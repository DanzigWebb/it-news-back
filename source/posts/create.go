package posts

import (
	"fmt"
	"newsparser/source/models"
)

func ToTelegram(post models.Post) string {
	title := fmt.Sprintf("<b>%s</b>", post.Title)
	createAt := fmt.Sprintf("<i>%s</i>", post.Published)
	link := fmt.Sprintf("<a href='%s'>На хабр</a>", post.Link)

	return title + "\n" + createAt + "\n" + link
}
