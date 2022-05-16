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
	"syscall"
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
	opts   *options
	props  *evntrace.EventTraceProperties
	handle evntrace.TraceHandle
}

func NewEventTrace(options ...option) *EventTrace {
	opts := newopts()
	for _, o := range options {
		o(opts)
	}

	return &EventTrace{opts: opts}
}

func (e *EventTrace) Start() error {
	errno := e.StartTrace()
	if errno != types.ULONG(types.ERROR_SUCCESS) {
		if errno == types.ULONG(types.ERROR_ALREADY_EXISTS) {
			errno = evntrace.ControlTrace(0, evntrace.KernelLoggerName, e.props, types.ULONG(evntrace.EventTraceControlQuery))
			if errno != types.ULONG(types.ERROR_SUCCESS) {
				return fmt.Errorf("control query trace error: %d", errno)
			}

			errno = evntrace.ControlTrace(0, evntrace.KernelLoggerName, e.props, types.ULONG(evntrace.EventTraceControlStop))
			if errno != types.ULONG(types.ERROR_SUCCESS) {
				return fmt.Errorf("control stop trace error: %d", errno)
			}

			time.Sleep(time.Millisecond * 100)
			if errno := e.StartTrace(); errno != types.ULONG(types.ERROR_SUCCESS) {
				return fmt.Errorf("start trace error: %d", errno)
			}
		} else {
			return fmt.Errorf("start trace error: %d", errno)
		}
	}

	handle := e.handle
	sysFlags := make([]uint32, 8)
	errno = evntrace.TraceSetInformation(handle, evntrace.TraceSystemTraceEnableFlagsInfo, sysFlags, types.ULONG(len(sysFlags)))
	if errno == types.ULONG(types.ERROR_SUCCESS) {
		sysFlags[0] = uint32(e.opts.flags)
		sysFlags[4] = uint32(evntrace.EventTraceFlagHandle)
		errno = evntrace.TraceSetInformation(handle, evntrace.TraceSystemTraceEnableFlagsInfo, sysFlags, types.ULONG(len(sysFlags)))
		if errno == types.ULONG(types.ERROR_SUCCESS) {
			return nil
		}
	}
	fmt.Println("set trace infotmation error", errno)

	return nil
}

func (e *EventTrace) StartTrace() types.ULONG {
	bufferSize := maxBufferSize
	minBuffers := uint32(runtime.NumCPU() * 2)
	maxBuffers := minBuffers + 20
	flushTimer := time.Second

	e.props = nil
	e.props = &evntrace.EventTraceProperties{
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
		EnableFlags:       e.opts.flags,
		LogFileNameOffset: 0,
		LoggerNameOffset:  types.ULONG(unsafe.Sizeof(evntrace.EventTraceProperties{})),
	}

	return evntrace.StartTrace(&e.handle, evntrace.KernelLoggerName, e.props)
}

func (e *EventTrace) Process() error {
	logname, err := syscall.UTF16PtrFromString(evntrace.KernelLoggerName)
	if err != nil {
		return err
	}

	logFile := evntrace.EventTraceLogFile{
		LoggerName:     logname,
		LogFileMode:    evntrace.EventTraceRealTimeMode | evntrace.EventTraceNoPerProcessorBuffering,
		BufferCallback: syscall.NewCallback(e.opts.bufferCallback),
		EventCallback:  syscall.NewCallback(e.opts.eventCallback),
	}

	handle := evntrace.OpenTrace(&logFile)

	go func() {
		errno := evntrace.ProcessTrace(&handle, 1, nil, nil)
		if errno != types.ULONG(types.ERROR_SUCCESS) {
			return
		}
	}()

	return nil
}

func (e *EventTrace) Close() error {
	return nil
}
