package crawler

import (
	"github.com/gocolly/colly"
	newsv1 "github.com/mozyy/empty-news/proto/news/news/v1"
)

// Detail get news detail
func Detail(link string) (response *newsv1.DetailResponse, err error) {
	c := colly.NewCollector(
	// colly.AllowedDomains("cnbeta.com"),
	)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		contents := make([]*newsv1.DetailResponse_DetailContent, 0)
		e.ForEach(".articleCont > p", func(_ int, el *colly.HTMLElement) {
			img := el.ChildAttr("img[src]", "src")
			if img != "" {
				contents = append(contents, &newsv1.DetailResponse_DetailContent{Type: 2, Content: img})
			} else {
				contents = append(contents, &newsv1.DetailResponse_DetailContent{Type: 1, Content: el.Text})
			}
		})
		response = &newsv1.DetailResponse{
			Title:   e.ChildText(".article-tit"),
			From:    e.ChildText(".article-byline > span"),
			Time:    e.ChildText(".article-byline > .time"),
			Summary: e.ChildText(".article-summ > p"),
			Content: contents,
		}
	})
	c.OnError(func(res *colly.Response, e error) {
		err = e
	})
	c.Visit("https://m.cnbeta.com/view/" + link + ".htm")
	return
}
