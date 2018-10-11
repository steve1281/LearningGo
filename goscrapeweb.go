package main

import ("fmt"
	"net/http"
	"io/ioutil")

func main() {
	fmt.Println("");

	// get info from internet
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body);
	string_body := string(bytes)
	resp.Body.Close()
	fmt.Println(string_body);


}

