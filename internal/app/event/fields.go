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
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
	"unsafe"

	"github.com/mel2oo/win32/tdh"
	"github.com/mel2oo/win32/types"
	"github.com/saferun/owl/pkg/utf16"
)

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
		typ Type
		val Value
	)

	switch nonStructType.InType {
	case tdh.TdhIntypeUnicodestring:
		typ, val = UnicodeString, utf16.PtrToString(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeAnsiString:
		typ, val = AnsiString, string((*[1<<30 - 1]byte)(unsafe.Pointer(&buffer[0]))[:size-1:size-1])

	case tdh.TdhIntypeInt8:
		typ, val = Int8, *(*int8)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeUint8:
		typ, val = Uint8, *(*uint8)(unsafe.Pointer(&buffer[0]))
		if nonStructType.OutType == tdh.TdhOutypeHexInt8 {
			typ = HexInt8
		}

	case tdh.TdhIntypeBoolean:
		typ, val = Bool, *(*bool)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeInt16:
		typ, val = Int16, *(*int16)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeUint16:
		typ, val = Uint16, *(*uint16)(unsafe.Pointer(&buffer[0]))
		switch nonStructType.OutType {
		case tdh.TdhOutypeHexInt16:
			typ = HexInt16
		case tdh.TdhOutypePort:
			typ = Port
		}

	case tdh.TdhIntypeInt32:
		typ, val = Int32, *(*int32)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeUint32:
		typ, val = Uint32, *(*uint32)(unsafe.Pointer(&buffer[0]))
		switch nonStructType.OutType {
		case tdh.TdhOutypeHexInt32:
			typ = HexInt32
		case tdh.TdhOutypeIpv4:
			typ = IPv4
		}

	case tdh.TdhIntypeInt64:
		typ, val = Int64, *(*int64)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeUint64:
		typ, val = Uint64, *(*uint64)(unsafe.Pointer(&buffer[0]))
		if nonStructType.OutType == tdh.TdhOutypeHexInt64 {
			typ = HexInt64
		}

	case tdh.TdhIntypeFloat:
		typ, val = Float, *(*float32)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeDouble:
		typ, val = Double, *(*float64)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeHexInt32:
		typ, val = HexInt32, *(*int32)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeHexInt64:
		typ, val = HexInt64, *(*int64)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypePointer, tdh.TdhIntypeSizet:
		typ, val = HexInt64, *(*uint64)(unsafe.Pointer(&buffer[0]))

	case tdh.TdhIntypeSid:
		typ, val = SID, buffer

	case tdh.TdhIntypeWbemSid:
		typ, val = WbemSID, buffer

	case tdh.TdhIntypeBinary:
		if nonStructType.OutType == tdh.TdhOutypeIpv6 {
			typ, val = IPv6, buffer
		} else {
			typ, val = Binary, buffer
		}

	default:
		return nil, fmt.Errorf("unknown type for %q parameter", name)
	}

	return &Param{name, typ, val}, nil
}

func (p *Param) String() string {
	if p.Value == nil {
		return ""
	}

	switch p.Type {
	case UnicodeString, AnsiString, SID, WbemSID:
		return p.Value.(string)
	case HexInt32, HexInt64, HexInt16, HexInt8:
		return string(p.Value.(Hex))
	case Int8:
		return strconv.Itoa(int(p.Value.(int8)))
	case Uint8:
		return strconv.Itoa(int(p.Value.(uint8)))
	case Int16:
		return strconv.Itoa(int(p.Value.(int16)))
	case Uint16, Port:
		return strconv.Itoa(int(p.Value.(uint16)))
	case Uint32, PID, TID:
		return strconv.Itoa(int(p.Value.(uint32)))
	case Int32:
		return strconv.Itoa(int(p.Value.(int32)))
	case Uint64:
		return strconv.FormatUint(p.Value.(uint64), 10)
	case Int64:
		return strconv.Itoa(int(p.Value.(int64)))
	case IPv4, IPv6:
		return p.Value.(net.IP).String()
	case Bool:
		return strconv.FormatBool(p.Value.(bool))
	case Float:
		return strconv.FormatFloat(float64(p.Value.(float32)), 'f', 6, 32)
	case Double:
		return strconv.FormatFloat(p.Value.(float64), 'f', 6, 64)
	case Time:
		return p.Value.(time.Time).String()
	// case Enum:
	// 	switch typ := p.Value.(type) {
	// 	case fs.FileShareMode:
	// 		return typ.String()
	// 	case networp.L4Proto:
	// 		return typ.String()
	// 	case fs.FileDisposition:
	// 		return typ.String()
	// 	default:
	// 		return fmt.Sprintf("%v", p.Value)
	// 	}
	default:
		return fmt.Sprintf("%v", p.Value)
	}
}

const (
	// NTStatus is the parameter that identifies the NTSTATUS value.
	NTStatus = "status"

	// ProcessID represents the process identifier.
	ProcessID = "pid"
	// ProcessObject field represents the address of the process object in the kernel.
	ProcessObject = "proc"
	// ThreadID field represents the thread identifier.
	ThreadID = "tid"
	// ProcessParentID field represents the parent process identifier.
	ProcessParentID = "ppid"
	// SessionID fields represents the session identifier.
	SessionID = "session_id"
	// UserSID field is the security identifier associated to the process token under which it is run.
	UserSID = "sid"
	// ProcessName field denotes the process image name.
	ProcessName = "name"
	// Exe field denotes the full path of the executable.
	Exe = "exe"
	// Comm field represents the process command line.
	Comm = "command"
	// DTB field denotes the address of the process directory table.
	DTB = "directory_table_base"
	// ExitStatus is the field that represents the process exit status.
	ExitStatus = "exit_status"
	// StartTime field denotes the process start time.
	StartTime = "start_time"

	// BasePrio field is the thread base priority assigned by the scheduler.
	BasePrio = "base_prio"
	// IOPrio represents the filed that indicates the thread I/O priority.
	IOPrio = "io_prio"
	// PagePrio field denotes page priority.
	PagePrio = "page_prio"
	// KstackBase field is the start address of the kernel space stacp.
	KstackBase = "kstack"
	// KstackLimit field is the end address of the kernel space stacp.
	KstackLimit = "kstack_limit"
	// UstackBase field is the start address of the user space stacp.
	UstackBase = "ustack"
	// UstackLimit field is the end address of the user space stacp.
	UstackLimit = "ustack_limit"
	// ThreadEntrypoint field is the address of the thread main function.
	ThreadEntrypoint = "entrypoint"

	// FileObject determines the field name for the file object pointer.
	FileObject = "file_object"
	// FileName represents the field that designates the absolute path of the file.
	FileName = "file_name"
	// FileCreateOptions is the field that represents the values passed in the CreateDispositions parameter to the NtCreateFile function.
	FileCreateOptions = "options"
	// FileOperation is the field that represents the values passed in the CreateOptions parameter to the NtCreateFile function.
	FileOperation = "operation"
	// FileCreated represents the name for the file creation field.
	FileCreated = "created"
	// FileAccessed represents the name for the file access field.
	FileAccessed = "accessed"
	// FileModified represents the name for the file modification field.
	FileModified = "modified"
	// FileShareMask represents the field name for the share access masp.
	FileShareMask = "share_mask"
	// FileType represents the field name that indicates the file type.
	FileType = "type"
	// FileAttributes is the field that represents file attribute values.
	FileAttributes = "attributes"
	// FileIoSize is the filed that represents the number of bytes in file read/write operations.
	FileIoSize = "io_size"
	// FileOffset represents the file for the file offset in read/write operations.
	FileOffset = "offset"
	// FileInfoClass represents the file information class.
	FileInfoClass = "class"
	// FileKey represents the directory key identifier in EnumDirectory events.
	FileKey = "file_key"
	// FileDirectory represents the filed for the directory name in EnumDirectory events.
	FileDirectory = "dir"
	// FileIrpPtr represents the I/O request packet id.
	FileIrpPtr = "irp"
	// FileExtraInfo is the parameter that represents extra information returned by the file system for the operation. For example for a read request, the actual number of bytes that were read.
	FileExtraInfo = "extra_info"

	// RegKeyHandle identifies the parameter name for the registry key handle.
	RegKeyHandle = "key_handle"
	// RegKeyName represents the parameter name for the fully qualified key name.
	RegKeyName = "key_name"
	// RegValue identifies the parameter name that contains the value
	RegValue = "value"
	// RegValueType identifies the parameter that represents registry value type e.g (DWORD, BINARY)
	RegValueType = "type"

	// ImageBase identifies the parameter name for the base address of the process in which the image is loaded.
	ImageBase = "base_address"
	// ImageSize represents the parameter name for the size of the image in bytes.
	ImageSize = "image_size"
	// ImageCheckSum is the parameter name for image checksum.
	ImageCheckSum = "checksum"
	// ImageDefaultBase is the parameter name that represents image's base address.
	ImageDefaultBase = "default_address"
	// ImageFilename is the parameter name that denotes file name and extension of the DLL/executable image.
	ImageFilename = "file_name"

	// NetSize identifies the parameter name that represents the packet size.
	NetSize = "size"
	// NetDIP is the parameter name that denotes the destination IP address.
	NetDIP = "dip"
	// NetSIP is the parameter name that denotes the source IP address.
	NetSIP = "sip"
	// NetDport identifies the parameter name that represents destination port number.
	NetDport = "dport"
	// NetSport identifies the parameter name that represents source port number.
	NetSport = "sport"
	// NetMSS is the parameter name that represents the maximum TCP segment size.
	NetMSS = "mss"
	// NetRcvWin is the parameter name that represents TCP segment's receive window size.
	NetRcvWin = "rcvwin"
	// NetSAckopt is the parameter name that represents Selective Acknowledgment option in TCP header.
	NetSAckopt = "sack_opt"
	// NetTsopt is the parameter name that represents the time stamp option in TCP header.
	NetTsopt = "timestamp_opt"
	// NetWsopt is the parameter name that represents the window scale option in TCP header.
	NetWsopt = "window_scale_opt"
	// NetRcvWinScale is the parameter name that represents the TCP receive window scaling factor.
	NetRcvWinScale = "recv_winscale"
	// NetSendWinScale is the parameter name that represents the TCP send window scaling factor.
	NetSendWinScale = "send_winscale"
	// NetSeqNum is the parameter name that represents that represents the TCP sequence number.
	NetSeqNum = "seqnum"
	// NetConnID is the parameter name that represents a unique connection identifier.
	NetConnID = "connid"
	// NetL4Proto is the parameter name that identifies the Layer 4 protocol name.
	NetL4Proto = "l4_proto"
	// NetDportName is the field that denotes the destination port name.
	NetDportName = "dport_name"
	// NetSportName is the field that denotes the source port name.
	NetSportName = "sport_name"
	// NetSIPNames is the field that denotes the source IP address names.
	NetSIPNames = "sip_names"
	// NetDIPNames is the field that denotes the destination IP address names.
	NetDIPNames = "dip_names"

	// HandleID identifies the parameter that specifies the handle identifier.
	HandleID = "handle_id"
	// HandleObject identifies the parameter that represents the kernel object to which handle is associated.
	HandleObject = "handle_object"
	// HandleObjectName identifies the parameter that represents the kernel object name.
	HandleObjectName = "handle_name"
	// HandleObjectTypeID identifies the parameter that represents the kernel object type identifier.
	HandleObjectTypeID = "type_id"
	// HandleObjectTypeName identifies the parameter that represents the kernel object type name.
	HandleObjectTypeName = "handle_type"
)

func SizeOf(field string) uint32 {
	switch field {
	case RegKeyHandle, KstackLimit, KstackBase, UstackLimit,
		UstackBase, ThreadEntrypoint, ImageBase, ImageSize,
		ImageDefaultBase, DTB, ProcessObject, FileIrpPtr, FileObject,
		FileExtraInfo, HandleObject, FileKey, FileOffset:
		return 8
	case NTStatus, ProcessID, ThreadID, ProcessParentID,
		SessionID, ExitStatus, FileCreateOptions, FileShareMask,
		HandleID, FileIoSize, FileInfoClass:
		return 4
	case NetDport, NetSport, HandleObjectTypeID:
		return 2
	case PagePrio, BasePrio, IOPrio:
		return 1
	default:
		return 0
	}
}
