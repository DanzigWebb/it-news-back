package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newsparser/source/models"
)

func mockRoute() {
	http.HandleFunc("/mocks", mockRouterHandler)
}

func mockRouterHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //анализ аргументов,

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	posts := getMock()

	fmt.Fprintf(w, string(posts))
}

func getMock() []byte {
	posts, _ := json.Marshal(createMock())
	return posts
}

func createMock() []models.Post {
	var posts []models.Post

	for i := 1; i < 10; i++ {
		posts = append(posts, models.Post{
			Title: fmt.Sprintf("Title %d", i),
			Link:  fmt.Sprintf("Link to postnumber %d", i)})
	}

	return posts
}
