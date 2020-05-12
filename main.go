package main

import (
	"app/router"
)

func main() {
	r := router.GetServer()

	r.Run(":9501")
}
