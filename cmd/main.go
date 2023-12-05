package main

import (
	_ "github.com/go-sql-driver/mysql"
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
