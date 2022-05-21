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
)

type Controller struct {
	producer *stream.Producer
	consumer *stream.Consumer
}

func NewController(config *config.Config) *Controller {
	return &Controller{
		producer: stream.NewProducer(config),
		consumer: stream.NewConsumer(),
	}
}

func (c *Controller) Start() error {
	if err := c.producer.Start(); err != nil {
		return err
	}

	return nil
}
