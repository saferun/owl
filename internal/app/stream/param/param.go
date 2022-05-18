/*
 * Copyright 2022 by Mel2oo <https://github.com/saferun/owl>
 *
 * Licensed under the GNU General Public License version 3 (GPLv3)
 *
 * If you distribute GPL-licensed software the license requires
 * that you also distribute the complete, corresponding source
 * code (as defined by GPL) to that GPL-licensed software.
 *
 * You should have received a copy of the GNU General Public License
 * with this program. If not, see <https://www.gnu.org/licenses/>
 */

package param

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/mel2oo/win32/tdh"
	"github.com/mel2oo/win32/types"
	"github.com/saferun/owl/internal/app/event"
	"github.com/saferun/owl/pkg/utf16"
)

func Parse(etype event.EType, evt *tdh.EventRecord, info *tdh.TraceEventInfo) map[string]interface{} {
	var (
		count  = info.PropertyCount
		params = make(map[string]interface{}, count)
		props  = (*[1 << 30]tdh.EventPropertyInfo)(unsafe.Pointer(&info.EventPropertyInfoArray[0]))[:count:count]
	)

	for _, property := range props {
		pprop := unsafe.Pointer(uintptr(unsafe.Pointer(info)) + uintptr(property.NameOffset))
		pname := utf16.PtrToString(pprop)

		if Ignored(pname) {
			continue
		}

		pname = Canonicalize(pname)
		if len(pname) == 0 {
			continue
		}

		descriptor := &tdh.PropertyDataDescriptor{
			PropertyName: types.ULONGLONG(uintptr(unsafe.Pointer(pprop))),
			ArrayIndex:   0xFFFFFFFF,
		}

		var size types.ULONG = types.ULONG(SizeOf(pname))
		if size == 0 {
			errno := tdh.TdhGetPropertySize(evt, 0, nil, 1, descriptor, &size)
			if errno != types.ERROR_SUCCESS || size == 0 {
				continue
			}
		}

		buffer := make([]types.BYTE, size)
		errno := tdh.TdhGetProperty(evt, 0, nil, 1, descriptor, size, &buffer[0])
		if errno != types.ERROR_SUCCESS {
			continue
		}

		param, err := getParam(pname, buffer, size, property.StructType)
		if err != nil {
			continue
		}

		fmt.Println(param)
	}

	return params
}

type Param struct {
	Name  string
	Type  Type
	Value Value
}

func getParam(name string, buffer []types.BYTE, size types.ULONG, nonStructType tdh.StructType) (*Param, error) {
	if len(buffer) == 0 {
		return nil, errors.New("property buffer is empty")
	}

	var (
		typ   Type
		value Value
	)

	switch nonStructType.InType {
	case tdh.TdhIntypeUnicodestring:
		typ, value = UnicodeString, utf16.PtrToString(unsafe.Pointer(&buffer[0]))
	case tdh.TdhIntypeAnsiString:
		typ, value = AnsiString, string((*[1<<30 - 1]byte)(unsafe.Pointer(&buffer[0]))[:size-1:size-1])

	case tdh.TdhIntypeInt8:
		typ, value = Int8, *(*int8)(unsafe.Pointer(&buffer[0]))
	case tdh.TdhIntypeUint8:
		typ, value = Uint8, *(*uint8)(unsafe.Pointer(&buffer[0]))
		if nonStructType.OutType == tdh.TdhOutypeHexInt8 {
			typ = HexInt8
		}
	case tdh.TdhIntypeBoolean:
		typ, value = Bool, *(*bool)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeInt16:
		typ, value = Int16, *(*int16)(unsafe.Pointer(&buffer[0]))
	case tdh.TdhIntypeUint16:
		typ, value = Uint16, *(*uint16)(unsafe.Pointer(&buffer[0]))
		switch nonStructType.OutType {
		case tdh.TdhOutypeHexInt16:
			typ = HexInt16
		case tdh.TdhOutypePort:
			typ = Port
		}

	case tdh.TdhIntypeInt32:
		typ, value = Int32, *(*int32)(unsafe.Pointer(&buffer[0]))
	case tdh.TdhIntypeUint32:
		typ, value = Uint32, *(*uint32)(unsafe.Pointer(&buffer[0]))
		switch nonStructType.OutType {
		case tdh.TdhOutypeHexInt32:
			typ = HexInt32
		case tdh.TdhOutypeIpv4:
			typ = IPv4
		}

	case tdh.TdhIntypeInt64:
		typ, value = Int64, *(*int64)(unsafe.Pointer(&buffer[0]))
	case tdh.TdhIntypeUint64:
		typ, value = Uint64, *(*uint64)(unsafe.Pointer(&buffer[0]))
		if nonStructType.OutType == tdh.TdhOutypeHexInt64 {
			typ = HexInt64
		}

	case tdh.TdhIntypeFloat:
		typ, value = Float, *(*float32)(unsafe.Pointer(&buffer[0]))
	case tdh.TdhIntypeDouble:
		typ, value = Double, *(*float64)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeHexInt32:
		typ, value = HexInt32, *(*int32)(unsafe.Pointer(&buffer[0]))
	case tdh.TdhIntypeHexInt64:
		typ, value = HexInt64, *(*int64)(unsafe.Pointer(&buffer[0]))
	case tdh.TdhIntypePointer, tdh.TdhIntypeSizet:
		typ, value = HexInt64, *(*uint64)(unsafe.Pointer(&buffer[0]))
	case tdh.TdhIntypeSid:
		typ, value = SID, buffer
	case tdh.TdhIntypeWbemSid:
		typ, value = WbemSID, buffer
	case tdh.TdhIntypeBinary:
		if nonStructType.OutType == tdh.TdhOutypeIpv6 {
			typ, value = IPv6, buffer
		} else {
			typ, value = Binary, buffer
		}
	default:
		return nil, fmt.Errorf("unknown type for %q parameter", name)
	}

	return &Param{name, typ, value}, nil
}
