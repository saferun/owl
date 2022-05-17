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
	"github.com/saferun/owl/internal/app/stream"
	"github.com/saferun/owl/internal/config"
	"github.com/saferun/owl/pkg/etw"
)

type Controller struct {
	config   *config.Config
	producer *stream.Producer
	consumer *stream.Consumer
	etw      *etw.EventTrace
}

func NewController(config *config.Config) *Controller {
	c := &Controller{config: config}
	c.producer = stream.NewProducer()
	c.consumer = stream.NewConsumer()
	c.etw = etw.NewEventTrace(
		etw.WithProcess(config.Etw.Process.Enabled),
		etw.WithBufferCallback(c.producer.BufferStatsCallback),
		etw.WithEventCallback(c.producer.ProcessEventCallback),
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
