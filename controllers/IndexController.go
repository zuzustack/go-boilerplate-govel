package controllers

import (
	view "zuzuse/utils/view"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
    var data = map[string]interface{}{
        "title": "Boilerplate by zuzustack",
        "name":  "Hello World",
    }

    view.Render("index.html", data, w)
}