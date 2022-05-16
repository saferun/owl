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
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Address string `mapstructure:"address"`
	} `mapstructure:"server"`
	Etw struct {
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
	} `mapstructure:"etw"`
	Yara struct {
		Enabled bool     `mapstructure:"enable"`
		Rules   []string `mapstructure:"rules"`
	} `mapstructure:"yara"`
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
