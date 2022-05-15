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
	"github.com/saferun/owl/pkg/etw"
	"github.com/saferun/owl/pkg/stream"
)

type Controller struct {
	etw    *etw.EventTrace
	stream *stream.Consumer
}

func NewController() *Controller {
	return &Controller{
		etw:    etw.NewEventTrace(),
		stream: stream.NewConsumer(),
	}
}

func (c *Controller) Start() error {
	return nil
}

func (c *Controller) Process() error {
	return nil
}
