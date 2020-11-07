package habr

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
	"log"
	"newsparser/source/models"
	"strings"
)

func ParseRSSByPageIndex(index int) []models.Post {
	var posts []models.Post

	fp := gofeed.NewParser()
	url := fmt.Sprintf("https://habr.com/ru/rss/all/all/page%d/", index)
	feed, err := fp.ParseURL(url)

	if err != nil {
		panic(err)
	}

	// Посты в rss ленте
	feedItems := feed.Items

	for _, item := range feedItems {

		post := models.Post{
			Title:       item.Title,
			Description: setBlankToLinks(item.Description),
			Link:        item.Link,
			ImgUrl:      "item.Image && item.Image.URL",
			Published:   item.Published,
			Categories:  item.Categories,
		}

		posts = append(posts, post)
	}

	return posts
}

func setBlankToLinks(str string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(str))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").SetAttr("target", "_blank")

	output, err := doc.Html()

	if err != nil {
		panic(err)
	}

	return output
}
