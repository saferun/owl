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

package event

import (
	"unsafe"

	"github.com/mel2oo/win32/tdh"
	"github.com/mel2oo/win32/types"
	"github.com/saferun/owl/pkg/utf16"
)

func Parse(etype EType, evt *tdh.EventRecord, info *tdh.TraceEventInfo) []Param {
	var (
		count  = info.PropertyCount
		params = make([]Param, 0)
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

		params = append(params, *param)
	}

	return params
}
