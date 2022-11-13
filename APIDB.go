package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	getApi()
}

func getApi() {

	url := "https://api.jikan.moe/v4/anime/40591/full"

	response, err := http.Get(url)

	type Anime struct {
		Data struct {
			Id       int     `json:"mal_id"`
			Type     string  `json:"type"`
			Episodes int     `json:"episodes"`
			Duration string  `json:"duration"`
			Rating   string  `json:"rating"`
			Score    float64 `json:"score"`
			Rank     int     `json:"rank"`

			Year      int    `json:"year"`
			Season    string `json:"season"`
			Licensors []struct {
				MalID int    `json:"mal_id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				URL   string `json:"url"`
			} `json:"licensors"`
			TitleEnglish string `json:"title_english"`
			Title        string `json:"title"`

			Producers []struct {
				MalID int    `json:"mal_id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				URL   string `json:"url"`
			} `json:"producers"`

			Genres []struct {
				MalID int `json:"mal_id"`
			} `json:"genres"`

			Themes []struct {
				MalID int    `json:"mal_id"`
				Type  string `json:"type"`
				Name  string `json:"name"`
				URL   string `json:"url"`
			} `json:"themes"`

			Relations []struct {
				Relation string `json:"relation"`
				Entry    []struct {
					MalID int `json:"mal_id"`

					Name string `json:"name"`
				} `json:"entry"`
			} `json:"relations"`
		} `json:"data"`
	}

	var deneme Anime

	if err != nil {
		fmt.Println(err)
	} else {
		data, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println("deneme hatası", err)
		}
		json.Unmarshal([]byte(data), &deneme)
		fmt.Println(deneme.Data)
	}
}
