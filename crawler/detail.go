package crawler

import (
	"github.com/gocolly/colly"
	pb "github.com/mozyy/empty-news/proto/news"
)

// Detail get news detail
func Detail(url string) (response *pb.DetailResponse, err error) {
	c := colly.NewCollector(
	// colly.AllowedDomains("cnbeta.com"),
	)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		contents := make([]*pb.DetailContent, 0)
		e.ForEach(".articleCont > p", func(_ int, el *colly.HTMLElement) {
			img := el.ChildAttr("img[src]", "src")
			if img != "" {
				contents = append(contents, &pb.DetailContent{Type: 2, Content: img})
			} else {
				contents = append(contents, &pb.DetailContent{Type: 1, Content: el.Text})
			}
		})
		response = &pb.DetailResponse{
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
	c.Visit(url)
	return
}
