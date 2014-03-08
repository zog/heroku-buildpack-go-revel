package controllers

import "github.com/revel/revel"

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	message := "hello, Revel"
	return c.Render(message)
}
