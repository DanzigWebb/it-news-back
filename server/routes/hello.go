package routes

import (
	"fmt"
	"net/http"
)

func helloRoute() {
	http.HandleFunc("/", helloRouterHandler) // установим роутер
}

func helloRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello to Home Page!")
}
