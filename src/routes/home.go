package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeRoute() *Routes {
	routes := new(Routes)
	routes.Initalize()

	routes.SetRoute("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("log")
		json.NewEncoder(res).Encode(map[string]string{
			"hello": "world",
		})
	})

	return routes
}
