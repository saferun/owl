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
	"fmt"
)

type Consumer struct {
	Queue chan *Event
	Quit  chan struct{}
}

func NewConsumer() *Consumer {
	return &Consumer{
		Queue: make(chan *Event, 1000),
		Quit:  make(chan struct{}),
	}
}

func (c *Consumer) Run() error {

	go func() {
		for {
			select {
			case evt := <-c.Queue:
				fmt.Println(evt.EType.String(), evt.Params)

			case <-c.Quit:
				return
			}
		}
	}()

	return nil
}
