package main

import (
	r "zuzuse/route"
	"fmt"
	"net/http"
)

func main()  {
	r.GetRoute()

    
    var address = "localhost:8080"
    fmt.Printf("server started at %s\n", address)
    err := http.ListenAndServe(address, nil)
    if err != nil {
        fmt.Println(err.Error())
    }
}