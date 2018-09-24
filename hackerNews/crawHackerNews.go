package hackerNews

import (
	"bytes"
	"encoding/json"
	"fmt"
	html2 "golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

func getAllArticleIDs() []int32 {
	res, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")

	if err != nil {
		log.Printf("go error %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("go error %v", err)
	}
	ids := make([]int32, 400)
	json.NewDecoder(res.Body).Decode(&ids)
	return ids

}

func getOneArticle(id int32) article {
	link := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", id)
	res, err := http.Get(link)

	if err != nil {
		log.Printf("go error %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("go error %v", err)
	}
	news := article{}
	json.NewDecoder(res.Body).Decode(&news)
	return news

}

func getFromHackerNews() []article {
	ids := getAllArticleIDs()
	top10Ids := ids[0:10]
	list := make([]article, 0)

	for _, e := range top10Ids {
		oneArticle := getOneArticle(e)
		list = append(list, oneArticle)
	}

	return list
}

func GetMarkDownFromHackNews() string {

	news := getFromHackerNews()
	var buffer bytes.Buffer

	for i, v := range news {
		if v.Title != "" {
			replacer := strings.NewReplacer("[", "", "]", "")
			sprintf := fmt.Sprintf(`
				%d.%s
				[detail](%s)`,
				i+1, replacer.Replace(v.Title),v.Url)
			buffer.WriteString(html2.EscapeString(sprintf))
		}
	}
	return buffer.String()
}
