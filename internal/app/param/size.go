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

func SizeOf(kpar string) uint32 {
	switch kpar {
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
