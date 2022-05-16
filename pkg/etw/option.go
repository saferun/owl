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

import "github.com/mel2oo/win32/advapi32/evntrace"

type Option func(*Options)

type Options struct {
	flags evntrace.EventEnableFlags
}

func newOpts() *Options {
	return &Options{
		flags: evntrace.EventTraceFlagProcess,
	}
}

func WithProcess(enabled bool) Option {
	return func(o *Options) {
		if enabled {
			o.flags |= evntrace.EventTraceFlagProcess
		}
	}
}

func WithThread(enabled bool) Option {
	return func(o *Options) {
		if enabled {
			o.flags |= evntrace.EventTraceFlagThread
		}
	}
}

func WithImageLoad(enabled bool) Option {
	return func(o *Options) {
		if enabled {
			o.flags |= evntrace.EventTraceFlagImageLoad
		}
	}
}

func WithTcpIP(enabled bool) Option {
	return func(o *Options) {
		if enabled {
			o.flags |= evntrace.EventTraceFlagNetworkTCPIP
		}
	}
}

func WithRegistry(enabled bool) Option {
	return func(o *Options) {
		if enabled {
			o.flags |= evntrace.EventTraceFlagRegistry
		}
	}
}

func WithDiskIO(enabled bool) Option {
	return func(o *Options) {
		if enabled {
			o.flags |= evntrace.EventTraceFlagDiskFileIO
			o.flags |= evntrace.EventTraceFlagFileIO
			o.flags |= evntrace.EventTraceFlagFileIOInit
		}
	}
}
