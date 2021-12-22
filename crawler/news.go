package crawler

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/mozyy/empty-news/proto/pbnews"
)

// News is
func News() (items []*pbnews.NewsItem, err error) {
	colletcor := colly.NewCollector(
	// colly.AllowedDomains("cnbeta.com"),
	)
	colletcor.OnHTML("body", func(body *colly.HTMLElement) {
		list := make([]*pbnews.NewsItem, 0)
		body.ForEach(".titleAnc", func(_ int, section *colly.HTMLElement) {
			typeStr := section.Attr("id")
			section.ForEach("li", func(_ int, li *colly.HTMLElement) {
				if strings.HasPrefix(li.ChildAttr(".txt_area a", "href"), "/") {
					view, _ := strconv.Atoi(li.ChildText(".ico_view"))
					comment, _ := strconv.Atoi(li.ChildText(".ico_comment"))
					list = append(list, &pbnews.NewsItem{
						Link:    li.ChildAttr(".txt_area a", "href")[6:13], // /view/1217435.htm => 1217435
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
			err = errors.New("not fund news")
		} else {
			items = list
		}
	})
	colletcor.OnError(func(res *colly.Response, e error) {
		err = e
	})
	colletcor.Visit("https://m.cnbeta.com")
	return
}
