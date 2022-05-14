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

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/mel2oo/win32/kernel32"
	"github.com/mel2oo/win32/shell32"
	"github.com/mel2oo/win32/types"
)

type Option func(*ExecOption)

type ExecOption struct {
	File    string
	Command string
	Admin   bool
}

func NewExecuteInfo() *ExecOption {
	return &ExecOption{
		File: "C:\\Windows\\System32\\cmd.exe",
	}
}

func WithFilePath(file string) Option {
	return func(ei *ExecOption) {
		ei.File = file
	}
}

func WithCommand(cmd string) Option {
	return func(ei *ExecOption) {
		ei.Command = cmd
	}
}

func WithAdmin() Option {
	return func(ei *ExecOption) {
		ei.Admin = true
	}
}

func Exec(opts ...Option) error {
	ei := NewExecuteInfo()
	for _, o := range opts {
		o(ei)
	}

	var lpExecInfo types.SHELLEXECUTEINFO

	lpExecInfo.Size = types.DWORD(unsafe.Sizeof(lpExecInfo))
	lpExecInfo.Mask = types.SEE_MASK_NOCLOSEPROCESS

	if ei.Admin {
		verb, err := syscall.UTF16PtrFromString("runas")
		if err != nil {
			return err
		}
		lpExecInfo.Verb = verb
	} else {
		verb, err := syscall.UTF16PtrFromString("open")
		if err != nil {
			return err
		}
		lpExecInfo.Verb = verb
	}

	file, err := syscall.UTF16PtrFromString(ei.File)
	if err != nil {
		return err
	}
	lpExecInfo.File = file
	lpExecInfo.Show = types.SW_SHOW

	param, err := syscall.UTF16PtrFromString(ei.Command)
	if err != nil {
		return err
	}
	lpExecInfo.Parameters = param

	if b := shell32.ShellExecuteExW(&lpExecInfo); b == 0 {
		return fmt.Errorf("shell execute error, %d", kernel32.GetLastError())
	}

	return nil
}
