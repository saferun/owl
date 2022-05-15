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

package etw

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"

	"github.com/mel2oo/win32/advapi32/evntrace"
	"github.com/mel2oo/win32/types"
)

const (
	maxBufferSize = 1024
	maxStringLen  = 1024
)

type EventTrace struct {
	Handle evntrace.TraceHandle
}

func NewEventTrace() *EventTrace {
	return nil
}

func (e *EventTrace) Start() error {
	flags := evntrace.EventTraceFlagProcess
	flags |= evntrace.EventTraceFlagThread
	flags |= evntrace.EventTraceFlagImageLoad
	flags |= evntrace.EventTraceFlagNetworkTCPIP
	flags |= evntrace.EventTraceFlagRegistry
	flags |= evntrace.EventTraceFlagDiskFileIO
	flags |= evntrace.EventTraceFlagFileIO
	flags |= evntrace.EventTraceFlagFileIOInit

	bufferSize := maxBufferSize
	minBuffers := uint32(runtime.NumCPU() * 2)
	maxBuffers := minBuffers + 20
	flushTimer := time.Second

	props := &evntrace.EventTraceProperties{
		Wnode: evntrace.WnodeHeader{
			BufferSize: types.ULONG(unsafe.Sizeof(evntrace.EventTraceProperties{}) + 2*maxStringLen),
			Guid:       types.GUID(evntrace.SystemTraceControlGuid),
			Flags:      evntrace.WnodeFlagTracedGUID,
		},
		BufferSize:        types.ULONG(bufferSize),
		MinimumBuffers:    types.ULONG(minBuffers),
		MaximumBuffers:    types.ULONG(maxBuffers),
		LogFileMode:       evntrace.EventTraceRealTimeMode,
		FlushTimer:        types.ULONG(flushTimer.Seconds()),
		EnableFlags:       flags,
		LogFileNameOffset: 0,
		LoggerNameOffset:  types.ULONG(unsafe.Sizeof(evntrace.EventTraceProperties{})),
	}

	errno := evntrace.StartTrace(&e.Handle, evntrace.KernelLoggerName, props)
	if errno != 0 {
		fmt.Println("StartTrace failure", errno)
		return nil
	}
	return nil
}

func (e *EventTrace) Process() error {
	return nil
}

func (e *EventTrace) Close() error {
	return nil
}
