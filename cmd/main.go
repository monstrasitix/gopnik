package main

import (
	"github.com/monstrasitix/gopnik/core"
	"github.com/monstrasitix/gopnik/env"
)

func init() {
	env.Configure()
}

func main() {
	server := core.NewServer(
		env.PORT(),
		core.Routing(),
	)

	core.Run(server)
}
