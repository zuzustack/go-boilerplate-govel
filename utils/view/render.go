package utils

import (
	"html/template"
	"net/http"
	"path"
)

func Render(views string , data any ,w http.ResponseWriter) error  {
    filepath := path.Join("views", views)
    tmpl, err := template.ParseFiles(filepath)

    if err != nil {
        return err
    }
    
    err = tmpl.Execute(w, data)

    if err != nil {
        return err
    }

	return nil
}