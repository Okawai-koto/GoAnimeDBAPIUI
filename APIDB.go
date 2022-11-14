package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func tes() {
	tickTime := time.NewTicker(time.Second * 1)
	sayac := 0
	_ = tickTime.C
	for v := range tickTime.C {
		_ = v
		x := getApi(sayac)
		if x.Data.Id != 0 {
			fmt.Println(x.Data.Title)
			AddToDB(columnsAnime(), x)
		}
		sayac = sayac + 1
	}
}

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
			LicensorId int `json:"mal_id"`
		} `json:"licensors"`
		TitleEnglish string `json:"title_english"`
		Title        string `json:"title"`

		Producers []struct {
			ProducersId int `json:"mal_id"`
		} `json:"producers"`

		Genres []struct {
			GenresId int `json:"mal_id"`
		} `json:"genres"`

		Themes []struct {
			ThemesId int `json:"mal_id"`
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

func getApi(num int) Anime {

	url := "https://api.jikan.moe/v4/anime/" + strconv.Itoa(num) + "/full"

	response, err := http.Get(url)

	var deneme Anime

	if err != nil {
		fmt.Println(err)
	} else {
		data, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println("deneme hatası", err)
		}
		json.Unmarshal([]byte(data), &deneme)
		// fmt.Println("Id: ", deneme.Data.Id)
		// fmt.Printf(deneme.Data.Title)

		// for i, v := range deneme.Data.Licensors {
		// 	fmt.Println(i, ". licensor:", v.LicensorId)
		// }

		// for i, v := range deneme.Data.Producers {
		// 	fmt.Println(strconv.Itoa(i) + ". producer:" + strconv.Itoa(v.ProducersId))
		// }

	}
	return deneme
}
