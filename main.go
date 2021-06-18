package main

import (
	"github.com/z24tao/sapphire_core/agent"
	_ "github.com/z24tao/sapphire_core/server"
	_ "github.com/z24tao/sapphire_core/world"
	"time"
)

// test commit 1
func main() {
	go addAgent(1)
	for {
		time.Sleep(time.Second)
	}
}

func addAgent(num int) {
	var agents []*agent.Agent
	for i := 0; i < num; i++ {

		agents = append(agents, agent.NewAgent())
	}
	for {
		for _, a := range agents {
			a.TimeStep()
			time.Sleep(time.Second / 20)
		}
	}
}
