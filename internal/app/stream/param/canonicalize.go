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

const (
	uniqueProcessKey = "UniqueProcessKey"
	processID        = "ProcessId"
	parentID         = "ParentId"
	sessionID        = "SessionId"
	exitStatus       = "ExitStatus"
	dirTableBase     = "DirectoryTableBase"
	userSID          = "UserSID"
	imageFileName    = "ImageFileName"
	commandLine      = "CommandLine"

	tthreadID       = "TThreadId"
	issuingThreadID = "IssuingThreadId"
	ttid            = "TTID"

	pid = "PID"

	basePriority   = "BasePriority"
	ioPriority     = "IoPriority"
	pagePriority   = "PagePriority"
	stackBase      = "StackBase"
	stackLimit     = "StackLimit"
	userStackBase  = "UserStackBase"
	userStackLimit = "UserStackLimit"
	win32StartAddr = "Win32StartAddr"

	fileObject        = "FileObject"
	fileName          = "FileName"
	openPath          = "OpenPath"
	fileCreateOptions = "CreateOptions"
	fileShareAccess   = "ShareAccess"
	fileOffset        = "Offset"
	fileIoSize        = "IoSize"
	fileInfoClass     = "InfoClass"
	fileKey           = "FileKey"
	fileExtraInfo     = "ExtraInfo"
	fileIrpPtr        = "IrpPtr"

	keyName        = "KeyName"
	keyHandle      = "KeyHandle"
	registryStatus = "Status"
	ntStatus       = "NtStatus"

	imageBase        = "ImageBase"
	imageSize        = "ImageSize"
	imageChecksum    = "ImageCheckSum"
	imageDefaultBase = "DefaultBase"

	netDaddr = "daddr"
	netSaddr = "saddr"
	netDport = "dport"
	netSport = "sport"
	netSize  = "size"

	handleID         = "Handle"
	handleObject     = "Object"
	handleObjectName = "ObjectName"
	handleObjectType = "ObjectType"
)

var ignore = map[string]bool{
	"Reserved0":       true,
	"Reserved1":       true,
	"Reserved2":       true,
	"Reserved3":       true,
	"Reserved4":       true,
	"ThreadFlags":     true,
	"ApplicationId":   true,
	"PackageFullName": true,
	"IoFlags":         true,
}

// Ignored returns the collection of parameters that are ignored by kernel stream consumer.
func Ignored(name string) bool {
	return ignore[name]
}

// Canonicalize takes an original kernel event property name and normalizes it
// to canonical parameter name.
func Canonicalize(name string) string {
	switch name {
	case tthreadID, issuingThreadID, ttid:
		return ThreadID
	case processID, pid:
		return ProcessID
	case uniqueProcessKey:
		return ProcessObject
	case parentID:
		return ProcessParentID
	case sessionID:
		return SessionID
	case imageFileName:
		return ProcessName
	case commandLine:
		return Comm
	case userSID:
		return UserSID
	case exitStatus:
		return ExitStatus
	case dirTableBase:
		return DTB
	case basePriority:
		return BasePrio
	case ioPriority:
		return IOPrio
	case pagePriority:
		return PagePrio
	case stackBase:
		return KstackBase
	case stackLimit:
		return KstackLimit
	case userStackBase:
		return UstackBase
	case userStackLimit:
		return UstackLimit
	case win32StartAddr:
		return ThreadEntrypoint
	case fileObject:
		return FileObject
	case fileName, openPath:
		return FileName
	case fileCreateOptions:
		return FileCreateOptions
	case fileShareAccess:
		return FileShareMask
	case fileOffset:
		return FileOffset
	case fileIoSize:
		return FileIoSize
	case fileInfoClass:
		return FileInfoClass
	case fileKey:
		return FileKey
	case fileExtraInfo:
		return FileExtraInfo
	case fileIrpPtr:
		return FileIrpPtr
	case keyName:
		return RegKeyName
	case keyHandle:
		return RegKeyHandle
	case imageBase:
		return ImageBase
	case imageDefaultBase:
		return ImageDefaultBase
	case imageSize:
		return ImageSize
	case imageChecksum:
		return ImageCheckSum
	case netDaddr:
		return NetDIP
	case netSaddr:
		return NetSIP
	case netDport:
		return NetDport
	case netSport:
		return NetSport
	case netSize:
		return NetSize
	case handleID:
		return HandleID
	case handleObject:
		return HandleObject
	case handleObjectName:
		return HandleObjectName
	case handleObjectType:
		return HandleObjectTypeID
	case registryStatus, ntStatus:
		return NTStatus
	default:
		return ""
	}
}
