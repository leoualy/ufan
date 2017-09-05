package server

import (
	"html/template"
	"net/http"
)

func View(w http.ResponseWriter, htmldir string,
	data interface{}) error {
	t, err := template.ParseFiles(htmldir)
	if err != nil {
		return err
	}
	err = t.Execute(w, data)
	return err
}
