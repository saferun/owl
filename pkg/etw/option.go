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
	"github.com/mel2oo/win32/advapi32/evntrace"
	"github.com/mel2oo/win32/tdh"
	"github.com/sirupsen/logrus"
)

type option func(*options)

type options struct {
	flags          evntrace.EventEnableFlags
	bufferCallback BufferStatsCallback
	eventCallback  ProcessEventCallback
}

func newopts() *options {
	return &options{}
}

func WithProcess(enabled bool) option {
	return func(o *options) {
		logrus.Infof("set process event flag: %v", enabled)

		if enabled {
			o.flags |= evntrace.EventTraceFlagProcess
		}
	}
}

func WithThread(enabled bool) option {
	return func(o *options) {
		logrus.Infof("set thread event flag: %v", enabled)

		if enabled {
			o.flags |= evntrace.EventTraceFlagThread
		}
	}
}

func WithImage(enabled bool) option {
	return func(o *options) {
		logrus.Infof("set image event flag: %v", enabled)

		if enabled {
			o.flags |= evntrace.EventTraceFlagImageLoad
		}
	}
}

func WithFile(enabled bool) option {
	return func(o *options) {
		logrus.Infof("set file event flag: %v", enabled)

		if enabled {
			o.flags |= evntrace.EventTraceFlagFileIO
			o.flags |= evntrace.EventTraceFlagFileIOInit
		}
	}
}

func WithRegistry(enabled bool) option {
	return func(o *options) {
		logrus.Infof("set registry event flag: %v", enabled)

		if enabled {
			o.flags |= evntrace.EventTraceFlagRegistry
		}
	}
}

func WithTcpIP(enabled bool) option {
	return func(o *options) {
		logrus.Infof("set tcp/ip event flag: %v", enabled)

		if enabled {
			o.flags |= evntrace.EventTraceFlagNetworkTCPIP
		}
	}
}

func WithDiskIO(enabled bool) option {
	return func(o *options) {
		if enabled {
			o.flags |= evntrace.EventTraceFlagDiskFileIO
			o.flags |= evntrace.EventTraceFlagFileIO
			o.flags |= evntrace.EventTraceFlagFileIOInit
		}
	}
}

type (
	BufferStatsCallback  func(*evntrace.EventTraceLogFile) uintptr
	ProcessEventCallback func(*tdh.EventRecord) uintptr
)

func WithBufferCallback(fn BufferStatsCallback) option {
	return func(o *options) {
		o.bufferCallback = fn
	}
}

func WithEventCallback(fn ProcessEventCallback) option {
	return func(o *options) {
		o.eventCallback = fn
	}
}
