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
	"github.com/saferun/owl/internal/server"
	"github.com/urfave/cli/v2"
)

func runCmd() *cli.Command {
	return &cli.Command{
		Name:    "run",
		Aliases: []string{"start"},
		Usage:   `Start etw for tarcaing windows kernel event and run http server`,
		Action:  action,
		Flags: []cli.Flag{
			&ConfigFlag,
			&BinaryFlag,
		},
	}
}

func action(*cli.Context) error {
	// TODO: query process map

	// TODO: start etw

	// TODO: start http
	if err := server.New().Start(":"); err != nil {
		return err
	}

	return nil
}
