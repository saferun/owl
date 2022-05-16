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

package cli

import (
	"github.com/saferun/owl/internal/app"
	"github.com/saferun/owl/internal/config"
	"github.com/saferun/owl/internal/server"
	"github.com/saferun/owl/pkg/logger"
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
			&FilterFlag,
		},
	}
}

func action(ctx *cli.Context) error {
	// load cmdline params
	cfg := loadctx(ctx)

	// load toml config
	toml, err := config.Load(cfg.Config)
	if err != nil {
		return err
	}

	if err = logger.Init(toml.Logger); err != nil {
		return err
	}

	if err := app.NewController(toml).Start(); err != nil {
		return err
	}

	// start http
	if err := server.New().Start(toml.Server.Address); err != nil {
		return err
	}

	return nil
}

type Config struct {
	Config string
	Binary string
	Filter string
}

func loadctx(ctx *cli.Context) *Config {
	return &Config{
		Config: ctx.String("config"),
		Binary: ctx.String("binary"),
		Filter: ctx.String("filter"),
	}
}
