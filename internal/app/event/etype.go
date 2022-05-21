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
	"syscall"
)

type EType [17]byte

func Pack(g syscall.GUID, op uint8) EType {
	return EType([17]byte{
		byte(g.Data1 >> 24), byte(g.Data1 >> 16), byte(g.Data1 >> 8), byte(g.Data1),
		byte(g.Data2 >> 8), byte(g.Data2), byte(g.Data3 >> 8), byte(g.Data3),
		g.Data4[0], g.Data4[1], g.Data4[2], g.Data4[3], g.Data4[4], g.Data4[5], g.Data4[6], g.Data4[7],
		op,
	})
}

var (
	ProcessGuid = syscall.GUID{Data1: 0x3d6fa8d0, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	ThreadGuid  = syscall.GUID{Data1: 0x3d6fa8d1, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	ImageGuid   = syscall.GUID{Data1: 0x2cb15d1d, Data2: 0x5fc1, Data3: 0x11d2, Data4: [8]byte{0xab, 0xe1, 0x00, 0xa0, 0xc9, 0x11, 0xf5, 0x18}}
	FileIOGuid  = syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x00, 0x00, 0xf8, 0x04, 0x64, 0xe3}}

	ALPCGuid             = syscall.GUID{Data1: 0x45d8cccd, Data2: 0x539f, Data3: 0x4b72, Data4: [8]byte{0xa8, 0xb7, 0x5c, 0x68, 0x31, 0x42, 0x60, 0x9a}}
	DiskIoGuid           = syscall.GUID{Data1: 0x3d6fa8d4, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	EventTraceConfigGuid = syscall.GUID{Data1: 0x01853a65, Data2: 0x418f, Data3: 0x4f36, Data4: [8]byte{0xae, 0xfc, 0xdc, 0x0f, 0x1d, 0x2f, 0xd2, 0x35}}
	ObTraceGuid          = syscall.GUID{Data1: 0x89497f50, Data2: 0xeffe, Data3: 0x4440, Data4: [8]byte{0x8c, 0xf2, 0xce, 0x6b, 0x1c, 0xdc, 0xac, 0xa7}}
	PageFaultGuid        = syscall.GUID{Data1: 0x3d6fa8d3, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	PerfInfoGuid         = syscall.GUID{Data1: 0xce1dbfb4, Data2: 0x137e, Data3: 0x4da6, Data4: [8]byte{0x87, 0xb0, 0x3f, 0x59, 0xaa, 0x10, 0x2c, 0xbc}}
	RegistryGuid         = syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}
	SplitIoGuid          = syscall.GUID{Data1: 0xd837ca92, Data2: 0x12b9, Data3: 0x44a5, Data4: [8]byte{0xad, 0x6a, 0x3a, 0x65, 0xb3, 0x57, 0x8a, 0xa8}}
	TcpIpGuid            = syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x00, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}
	UdpIpGuid            = syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xf0, 0xb7, 0xc8, 0x0f, 0x80}}

	// Process event
	OpProcessCreate    = Pack(ProcessGuid, 1)
	OpProcessTerminate = Pack(ProcessGuid, 2)
	OpProcessEnum      = Pack(ProcessGuid, 3)

	// Thread event
	OpThreadCreate    = Pack(ThreadGuid, 1)
	OpThreadTerminate = Pack(ThreadGuid, 2)
	OpThreadEnum      = Pack(ThreadGuid, 3)

	// ImageLoad event
	OpImageUnload = Pack(ImageGuid, 2)
	OpImageEnum   = Pack(ImageGuid, 3)
	OpImageLoad   = Pack(ImageGuid, 10)

	// file Event
	OpFileCreate  = Pack(FileIOGuid, 64)
	OpFileCleanup = Pack(FileIOGuid, 65)
	OpFileDelete  = Pack(FileIOGuid, 70)

	OpClose         = Pack(FileIOGuid, 66)
	OpReadFile      = Pack(FileIOGuid, 67)
	OpWriteFile     = Pack(FileIOGuid, 68)
	OpSetFileInfo   = Pack(FileIOGuid, 69)
	OpRenameFile    = Pack(FileIOGuid, 71)
	OpEnumDir       = Pack(FileIOGuid, 72)
	OpFlush         = Pack(FileIOGuid, 73)
	OpQueryFileInfo = Pack(FileIOGuid, 74)
	OpFSControl     = Pack(FileIOGuid, 75)
	OpOpEnd         = Pack(FileIOGuid, 76)
	OpNotifyDir     = Pack(FileIOGuid, 77)

	OpALPCSendMessage       = Pack(ALPCGuid, 33)
	OpALPCReceiveMessage    = Pack(ALPCGuid, 34)
	OpALPCWaitForReply      = Pack(ALPCGuid, 35)
	OpALPCWaitForNewMessage = Pack(ALPCGuid, 36)
	OpALPCUnwait            = Pack(ALPCGuid, 37)

	OpDiskIORead                  = Pack(DiskIoGuid, 10)
	OpDiskIOWrite                 = Pack(DiskIoGuid, 11)
	OpDiskIOReadInit              = Pack(DiskIoGuid, 12)
	OpDiskIOWriteInit             = Pack(DiskIoGuid, 13)
	OpDiskIOFlush                 = Pack(DiskIoGuid, 14)
	OpDiskIOFlushInit             = Pack(DiskIoGuid, 15)
	OpDiskIORedirectedInit        = Pack(DiskIoGuid, 16)
	OpDriverMajorFunctionCall     = Pack(DiskIoGuid, 34)
	OpDriverMajorFunctionReturn   = Pack(DiskIoGuid, 35)
	OpDriverCompletionRoutine     = Pack(DiskIoGuid, 37)
	OpDriverCompleteRequest       = Pack(DiskIoGuid, 52)
	OpDriverCompleteRequestReturn = Pack(DiskIoGuid, 53)

	OpRegCreateKey        = Pack(RegistryGuid, 10)
	OpRegOpenKey          = Pack(RegistryGuid, 11)
	OpRegDeleteKey        = Pack(RegistryGuid, 12)
	OpRegQueryKey         = Pack(RegistryGuid, 13)
	OpRegSetValue         = Pack(RegistryGuid, 14)
	OpRegDeleteValue      = Pack(RegistryGuid, 15)
	OpRegQueryValue       = Pack(RegistryGuid, 16)
	OpRegEnumKey          = Pack(RegistryGuid, 17)
	OpRegEnumValueKey     = Pack(RegistryGuid, 18)
	OpRegQueryMultValue   = Pack(RegistryGuid, 19)
	OpRegSetInformation   = Pack(RegistryGuid, 20)
	OpRegFlush            = Pack(RegistryGuid, 21)
	OpRegCreateKCB        = Pack(RegistryGuid, 22)
	OpRegDeleteKCB        = Pack(RegistryGuid, 23)
	OpRegKCBRundownBegine = Pack(RegistryGuid, 24)
	OpRegKCBRundownEnd    = Pack(RegistryGuid, 25)
	OpRegOpenKeyV1        = Pack(RegistryGuid, 27)

	OpSendTCPv4       = Pack(TcpIpGuid, 10)
	OpRecvTCPv4       = Pack(TcpIpGuid, 11)
	OpConnectTCPv4    = Pack(TcpIpGuid, 12)
	OpDisconnectTCPv4 = Pack(TcpIpGuid, 13)
	OpRetransmitTCPv4 = Pack(TcpIpGuid, 14)
	OpAcceptTCPv4     = Pack(TcpIpGuid, 15)
	OpReconnectTCPv4  = Pack(TcpIpGuid, 16)
	OpTCPFail         = Pack(TcpIpGuid, 17)
	OpCopyTCPv4       = Pack(TcpIpGuid, 18)
	OpSendTCPv6       = Pack(TcpIpGuid, 26)
	OpRecvTCPv6       = Pack(TcpIpGuid, 27)
	OpConnectTCPv6    = Pack(TcpIpGuid, 28)
	OpDisconnectTCPv6 = Pack(TcpIpGuid, 29)
	OpRetransmitTCPv6 = Pack(TcpIpGuid, 30)
	OpAcceptTCPv6     = Pack(TcpIpGuid, 31)
	OpReconnectTCPv6  = Pack(TcpIpGuid, 32)
	OpCopyTCPv6       = Pack(TcpIpGuid, 34)

	OpSendUDPv4 = Pack(UdpIpGuid, 10)
	OpRecvUDPV4 = Pack(UdpIpGuid, 11)
	OpUDPFail   = Pack(UdpIpGuid, 17)
	OpSendUDPv6 = Pack(UdpIpGuid, 26)
	OpRecvUDPV6 = Pack(UdpIpGuid, 27)

	OpConnect    = Pack(TcpIpGuid, 40)
	OpDisconnect = Pack(TcpIpGuid, 42)
	OpRetransmit = Pack(TcpIpGuid, 44)
	OpAccept     = Pack(TcpIpGuid, 46)
	OpReconnect  = Pack(TcpIpGuid, 47)
	OpSend       = Pack(TcpIpGuid, 72)
	OpRecv       = Pack(UdpIpGuid, 75)

	OpCreateHandle = Pack(ObTraceGuid, 32)
	OpCloseHandle  = Pack(ObTraceGuid, 33)
)

func (e EType) String() string {
	switch e {
	case OpProcessCreate:
		return "ProcessCreate"
	case OpProcessTerminate:
		return "ProcessTerminate"
	case OpProcessEnum:
		return "ProcessEnum"

	case OpThreadCreate:
		return "ThreadCreate"
	case OpThreadTerminate:
		return "ThreadTerminate"
	case OpThreadEnum:
		return "ThreadEnum"

	case OpImageUnload:
		return "ImageUnload"
	case OpImageEnum:
		return "ImageEnum"
	case OpImageLoad:
		return "ImageLoad"

	case OpFileCreate:
		return "FileCreate"
	case OpFileCleanup:
		return "FileCleanup"
	case OpFileDelete:
		return "FileDelete"
	}

	return ""
}

func (e EType) Exist() bool {
	switch e {
	case OpProcessCreate, OpProcessTerminate, OpProcessEnum:
		return true
	case OpThreadCreate, OpThreadTerminate, OpThreadEnum:
		return true
	case OpImageUnload, OpImageEnum, OpImageLoad:
		return true
	// case OpFileDelete, OpFileDelete:
	case OpFileCleanup:
		return true
	default:
		return false
	}
}

func (e EType) Dropped() bool {
	switch e {
	case OpProcessEnum, OpThreadEnum, OpImageEnum:
		return true
	default:
		return false
	}
}
