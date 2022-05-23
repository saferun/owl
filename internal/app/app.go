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
	"github.com/sirupsen/logrus"
)

type Controller struct {
	consumer *stream.Consumer
	producer *stream.Producer
}

func NewController(config *config.Config) *Controller {
	consumer := stream.NewConsumer()
	producer := stream.NewProducer(config, consumer)
	return &Controller{
		consumer: consumer,
		producer: producer,
	}
}

func (c *Controller) Start() error {
	logrus.Info("consumer run")
	if err := c.consumer.Run(); err != nil {
		return err
	}

	logrus.Info("producer start")
	if err := c.producer.Start(); err != nil {
		return err
	}

	return nil
}
