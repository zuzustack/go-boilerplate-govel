package route

import (
	"boilerplate/controllers"
	route "boilerplate/utils/route"
)

func GetRoute()  {
	// Your method and url
	route.Get("/", controllers.GetIndex)
}