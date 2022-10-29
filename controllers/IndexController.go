package controllers

import (
	view "boilerplate/utils/view"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
    var data = map[string]interface{}{
        "title": "Boilerplate by zuzustack",
        "name":  "Hello World",
    }

    view.Render("index.html", data, w)
}