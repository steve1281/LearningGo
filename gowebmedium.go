package main

import ("fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml")


/* example format from website if website is down: */
var washPostXML = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>`)


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

func main() {
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

	for idx, data := range news_map {
		fmt.Println("\n\n\n",idx)
		fmt.Println("\n",data.Keyword)
		fmt.Println("\n",data.Location)
	}
}

