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

package ps

type Process struct {
	Pid       uint32
	Name      string
	Command   string
	FullPath  string
	Cwd       string
	SID       string
	SessionID uint8
	Integrity uint32
	Envs      map[string]string
	Threads   map[uint32]Thread
	Images    []Image
	Parent    *Process
}

type Thread struct {
	Tid            uint32
	Pid            uint32
	UserStackBase  string
	UserStackLimit string
	KrnlStackBase  string
	KrnlStackLimit string
	EntryPoint     string
}

type Image struct {
	Name     string
	Base     string
	Size     uint32
	Checksum uint32
}
