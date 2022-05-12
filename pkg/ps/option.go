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

type Option func(*ExecuteInfo)

type ExecuteInfo struct {
	File    string
	Command string
	Admin   bool
}

func NewExecuteInfo() *ExecuteInfo {
	return &ExecuteInfo{
		File: "C:\\Windows\\System32\\cmd.exe",
	}
}

func WithFilePath(file string) Option {
	return func(ei *ExecuteInfo) {
		ei.File = file
	}
}

func WithCommand(cmd string) Option {
	return func(ei *ExecuteInfo) {
		ei.Command = cmd
	}
}

func WithAdmin() Option {
	return func(ei *ExecuteInfo) {
		ei.Admin = true
	}
}
