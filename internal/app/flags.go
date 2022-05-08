/*
 * Copyright 2022 by Mel2oo <https://github.com/saferun/monitor-windows>
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

import "github.com/urfave/cli/v2"

var (
	BinaryFlag = cli.StringFlag{
		Name:    "binary",
		Aliases: []string{"bin"},
		Usage:   "Input sample path for running and tracing",
		Value:   "C:\\Windows\\System32\\cmd.exe",
	}
)