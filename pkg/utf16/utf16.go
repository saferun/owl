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

package utf16

import (
	"reflect"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// StringToUTF16Ptr returns the pointer to UTF-8 encoded string. It will silently return
// an invalid pointer if `s` argument contains a NUL byte at any location.
func StringToUTF16Ptr(s string) *uint16 {
	var p *uint16
	p, _ = syscall.UTF16PtrFromString(s)
	return p
}

// PtrToString is like UTF16ToString, but takes *uint16
// as a parameter instead of []uint16.
func PtrToString(p unsafe.Pointer) string {
	if p == nil {
		return ""
	}
	var s []uint16
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(p)
	hdr.Cap = 1
	hdr.Len = 1
	for s[len(s)-1] != 0 {
		hdr.Cap++
		hdr.Len++
	}
	// Remove trailing NUL and decode into a Go string.
	return string(utf16.Decode(s[:len(s)-1]))
}
