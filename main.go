package main

import (
	"app/bootstrap"
	"app/router"
)

func main() {
	bootstrap.BootStrap()
	bootstrap.PreLoad()

	r := router.GetServer()

	r.Run(":9501")
}
