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

package config

import (
	"github.com/saferun/owl/pkg/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Server Server        `mapstructure:"server"`
	Logger logger.Config `mapstructure:"logger"`
	Etw    ETW           `mapstructure:"etw"`
	Yara   Yara          `mapstructure:"yara"`
}

type Server struct {
	Address string `mapstructure:"address"`
}

type ETW struct {
	Process struct {
		Enabled bool `mapstructure:"enable"`
	} `mapstructure:"process"`
	Thread struct {
		Enabled bool `mapstructure:"enable"`
	} `mapstructure:"thread"`
	Image struct {
		Enabled bool `mapstructure:"enable"`
	} `mapstructure:"image"`
	TcpIP struct {
		Enabled bool `mapstructure:"enable"`
	} `mapstructure:"tcpip"`
	Reg struct {
		Enabled bool `mapstructure:"enable"`
	} `mapstructure:"registry"`
	DiskIO struct {
		Enabled bool `mapstructure:"enable"`
	} `mapstructure:"diskio"`
}

type Yara struct {
	Enabled bool     `mapstructure:"enable"`
	Rules   []string `mapstructure:"rules"`
}

func Load(path string) (v *Config, err error) {
	viper.SetConfigFile(path)

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = viper.Unmarshal(&v); err != nil {
		return nil, err
	}

	return
}
