package crawler

import (
	pb "empty/proto/news"
	"errors"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// News is
func News(callback func([]*pb.NewsItem, error)) {
	colletcor := colly.NewCollector(
	// colly.AllowedDomains("cnbeta.com"),
	)
	colletcor.OnHTML("body", func(body *colly.HTMLElement) {
		list := make([]*pb.NewsItem, 0)
		body.ForEach(".titleAnc", func(_ int, section *colly.HTMLElement) {
			typeStr := section.Attr("id")
			section.ForEach("li", func(_ int, li *colly.HTMLElement) {
				if strings.HasPrefix(li.ChildAttr(".txt_area a", "href"), "/") {
					view, _ := strconv.Atoi(li.ChildText(".ico_view"))
					comment, _ := strconv.Atoi(li.ChildText(".ico_comment"))
					list = append(list, &pb.NewsItem{
						Link:    li.Request.AbsoluteURL(li.ChildAttr(".txt_area a", "href")),
						Image:   li.ChildAttr("img", "src"),
						Title:   li.ChildText(".txt_detail"),
						Time:    li.ChildText(".ico_time"),
						View:    int32(view),
						Comment: int32(comment),
						Type:    typeStr,
					})
				}
			})
		})
		if len(list) == 0 {
			callback(list, errors.New("not fund news"))
		} else {
			callback(list, nil)
		}
	})
	colletcor.OnError(func(e *colly.Response, err error) {
		callback(nil, nil)
	})
	colletcor.Visit("https://m.cnbeta.com")
}
