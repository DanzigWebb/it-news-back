package habr

import (
	"fmt"
	"net/http"
	"net/url"
	"newsparser/source/models"
)

var (
	token   = "<fake_token>"
	chatIDs = []string{"793490251", "693419291"}
)

func SendPostToBot(posts []models.Post) {
	for _, chatID := range chatIDs {
		for _, post := range posts {
			urlBot := createUrl(post, chatID)
			sendItem(urlBot.String())
		}
	}
}

func sendItem(urlToBOt string) {
	resp, err := http.Get(urlToBOt)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

func createUrl(post models.Post, chatID string) *url.URL {
	urlToBot := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	baseUrl, _ := url.Parse(urlToBot)
	params := url.Values{}

	params.Add("parse_mode", "html")
	params.Add("chat_id", chatID)
	params.Add("text", createMessage(post))

	baseUrl.RawQuery = params.Encode()

	return baseUrl
}

func createMessage(post models.Post) string {
	title := fmt.Sprintf("<b>%s</b>\n", post.Title)
	link := fmt.Sprintf("<a href='%s'>На хабр</a>", post.Link)

	return title + link
}
