package habr

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"newsparser/source/models"
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
			Description: item.Description,
			Link:        item.Link,
			ImgUrl:      "item.Image && item.Image.URL",
			Published:   item.Published,
			Categories:  item.Categories,
		}

		posts = append(posts, post)
	}

	return posts
}
