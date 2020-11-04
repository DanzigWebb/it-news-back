package habr

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"newsparser/source/models"
)

func ParseRSSByPageIndex(index int) []models.HabrPost {
	var posts []models.HabrPost

	fp := gofeed.NewParser()
	url := fmt.Sprintf("https://habr.com/ru/rss/all/all/page%d/", index)
	feed, err := fp.ParseURL(url)

	if err != nil {
		panic(err)
	}

	// Посты в rss ленте
	feedItems := feed.Items

	for _, item := range feedItems {
		post := models.HabrPost{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			ImgUrl:      item.Image.URL,
			Published:   item.Published,
			Categories:  item.Categories,
		}

		posts = append(posts, post)
	}

	return posts
}
