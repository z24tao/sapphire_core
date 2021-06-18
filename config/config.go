package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

var Cfg Configuration

type Configuration struct {
	Debug                 bool `env:"DEBUG_" envDefault:"false"`
	DebugMindChanges      bool `env:"DEBUG_MIND_CHANGES" envDefault:"false"`
	DebugActionHypotheses bool `env:"DEBUG_ACTION_HYPOTHESES" envDefault:"false"`
}

func init() {
	err := Cfg.Load()
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Configuration) Load() error {
	return env.Parse(c)
}
