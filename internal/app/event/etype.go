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
	ALPCGuid             = syscall.GUID{Data1: 0x45d8cccd, Data2: 0x539f, Data3: 0x4b72, Data4: [8]byte{0xa8, 0xb7, 0x5c, 0x68, 0x31, 0x42, 0x60, 0x9a}}
	DiskIoGuid           = syscall.GUID{Data1: 0x3d6fa8d4, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	EventTraceConfigGuid = syscall.GUID{Data1: 0x01853a65, Data2: 0x418f, Data3: 0x4f36, Data4: [8]byte{0xae, 0xfc, 0xdc, 0x0f, 0x1d, 0x2f, 0xd2, 0x35}}
	FileIOGuid           = syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x00, 0x00, 0xf8, 0x04, 0x64, 0xe3}}
	ImageLoadGuid        = syscall.GUID{Data1: 0x2cb15d1d, Data2: 0x5fc1, Data3: 0x11d2, Data4: [8]byte{0xab, 0xe1, 0x00, 0xa0, 0xc9, 0x11, 0xf5, 0x18}}
	ObTraceGuid          = syscall.GUID{Data1: 0x89497f50, Data2: 0xeffe, Data3: 0x4440, Data4: [8]byte{0x8c, 0xf2, 0xce, 0x6b, 0x1c, 0xdc, 0xac, 0xa7}}
	PageFaultGuid        = syscall.GUID{Data1: 0x3d6fa8d3, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	PerfInfoGuid         = syscall.GUID{Data1: 0xce1dbfb4, Data2: 0x137e, Data3: 0x4da6, Data4: [8]byte{0x87, 0xb0, 0x3f, 0x59, 0xaa, 0x10, 0x2c, 0xbc}}
	ProcessGuid          = syscall.GUID{Data1: 0x3d6fa8d0, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	RegistryGuid         = syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}
	SplitIoGuid          = syscall.GUID{Data1: 0xd837ca92, Data2: 0x12b9, Data3: 0x44a5, Data4: [8]byte{0xad, 0x6a, 0x3a, 0x65, 0xb3, 0x57, 0x8a, 0xa8}}
	TcpIpGuid            = syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x00, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}
	ThreadGuid           = syscall.GUID{Data1: 0x3d6fa8d1, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	UdpIpGuid            = syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xf0, 0xb7, 0xc8, 0x0f, 0x80}}

	OpcodeALPCSendMessage       = Pack(ALPCGuid, 33)
	OpcodeALPCReceiveMessage    = Pack(ALPCGuid, 34)
	OpcodeALPCWaitForReply      = Pack(ALPCGuid, 35)
	OpcodeALPCWaitForNewMessage = Pack(ALPCGuid, 36)
	OpcodeALPCUnwait            = Pack(ALPCGuid, 37)

	OpcodeDiskIORead                  = Pack(DiskIoGuid, 10)
	OpcodeDiskIOWrite                 = Pack(DiskIoGuid, 11)
	OpcodeDiskIOReadInit              = Pack(DiskIoGuid, 12)
	OpcodeDiskIOWriteInit             = Pack(DiskIoGuid, 13)
	OpcodeDiskIOFlush                 = Pack(DiskIoGuid, 14)
	OpcodeDiskIOFlushInit             = Pack(DiskIoGuid, 15)
	OpcodeDiskIORedirectedInit        = Pack(DiskIoGuid, 16)
	OpcodeDriverMajorFunctionCall     = Pack(DiskIoGuid, 34)
	OpcodeDriverMajorFunctionReturn   = Pack(DiskIoGuid, 35)
	OpcodeDriverCompletionRoutine     = Pack(DiskIoGuid, 37)
	OpcodeDriverCompleteRequest       = Pack(DiskIoGuid, 52)
	OpcodeDriverCompleteRequestReturn = Pack(DiskIoGuid, 53)

	OpcodeFileName      = Pack(FileIOGuid, 0)
	OpcodeFileCreate    = Pack(FileIOGuid, 32)
	OpcodeFileDelete    = Pack(FileIOGuid, 35)
	OpcodeFileRundown   = Pack(FileIOGuid, 36)
	OpcodeCreateFile    = Pack(FileIOGuid, 64)
	OpcodeCleanup       = Pack(FileIOGuid, 65)
	OpcodeClose         = Pack(FileIOGuid, 66)
	OpcodeReadFile      = Pack(FileIOGuid, 67)
	OpcodeWriteFile     = Pack(FileIOGuid, 68)
	OpcodeSetFileInfo   = Pack(FileIOGuid, 69)
	OpcodeDeleteFile    = Pack(FileIOGuid, 70)
	OpcodeRenameFile    = Pack(FileIOGuid, 71)
	OpcodeEnumDir       = Pack(FileIOGuid, 72)
	OpcodeFlush         = Pack(FileIOGuid, 73)
	OpcodeQueryFileInfo = Pack(FileIOGuid, 74)
	OpcodeFSControl     = Pack(FileIOGuid, 75)
	OpcodeOpEnd         = Pack(FileIOGuid, 76)
	OpcodeNotifyDir     = Pack(FileIOGuid, 77)

	OpcodeImageUnload    = Pack(ImageLoadGuid, 2)
	OpcodeImageEnumStart = Pack(ImageLoadGuid, 3)
	OpcodeImageEnumEnd   = Pack(ImageLoadGuid, 4)
	OpcodeImageLoad      = Pack(ImageLoadGuid, 10)

	OpcodeProcessCreate    = Pack(ProcessGuid, 1)
	OpcodeProcessTerminate = Pack(ProcessGuid, 2)
	// OpcodeProcessEnum      = Pack(ProcessGuid, 3)

	OpcodeThreadCreate    = Pack(ThreadGuid, 1)
	OpcodeThreadTerminate = Pack(ThreadGuid, 2)
	OpcodeThreadEnumStart = Pack(ThreadGuid, 3)
	OpcodeThreadEnumEnd   = Pack(ThreadGuid, 4)

	OpcodeRegCreateKey        = Pack(RegistryGuid, 10)
	OpcodeRegOpenKey          = Pack(RegistryGuid, 11)
	OpcodeRegDeleteKey        = Pack(RegistryGuid, 12)
	OpcodeRegQueryKey         = Pack(RegistryGuid, 13)
	OpcodeRegSetValue         = Pack(RegistryGuid, 14)
	OpcodeRegDeleteValue      = Pack(RegistryGuid, 15)
	OpcodeRegQueryValue       = Pack(RegistryGuid, 16)
	OpcodeRegEnumKey          = Pack(RegistryGuid, 17)
	OpcodeRegEnumValueKey     = Pack(RegistryGuid, 18)
	OpcodeRegQueryMultValue   = Pack(RegistryGuid, 19)
	OpcodeRegSetInformation   = Pack(RegistryGuid, 20)
	OpcodeRegFlush            = Pack(RegistryGuid, 21)
	OpcodeRegCreateKCB        = Pack(RegistryGuid, 22)
	OpcodeRegDeleteKCB        = Pack(RegistryGuid, 23)
	OpcodeRegKCBRundownBegine = Pack(RegistryGuid, 24)
	OpcodeRegKCBRundownEnd    = Pack(RegistryGuid, 25)
	OpcodeRegOpenKeyV1        = Pack(RegistryGuid, 27)

	OpcodeSendTCPv4       = Pack(TcpIpGuid, 10)
	OpcodeRecvTCPv4       = Pack(TcpIpGuid, 11)
	OpcodeConnectTCPv4    = Pack(TcpIpGuid, 12)
	OpcodeDisconnectTCPv4 = Pack(TcpIpGuid, 13)
	OpcodeRetransmitTCPv4 = Pack(TcpIpGuid, 14)
	OpcodeAcceptTCPv4     = Pack(TcpIpGuid, 15)
	OpcodeReconnectTCPv4  = Pack(TcpIpGuid, 16)
	OpcodeTCPFail         = Pack(TcpIpGuid, 17)
	OpcodeCopyTCPv4       = Pack(TcpIpGuid, 18)
	OpcodeSendTCPv6       = Pack(TcpIpGuid, 26)
	OpcodeRecvTCPv6       = Pack(TcpIpGuid, 27)
	OpcodeConnectTCPv6    = Pack(TcpIpGuid, 28)
	OpcodeDisconnectTCPv6 = Pack(TcpIpGuid, 29)
	OpcodeRetransmitTCPv6 = Pack(TcpIpGuid, 30)
	OpcodeAcceptTCPv6     = Pack(TcpIpGuid, 31)
	OpcodeReconnectTCPv6  = Pack(TcpIpGuid, 32)
	OpcodeCopyTCPv6       = Pack(TcpIpGuid, 34)

	OpcodeSendUDPv4 = Pack(UdpIpGuid, 10)
	OpcodeRecvUDPV4 = Pack(UdpIpGuid, 11)
	OpcodeUDPFail   = Pack(UdpIpGuid, 17)
	OpcodeSendUDPv6 = Pack(UdpIpGuid, 26)
	OpcodeRecvUDPV6 = Pack(UdpIpGuid, 27)

	OpcodeConnect    = Pack(TcpIpGuid, 40)
	OpcodeDisconnect = Pack(TcpIpGuid, 42)
	OpcodeRetransmit = Pack(TcpIpGuid, 44)
	OpcodeAccept     = Pack(TcpIpGuid, 46)
	OpcodeReconnect  = Pack(TcpIpGuid, 47)
	OpcodeSend       = Pack(TcpIpGuid, 72)
	OpcodeRecv       = Pack(UdpIpGuid, 75)

	OpcodeCreateHandle = Pack(ObTraceGuid, 32)
	OpcodeCloseHandle  = Pack(ObTraceGuid, 33)
)

func (e EType) String() string {
	switch e {
	case OpcodeProcessCreate:
		return "ProcessCreate"
	case OpcodeProcessTerminate:
		return "ProcessTerminate"
	}

	return ""
}

func (e EType) Exist() bool {
	switch e {
	case OpcodeProcessCreate:
		return true
	case OpcodeProcessTerminate:
		return true
	}

	return false
}
