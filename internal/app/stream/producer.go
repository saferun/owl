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
	"unsafe"

	"github.com/mel2oo/win32/advapi32/evntrace"
	"github.com/mel2oo/win32/tdh"
	"github.com/mel2oo/win32/types"
	"github.com/saferun/owl/internal/app/event"
	"github.com/saferun/owl/internal/config"
	"github.com/saferun/owl/pkg/etw"
	"github.com/sirupsen/logrus"
)

const (
	callbackNext = uintptr(1)
	bufferSize   = 4096
)

type Producer struct {
	config *config.Config
	etw    *etw.EventTrace
}

func NewProducer(c *config.Config) *Producer {
	return &Producer{config: c}
}

func (p *Producer) Start() error {
	p.etw = etw.NewEventTrace(
		etw.WithProcess(p.config.Etw.Process.Enabled),
		etw.WithThread(p.config.Etw.Thread.Enabled),
		etw.WithImage(p.config.Etw.Image.Enabled),
		etw.WithFile(p.config.Etw.File.Enabled),
		etw.WithRegistry(p.config.Etw.Registry.Enabled),
		etw.WithTcpIP(p.config.Etw.TcpIP.Enabled),
		etw.WithDiskIO(p.config.Etw.DiskIO.Enabled),
		etw.WithBufferCallback(p.BufferStatsCallback),
		etw.WithEventCallback(p.ProcessEventCallback),
	)

	if err := p.etw.Start(); err != nil {
		return err
	}

	if err := p.etw.Process(); err != nil {
		return err
	}

	return nil
}

func (p *Producer) BufferStatsCallback(*evntrace.EventTraceLogFile) uintptr {
	return callbackNext
}

func (p *Producer) ProcessEventCallback(evt *tdh.EventRecord) uintptr {

	etype := event.Pack(syscall.GUID(evt.EventHeader.ProviderId),
		evt.EventHeader.EventDescriptor.Opcode)

	if !etype.Exist() || etype.Dropped() {
		return callbackNext
	}

	var (
		bufferSize types.ULONG = bufferSize
		buffer                 = make([]byte, bufferSize)
		info                   = (*tdh.TraceEventInfo)(unsafe.Pointer(&buffer[0]))
	)

	errno := tdh.TdhGetEventInformation(evt, 0, nil, info, &bufferSize)
	if errno != types.ERROR_SUCCESS {
		return callbackNext
	}

	params := event.Parse(etype, evt, info)

	logrus.Debugf("event:%s, params:%v", etype.String(), params)

	return callbackNext
}
