package main


import ("fmt"
		"net/http"
		"html/template"
		"io/ioutil"
		"encoding/xml"
		"sync")

var wg sync.WaitGroup

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

func newsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body);
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex
	news_map := make(map[string]NewsMap)

	fmt.Println("");

	// get info from internet
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body);
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	queue := make(chan News, 100)

	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
		wg.Add(1)
		go newsRoutine(queue, Location)
	}
	wg.Wait()
	close(queue)
	for elem := range queue {
		for idx, _ := range elem.Titles {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
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
