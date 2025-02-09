package main

import (
    "github.com/ikeikeikeike/go-sitemap-generator/stm"
    "log"
    "os"
    "net/http"
    "golang.org/x/net/html"
)

// Function to get all URLs from a web page
func getURLs(baseURL string) ([]string, error) {
    var urls []string
    resp, err := http.Get(baseURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    z := html.NewTokenizer(resp.Body)
    for {
        tt := z.Next()
        switch tt {
        case html.ErrorToken:
            return urls, nil
        case html.StartTagToken, html.SelfClosingTagToken:
            t := z.Token()
            if t.Data == "a" {
                for _, a := range t.Attr {
                    if a.Key == "href" {
                        urls = append(urls, a.Val)
                        break
                    }
                }
            }
        }
    }
}

func main() {
    baseURL := "https://imaba.web.id"
    urls, err := getURLs(baseURL)
    if err != nil {
        log.Fatal(err)
    }

    sm := stm.NewSitemap(0)
    sm.Create()

    for _, url := range urls {
        sm.Add(stm.URL{"loc": baseURL + url, "changefreq": "daily"})
    }

    out, err := os.Create("sitemap.xml")
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    sm.WriteTo(out)
}