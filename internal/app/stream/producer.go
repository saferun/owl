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

package stream

import (
	"syscall"

	"github.com/mel2oo/win32/advapi32/evntrace"
	"github.com/mel2oo/win32/tdh"
	"github.com/mel2oo/win32/types"
	"github.com/saferun/owl/internal/app/event"
)

const (
	callbackNext = uintptr(1)
	bufferSize   = 4096
)

type Producer struct {
}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) BufferStatsCallback(*evntrace.EventTraceLogFile) uintptr {
	return callbackNext
}

func (p *Producer) ProcessEventCallback(evt *tdh.EventRecord) uintptr {
	etype := event.Pack(syscall.GUID(evt.EventHeader.ProviderId), evt.EventHeader.EventDescriptor.Opcode)

	if !etype.Exist() {
		return callbackNext
	}

	var einfo tdh.TraceEventInfo
	var size types.ULONG = bufferSize

	errno := tdh.TdhGetEventInformation(evt, 0, nil, &einfo, &size)
	if errno != types.ERROR_SUCCESS {
		return callbackNext
	}

	return callbackNext
}
