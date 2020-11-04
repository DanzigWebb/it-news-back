package habr

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"newsparser/source/models"
)

func GetPostsByIndex(index int, subject []string) []models.Post {
	url := fmt.Sprintf("https://habr.com/ru/all/page%d/", index)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return ParsePage(string(html), subject)
}
