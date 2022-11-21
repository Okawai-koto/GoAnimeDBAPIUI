package main

func Kontrol(anime Anime) Anime {

	//type
	if anime.Data.Type == "" {
		anime.Data.Type = "Unknown"
	}

	//rating
	if anime.Data.Rating == "" {
		anime.Data.Rating = "Unknown"
	}

	//duration
	if anime.Data.Duration == "" {
		anime.Data.Duration = "Unknown"
	}

	//season
	if anime.Data.Season == "" {
		anime.Data.Season = "Unknown"
	}

	//title
	if anime.Data.Title == "" {
		anime.Data.Title = "Unknown"
	}

	//title eng
	if anime.Data.TitleEnglish == "" {
		anime.Data.TitleEnglish = "Unknown"
	}

	return anime
}
