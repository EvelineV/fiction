package lyrics

import (
  "log"
  "net/http"
  "os"
  "strings"

  "github.com/PuerkitoBio/goquery"

  "github.com/EvelineV/fiction/utils"
)

type Client struct {
  HTTPClient *http.Client
  RootURL string
  Prefix string
  ListContainer string
  TextContainer string
}

func (c *Client) relativeLink(link string) bool {
  if strings.HasPrefix(link, c.Prefix) {
    return true
  }
  return false
}

func (c *Client) processLink(index int, element *goquery.Selection) string {
  href, exists := element.Attr("href")
  if exists {
    if c.relativeLink(href) {
      return href
    }
  }
  return ""
}

func (c *Client) getSongListForArtist(artist string) ([]string, error) {
  request, err := http.NewRequest("GET", c.RootURL + c.Prefix + artist, strings.NewReader(""))
  if err != nil {
    log.Fatal(err)
  }

  response, err := c.HTTPClient.Do(request)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()

  document, err := goquery.NewDocumentFromReader(response.Body)
  if err != nil {
    log.Fatal(err)
  }

  var container = document.Find(c.ListContainer)
  var links []string
  links = container.Find("a").Map(c.processLink)
  return utils.FilterEmptyStrings(links), nil
}

func (c *Client) getFullURLs(link string) string {
  return c.RootURL + link
}

func (c *Client) getSongLyrics(url string) string {
  request, err := http.NewRequest("GET", url, strings.NewReader(""))
  if err != nil {
    log.Fatal(err)
  }

  response, err := c.HTTPClient.Do(request)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()
  document, err := goquery.NewDocumentFromReader(response.Body)
  if err != nil {
    log.Fatal(err)
  }
  var text = document.Find(c.TextContainer).Text()
  return text
}

func writeLyricsToFile(artist string, lyrics []string) error {
  var file, err = os.Create(artist + ".txt")
  if err != nil {
    return err
  }
  defer file.Close()
  for _, song := range lyrics {
     file.WriteString(song+"\n")
     file.Sync()
  }
  return nil
}

func GetLyricsForArtist(client *Client, artist string) ([]string, error) {
  var artistName = strings.Replace(strings.TrimSpace(strings.ToLower(artist)), " ", "_", -1)
  var links, err = client.getSongListForArtist(artistName)
  if err != nil {
    return nil, err
  }
  urls := utils.Map(links, client.getFullURLs)
  lyrics := utils.Map(urls, client.getSongLyrics)
  err = writeLyricsToFile(artistName, lyrics)
  if err != nil {
    log.Fatal(err)
  }
  return lyrics, nil
}

