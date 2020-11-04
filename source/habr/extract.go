package habr

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"newsparser/source/models"
	"strings"
)

func ParsePage(htmlString string, subject []string) []models.Post {

	var output []models.Post

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlString))

	if err != nil {
		log.Fatal(err)
	}

	posts := doc.Find("li.content-list__item_post")

	posts.Each(func(i int, post *goquery.Selection) {
		flags := post.Find(".inline-list__item_hub")

		flags.EachWithBreak(func(i int, flag *goquery.Selection) bool {
			flagText := strings.TrimSpace(flag.Text())

			if isContainCategory(flagText, subject) {
				updateOutput(post, &output)
				return false
			}
			return true
		})
	})

	return output
}

func updateOutput(post *goquery.Selection, output *[]models.Post) {
	postMeta := post.Find(".post__title_link")
	title := strings.TrimSpace(postMeta.Text())
	description := sliceDescription(strings.TrimSpace(post.Find(".post__text-html").Text()))
	img, _ := post.Find("img").Attr("src")
	link, _ := postMeta.Attr("href")

	postData := models.Post{Title: title, Link: link, Img: img, Description: description}
	*output = append(*output, postData)
}

func isContainCategory(flag string, subject []string) bool {
	lowerFlag := strings.ToLower(flag)

	for _, category := range subject {
		lowerCat := strings.ToLower(category)
		if strings.Contains(lowerFlag, lowerCat) {
			return true
		}
	}

	return false
}

func sliceDescription(text string) string {
	if len(text) > 400 {
		return text[0:399] + "..."
	} else {
		return text
	}
}
