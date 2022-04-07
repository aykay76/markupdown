package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

func DefaultController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	if strings.HasPrefix(r.URL.Path, "/static") {
		filename := "internal" + r.URL.Path
		body, _ := ioutil.ReadFile(filename)

		if strings.HasSuffix(filename, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}
		w.Write(body)
	} else {
		t, err := template.ParseFiles("internal/templates/index.html", "internal/templates/head.html", "internal/templates/tail.html")
		if err != nil {
			panic(err)
		}

		err = t.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			panic(err)
		}
	}
}
