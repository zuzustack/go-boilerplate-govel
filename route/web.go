package route

import (
	"zuzuse/controllers"
	route "zuzuse/utils/route"
)

func GetRoute()  {
	// Your method and url
	route.Get("/", controllers.GetIndex)
}