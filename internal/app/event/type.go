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

import "syscall"

type Pack struct {
	Guid   syscall.GUID
	Opcode uint8
}

var (
	ALPCGuid             = syscall.GUID{Data1: 0x45d8cccd, Data2: 0x539f, Data3: 0x4b72, Data4: [8]byte{0xa8, 0xb7, 0x5c, 0x68, 0x31, 0x42, 0x60, 0x9a}}
	DiskIoGuid           = syscall.GUID{Data1: 0x3d6fa8d4, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	EventTraceConfigGuid = syscall.GUID{Data1: 0x01853a65, Data2: 0x418f, Data3: 0x4f36, Data4: [8]byte{0xae, 0xfc, 0xdc, 0x0f, 0x1d, 0x2f, 0xd2, 0x35}}
	FileIOGuid           = syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x00, 0x00, 0xf8, 0x04, 0x64, 0xe3}}
	ImageLoadGuid        = syscall.GUID{Data1: 0x2cb15d1d, Data2: 0x5fc1, Data3: 0x11d2, Data4: [8]byte{0xab, 0xe1, 0x00, 0xa0, 0xc9, 0x11, 0xf5, 0x18}}
	PageFaultGuid        = syscall.GUID{Data1: 0x3d6fa8d3, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	PerfInfoGuid         = syscall.GUID{Data1: 0xce1dbfb4, Data2: 0x137e, Data3: 0x4da6, Data4: [8]byte{0x87, 0xb0, 0x3f, 0x59, 0xaa, 0x10, 0x2c, 0xbc}}
	ProcessGuid          = syscall.GUID{Data1: 0x3d6fa8d0, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	RegistryGuid         = syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}
	SplitIoGuid          = syscall.GUID{Data1: 0xd837ca92, Data2: 0x12b9, Data3: 0x44a5, Data4: [8]byte{0xad, 0x6a, 0x3a, 0x65, 0xb3, 0x57, 0x8a, 0xa8}}
	TcpIpGuid            = syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x00, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}
	ThreadGuid           = syscall.GUID{Data1: 0x3d6fa8d1, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	UdpIpGuid            = syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xf0, 0xb7, 0xc8, 0x0f, 0x80}}

	OpcodeALPCSendMessage       = Pack{ALPCGuid, 33}
	OpcodeALPCReceiveMessage    = Pack{ALPCGuid, 34}
	OpcodeALPCWaitForReply      = Pack{ALPCGuid, 35}
	OpcodeALPCWaitForNewMessage = Pack{ALPCGuid, 36}
	OpcodeALPCUnwait            = Pack{ALPCGuid, 37}

	OpcodeDiskIORead                  = Pack{DiskIoGuid, 10}
	OpcodeDiskIOWrite                 = Pack{DiskIoGuid, 11}
	OpcodeDiskIOReadInit              = Pack{DiskIoGuid, 12}
	OpcodeDiskIOWriteInit             = Pack{DiskIoGuid, 13}
	OpcodeDiskIOFlush                 = Pack{DiskIoGuid, 14}
	OpcodeDiskIOFlushInit             = Pack{DiskIoGuid, 15}
	OpcodeDiskIORedirectedInit        = Pack{DiskIoGuid, 16}
	OpcodeDriverCompleteRequest       = Pack{DiskIoGuid, 52}
	OpcodeDriverCompleteRequestReturn = Pack{DiskIoGuid, 53}
	OpcodeDriverCompletionRoutine     = Pack{DiskIoGuid, 37}
	OpcodeDriverMajorFunctionCall     = Pack{DiskIoGuid, 34}
	OpcodeDriverMajorFunctionReturn   = Pack{DiskIoGuid, 35}

	OpcodeFileName    = Pack{FileIOGuid, 0}
	OpcodeFileCreate  = Pack{FileIOGuid, 32}
	OpcodeFileDelete  = Pack{FileIOGuid, 35}
	OpcodeFileRundown = Pack{FileIOGuid, 36}

	OpcodeCreateFile    = Pack{FileIOGuid, 64}
	OpcodeEnumDir       = Pack{FileIOGuid, 72}
	OpcodeNotifyDir     = Pack{FileIOGuid, 77}
	OpcodeSetFileInfo   = Pack{FileIOGuid, 69}
	OpcodeDeleteFile    = Pack{FileIOGuid, 70}
	OpcodeRenameFile    = Pack{FileIOGuid, 71}
	OpcodeQueryFileInfo = Pack{FileIOGuid, 74}
	OpcodeFSControl     = Pack{FileIOGuid, 75}
	OpcodeOpEnd         = Pack{FileIOGuid, 76}
	OpcodeReadFile      = Pack{FileIOGuid, 67}
	OpcodeWriteFile     = Pack{FileIOGuid, 68}
	OpcodeCleanup       = Pack{FileIOGuid, 65}
	OpcodeClose         = Pack{FileIOGuid, 66}
	OpcodeFlush         = Pack{FileIOGuid, 73}

	OpcodeImageLoad      = Pack{ImageLoadGuid, 10}
	OpcodeImageUnload    = Pack{ImageLoadGuid, 2}
	OpcodeImageEnumStart = Pack{ImageLoadGuid, 3}
	OpcodeImageEnumEnd   = Pack{ImageLoadGuid, 4}

	OpcodeProcessCreate    = Pack{ProcessGuid, 1}
	OpcodeProcessTerminate = Pack{ProcessGuid, 2}
	OpcodeProcessEnumStart = Pack{ProcessGuid, 3}
	OpcodeProcessEnumEnd   = Pack{ProcessGuid, 4}
)

const (
	OpcodeCreateThread    = 1
	OpcodeTerminateThread = 2
	OpcodeEnumThread      = 3
)

var (
	// RegCreateKey represents registry key creation kernel events
	RegCreateKey = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 10)
	// RegOpenKey represents registry open key kernel events
	RegOpenKey = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 11)
	// RegDeleteKey represents registry key deletion kernel events
	RegDeleteKey = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 12)
	// RegQueryKey represents registry query key kernel events
	RegQueryKey = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 13)
	// RegSetValue represents registry set value kernel events
	RegSetValue = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 14)
	// RegDeleteValue are kernel events for registry value removals
	RegDeleteValue = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 15)
	// RegQueryValue are kernel events for registry value queries
	RegQueryValue = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 16)
	// RegCreateKCB represents kernel events for KCB (Key Control Block) creation requests
	RegCreateKCB = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 22)
	// RegDeleteKCB represents kernel events for KCB(Key Control Block) closures
	RegDeleteKCB = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 23)
	// RegKCBRundown enumerates the registry keys open at the start of the kernel session.
	RegKCBRundown = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 25)
	// RegOpenKeyV1 represents registry open key kernel event. It looks like this event type defines identical kernel event type as RegOpenKey
	RegOpenKeyV1 = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 27)

	// AcceptTCPv4 represents the TCPv4 kernel events for accepting connection requests from the socket queue.
	AcceptTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 15)
	// AcceptTCPv6 represents the TCPv6 kernel events for accepting connection requests from the socket queue.
	AcceptTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 31)
	// SendTCPv4 represents the TCPv4 kernel events for sending data to the connected socket.
	SendTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 10)
	// SendTCPv6 represents the TCPv6 kernel events for sending data to the connected socket.
	SendTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 26)
	// SendUDPv4 represents the UDPv4 kernel events for sending datagrams to connectionless sockets.
	SendUDPv4 = Pack(syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xf0, 0xb7, 0xc8, 0x0f, 0x80}}, 10)
	// SendUDPv6 represents the UDPv6 kernel events for sending datagrams to connectionless sockets.
	SendUDPv6 = Pack(syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xf0, 0xb7, 0xc8, 0x0f, 0x80}}, 26)

	// RecvTCPv4 represents the TCP IPv4 network receive event.
	RecvTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 11)
	// RecvTCPv6 represents the TCP IPv6 network receive event.
	RecvTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 27)
	// RecvUDPv4 represents the UDP IPv4 network receive event.
	RecvUDPv4 = Pack(syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xf0, 0xb7, 0xc8, 0x0f, 0x80}}, 11)
	// RecvUDPv6 represents the UDP IPv6 network receive event.
	RecvUDPv6 = Pack(syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xf0, 0xb7, 0xc8, 0x0f, 0x80}}, 27)

	// ConnectTCPv4 represents the TCP IPv4 network connect event.
	ConnectTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 12)
	// ConnectTCPv6 represents the TCP IPv6 network connect event.
	ConnectTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 28)

	// DisconnectTCPv4 is the TCP IPv4 network disconnect event.
	DisconnectTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 13)
	// DisconnectTCPv6 is the TCP IPv6 network disconnect event.
	DisconnectTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 29)
	// Disconnect encompasses TCP IPv4/6 network disconnect events.
	Disconnect = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 42)

	// ReconnectTCPv4 is the TCP IPv4 network reconnect event.
	ReconnectTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 16)
	// ReconnectTCPv6 is the TCP IPv6 network reconnect event.
	ReconnectTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 32)

	// RetransmitTCPv4 is the TCP IPv4 network retransmit event.
	RetransmitTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 14)
	// RetransmitTCPv6 is the TCP IPv6 network retransmit event.
	RetransmitTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 30)

	// Accept represents the global kernel event type for both TCP v4/v6 connections. Note this is an artificial kernel event that is never published by the provider.
	Accept = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 46)
	// Send represents the global kernel event for all variants of sending data to sockets. Note this is an artificial kernel event that is never published by the provider.
	Send = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xa9c9, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 72)
	// Recv represents the global kernel event for all variants of receiving data from sockets. Note this is an artificial kernel event that is never published by the provider.
	Recv = Pack(syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xc8e0, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xc0, 0xb7, 0xc8, 0x0f, 0x80}}, 75)
	// Connect represents the global kernel event for all variants of connecting to sockets sockets. Note this is an artificial kernel event that is never published by the provider.
	Connect = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 40)
	// Reconnect represents the global kernel event for all variants of reconnecting to sockets sockets. Note this is an artificial kernel event that is never published by the provider.
	Reconnect = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 47)
	// Retransmit represents the global kernel event for all variants of retransmitting TCP segments. Note this is an artificial kernel event that is never published by the provider.
	Retransmit = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 44)

	// CreateHandle represents handle creation kernel event
	CreateHandle = Pack(syscall.GUID{Data1: 0x89497f50, Data2: 0xeffe, Data3: 0x4440, Data4: [8]byte{0x8c, 0xf2, 0xce, 0x6b, 0x1c, 0xdc, 0xac, 0xa7}}, 32)
	// CloseHandle represents handle closure kernel event
	CloseHandle = Pack(syscall.GUID{Data1: 0x89497f50, Data2: 0xeffe, Data3: 0x4440, Data4: [8]byte{0x8c, 0xf2, 0xce, 0x6b, 0x1c, 0xdc, 0xac, 0xa7}}, 33)

	// UnknownKtype designates unknown kernel event type
	UnknownKtype = Pack(syscall.GUID{}, 0)
)
