package main


import ("fmt"
		"net/http"
		"html/template"
		"io/ioutil"
		"encoding/xml")


type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}
type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}
type NewsMap struct {
	Keyword string
	Location string
}

type NewsAggPage struct {
	Title string
	News map[string] NewsMap
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)

	fmt.Println("");

	// get info from internet
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body);
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body);
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}

	}

	p:= NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	t, _ := template.ParseFiles("newsaggTemplate.html")
	fmt.Println(t.Execute(w, p))
}

func main() {
	fmt.Println("Server on port 8000")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
