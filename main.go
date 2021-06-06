package main

import (
	"github.com/z24tao/sapphire_core/agent"
	_ "github.com/z24tao/sapphire_core/server"
	_ "github.com/z24tao/sapphire_core/world"
	"time"
)

func main() {
	a := agent.NewAgent()
	for {
		a.TimeStep()
		time.Sleep(time.Second / 200)
	}
}
