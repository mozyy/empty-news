package crawler

import (
	"encoding/json"
	"os"

	"github.com/gocolly/colly"
)

type Content struct {
	Type    int    `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}

type NewsDetail struct {
	Title   string    `json:"title,omitempty"`
	From    string    `json:"from,omitempty"`
	Time    string    `json:"time,omitempty"`
	Summary string    `json:"summary,omitempty"`
	Content []Content `json:"content,omitempty"`
}

// Detail get news detail
func Detail(url string) {
	c := colly.NewCollector(
	// colly.AllowedDomains("cnbeta.com"),
	)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		content := make([]Content, 0)
		e.ForEach(".articleCont > p", func(_ int, el *colly.HTMLElement) {
			img := el.ChildAttr("img[src]", "src")
			// if img == "" && el.Text == "" {
			// 	return
			// }
			if img != "" {
				content = append(content, Content{2, img})
			} else {
				content = append(content, Content{1, el.Text})
			}
		})
		detail := NewsDetail{
			e.ChildText(".article-tit"),
			e.ChildText(".article-byline > span"),
			e.ChildText(".article-byline > .time"),
			e.ChildText(".article-summ > p"),
			content,
		}
		// fmt.Println(detail)
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")

		enc.Encode(detail)
	})
	c.Visit(url)
}
