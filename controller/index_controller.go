package controller

import (
	"app/kernel"
)

type IndexController struct {
	Response kernel.Response
}

func (this IndexController) Index() {
	this.Response.Success("Hello World.")
}
