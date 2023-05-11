package routes

import (
	"encoding/json"
	"net/http"
)

func HomeRoute() *Routes {
	routes := new(Routes)
	routes.Initalize()

	routes.SetRoute("/", func(res http.ResponseWriter, req *http.Request) {
		json.NewEncoder(res).Encode(map[string]string{
			"hello": "world",
		})
	})

	return routes
}
