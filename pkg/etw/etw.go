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
	"github.com/sirupsen/logrus"
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
	if errno := e.startTrace(); errno != types.ULONG(types.ERROR_SUCCESS) {
		if errno == types.ULONG(types.ERROR_ALREADY_EXISTS) {
			if errno := evntrace.ControlTrace(
				0, evntrace.KernelLoggerName,
				e.props, types.ULONG(evntrace.EventTraceControlQuery),
			); errno != types.ULONG(types.ERROR_SUCCESS) {
				return fmt.Errorf("control trace query error: %d", errno)
			}

			if errno := evntrace.ControlTrace(
				0, evntrace.KernelLoggerName,
				e.props, types.ULONG(evntrace.EventTraceControlStop),
			); errno != types.ULONG(types.ERROR_SUCCESS) {
				return fmt.Errorf("control trace stop error: %d", errno)
			}

			time.Sleep(time.Millisecond * 100)

			if errno := e.startTrace(); errno != types.ULONG(types.ERROR_SUCCESS) {
				return fmt.Errorf("start trace error: %d", errno)
			}
		} else {
			return fmt.Errorf("start trace error: %d", errno)
		}
	}

	handle := e.handle
	sysFlags := make([]uint32, 8)
	if errno := evntrace.TraceSetInformation(handle, evntrace.TraceSystemTraceEnableFlagsInfo,
		sysFlags, types.ULONG(len(sysFlags))); errno != types.ULONG(types.ERROR_SUCCESS) {
		logrus.Warnf("set trace infotmation error %d", errno)
	}

	sysFlags[0] = uint32(e.opts.flags)
	sysFlags[4] = uint32(evntrace.EventTraceFlagHandle)

	if errno := evntrace.TraceSetInformation(handle, evntrace.TraceSystemTraceEnableFlagsInfo,
		sysFlags, types.ULONG(len(sysFlags))); errno != types.ULONG(types.ERROR_SUCCESS) {
		logrus.Warnf("set trace infotmation error %d", errno)
	}

	return nil
}

func (e *EventTrace) Close() error {
	if e.handle != 0 {
		if errno := evntrace.ControlTrace(e.handle, evntrace.KernelLoggerName,
			e.props, types.ULONG(evntrace.EventTraceControlFlush)); errno != 0 {
			return fmt.Errorf("control trace flush error: %d", errno)
		}

		time.Sleep(time.Millisecond * 100)

		if errno := evntrace.ControlTrace(e.handle, evntrace.KernelLoggerName,
			e.props, types.ULONG(evntrace.EventTraceControlStop)); errno != 0 {
			return fmt.Errorf("control trace stop error: %d", errno)
		}
	}

	return nil
}

func (e *EventTrace) Process() error {
	logname, err := syscall.UTF16PtrFromString(evntrace.KernelLoggerName)
	if err != nil {
		return err
	}

	logFile := evntrace.EventTraceLogFile{
		LoggerName:     logname,
		BufferCallback: syscall.NewCallback(e.opts.bufferCallback),
	}

	mode := evntrace.EventTraceRealTimeMode | evntrace.EventTraceNoPerProcessorBuffering
	cb := syscall.NewCallback(e.opts.eventCallback)

	*(*uint32)(unsafe.Pointer(&logFile.LogFileMode)) = uint32(mode)
	*(*uintptr)(unsafe.Pointer(&logFile.EventCallback)) = cb

	handle := evntrace.OpenTrace(&logFile)
	e.handle = handle

	go func() {
		errno := evntrace.ProcessTrace(&handle, 1, nil, nil)
		if errno != types.ULONG(types.ERROR_SUCCESS) {
			return
		}
	}()

	return nil
}

func (e *EventTrace) startTrace() types.ULONG {
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
