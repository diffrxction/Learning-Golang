package main

import (
	"fmt"		 // To print the data out, we'll use fmt
	"io/ioutil"  // To read the data, import io/ioutil
	"net/http"   // To make the request we'll use net/http
)

func main(){
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	//Next we use resp.Body to access the body of the website and we read this with the help of ioutio.ReadAll
	bytes, _ := ioutil.ReadAll(resp.Body)
	//Next convert to string usinf string()
	stringBody := string(bytes)
	fmt.Println(stringBody)
	//Close to free up resources with .close()
	resp.Body.Close()
}