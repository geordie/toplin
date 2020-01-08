package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/geordie/toplin/joplin"

	"github.com/geordie/toplin/libhttp"
)

// HomeHandler function
func HomeHandler(token string, port int) http.Handler {

	sURL := fmt.Sprintf("http://localhost:%v/notes?token=%v", port, url.QueryEscape(token))

	fn := func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(sURL)
		if err != nil {
			libhttp.HandleErrorJson(w, err)
			return
		}

		w.Header().Set("Content-Type", "text/html")

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			libhttp.HandleErrorJson(w, err)
			return
		}

		var todoItems []joplin.TodoItem

		err = json.Unmarshal([]byte(bodyBytes), &todoItems)
		if err != nil {
			fmt.Println("error:", err)
		}

		tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
		if err != nil {
			libhttp.HandleErrorJson(w, err)
			return
		}

		tmpl.Execute(w, todoItems)
	}
	return http.HandlerFunc(fn)
}
