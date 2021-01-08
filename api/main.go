package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	apiKey := os.Getenv("DICTIONARY_API_KEY")

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081"},
	}))

	r.GET("/api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/api/translate", func(c *gin.Context) {
		type Translate struct {
			Text string `form:"text"`
			Lang string `form:"lang"`
		}

		var t Translate

		c.BindQuery(&t)

		text := c.Query("text")
		lang := c.Query("lang")
		url := fmt.Sprintf("https://dictionary.yandex.net/api/v1/dicservice.json/lookup?key=%s&lang=%s&text=%s", apiKey, lang, text)

		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		type Message struct {
			Def []struct {
				Text string `json:"text"`
				Tr   []struct {
					Text string `json:"text"`
				} `json:"tr"`
			} `json:"def"`
		}

		dec := json.NewDecoder(strings.NewReader(string(body)))

		var m Message
		err = dec.Decode(&m)
		if err != nil {
			log.Println(err)
		}

		translations := make([]string, 0, 10)

		for _, def := range m.Def {
			for _, tr := range def.Tr {
				translations = append(translations, tr.Text)
			}
		}

		c.JSON(http.StatusOK, translations)
	})

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
