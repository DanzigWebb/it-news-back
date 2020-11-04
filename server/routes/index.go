package routes

import (
	"log"
	"net/http"
)

func Run() {
	helloRoute()
	mockRoute()
	postsRoute()

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
