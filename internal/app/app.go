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

package app

import (
	"syscall"

	"github.com/mel2oo/win32/advapi32/evntrace"
	"github.com/mel2oo/win32/tdh"
	"github.com/mel2oo/win32/types"
	"github.com/saferun/owl/internal/app/event"
	"github.com/saferun/owl/internal/config"
	"github.com/saferun/owl/pkg/etw"
	"github.com/saferun/owl/pkg/stream"
)

type Controller struct {
	config *config.Config
	etw    *etw.EventTrace
	stream *stream.Consumer
}

func NewController(config *config.Config) *Controller {
	c := &Controller{config: config}
	c.stream = stream.NewConsumer()
	c.etw = etw.NewEventTrace(
		etw.WithProcess(config.Etw.Process.Enabled),
		etw.WithBufferCallback(c.BufferStatsCallback),
		etw.WithEventCallback(c.ProcessEventCallback),
	)

	return c
}

func (c *Controller) Start() error {
	if err := c.etw.Start(); err != nil {
		return err
	}

	if err := c.etw.Process(); err != nil {
		return err
	}

	return nil
}

const callbackNext = uintptr(1)

func (c *Controller) BufferStatsCallback(*evntrace.EventTraceLogFile) uintptr {
	return callbackNext
}

func (c *Controller) ProcessEventCallback(evt *tdh.EventRecord) uintptr {
	etype := event.Pack(syscall.GUID(evt.EventHeader.ProviderId), evt.EventHeader.EventDescriptor.Opcode)

	if !etype.Exist() {
		return callbackNext
	}

	var einfo tdh.TraceEventInfo
	var size types.ULONG = 4096

	errno := tdh.TdhGetEventInformation(evt, 0, nil, &einfo, &size)
	if errno != types.ERROR_SUCCESS {
		return callbackNext
	}

	return callbackNext
}
