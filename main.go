package main

import "app/router"
import "app/kernel"

func main() {
	kernel.BootStrap()

	r := router.GetServer()

	r.Run(":9501")
}
