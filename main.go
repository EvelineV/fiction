package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/EvelineV/fiction/corpora/lyrics"
)

func main() {
  var client = &lyrics.Client{
    RootURL: "",
    Prefix: "",
    HTTPClient: &http.Client{
      Timeout: 30 * time.Second,
    },
    ListContainer: "",
    TextContainer: "",
  }
  lyrics, _ := lyrics.GetLyricsForArtist(client, "")
  fmt.Println(len(lyrics))
}

