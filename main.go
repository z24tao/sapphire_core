package main

import (
	"./agent"
	_ "./server"
	_ "./world"
	"time"
)

func main() {
	a := agent.NewAgent()
	for {
		a.TimeStep()
		time.Sleep(time.Second / 100)
	}
}
