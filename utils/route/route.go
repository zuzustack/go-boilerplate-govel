package utils

import (
	"net/http"
)

func Get(url string, callback func(http.ResponseWriter, *http.Request))  {
	http.HandleFunc(url ,callback)
}