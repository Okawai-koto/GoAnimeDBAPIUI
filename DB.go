package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

type Anime struct {
	data Data
}

type Data struct {
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
}

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("mssql", "server=localhost; Trusted_Connection=true;Database=GoAnimeDB")
	return
}

type dbColumn struct {
	name       string
	columnType string
}

func columnsAnime() []dbColumn {
	var dbColumns []dbColumn
	dbColumns = append(dbColumns,
		dbColumn{
			name:       "Id",
			columnType: "int",
		},
		dbColumn{
			name:       "Type",
			columnType: "varchar(250)",
		},
		dbColumn{
			name:       "Episodes",
			columnType: "int",
		},
		dbColumn{
			name:       "Rating",
			columnType: "varchar(250)",
		},
		dbColumn{
			name:       "Duration",
			columnType: "varchar(250)",
		},
		dbColumn{
			name:       "Score",
			columnType: "float",
		},
		dbColumn{
			name:       "Rank",
			columnType: "int",
		},
		dbColumn{
			name:       "Year",
			columnType: "varchar(250)",
		},
		dbColumn{
			name:       "Season",
			columnType: "varchar(250)",
		},
		dbColumn{
			name:       "Licensors",
			columnType: "varchar(250)",
		},
		dbColumn{
			name:       "Title",
			columnType: "varchar(250)",
		},
		dbColumn{
			name:       "TitleEnglish",
			columnType: "varchar(250)",
		},
		dbColumn{
			name:       "TitleTurkish",
			columnType: "varchar(250)",
		},
		dbColumn{
			name:       "Producers",
			columnType: "varchar(250)",
		},
		dbColumn{
			name:       "Genres",
			columnType: "int",
		},
	)
	return dbColumns
}

func sqlQuery(queryString string) {
	db, err := GetDB()
	if err != nil {
		fmt.Println("err1", err)

	} else {
		response, err := db.Query(queryString)
		if err != nil {
			fmt.Println("err2", err)
		} else {
			fmt.Println("%s", response)
		}
	}
}

func createTable(columns []dbColumn) {

	sqlQueryString := "CREATE TABLE Anime ("
	for _, v := range columns {
		sqlQueryString += v.name + " " + v.columnType + ", "
	}
	sqlQueryString = sqlQueryString[:len(sqlQueryString)-2]
	sqlQueryString += ");"
	fmt.Println(sqlQueryString)
	sqlQuery(sqlQueryString)

}

func addToDB(columns []dbColumn, anime Anime) {
	sqlQueryString := "INSERT INTO Anime ("
	for _, v := range columns {
		sqlQueryString += v.name + ", "
	}
	sqlQueryString = sqlQueryString[:len(sqlQueryString)-2]
	sqlQueryString += ") "
	score := fmt.Sprintf("%f", anime.data.Score)
	licensor := ""
	for _, v := range anime.data.Licensors {
		licensor += string(v.LicensorId)
	}
	producers := ""
	for _, v := range anime.data.Producers {
		producers += string(v.ProducersId)
	}
	genre := ""
	for _, v := range anime.data.Genres {
		genre += string(v.GenresId)
	}

	sqlQueryString += "VALUES (" +
		strconv.Itoa(anime.data.Id) + ", " +
		anime.data.Type + ", " +
		strconv.Itoa(anime.data.Episodes) + ", " +
		anime.data.Rating + ", " +
		anime.data.Duration + ", " +
		score + ", " +
		strconv.Itoa(anime.data.Year) + ", " +
		anime.data.Season + ", " +
		licensor + ", " +
		anime.data.Title + ", " +
		anime.data.TitleEnglish + ", " +
		"TR" + ", " +
		producers + ", " +
		genre + ", " +
		");"
	fmt.Println(sqlQueryString)
	//sqlQuery(sqlQueryString)

}

func addToDBV1(tableName string, test_structs Anime) {
	/*
		sqlQueryString := "INSERT INTO " + tableName + " ("
		test_s := Anime{}
		reflect := reflect.ValueOf(&test_s.Data).Elem()
		for i := 0; i < reflect.NumField(); i++ {
			sqlQueryString += reflect.Type().Field(i).Name + ",\n "
		}

		sqlQueryString = sqlQueryString[:len(sqlQueryString)-2]
		sqlQueryString += ") ("
		fmt.Println(sqlQueryString)
		//sqlQuery("(CustomerName, ContactName, Address, City, PostalCode, Country) "+
		//"VALUES ('Cardinal', 'Tom B. Erichsen', 'Skagen 21', 'Stavanger', '4006', 'Norway');")
	*/
}

func main() {
	//createTable("Anime", columnsAnime())
	//createTable("Images", columns())
	//createTable("Genres", columns())
	//test_structs := Anime{}
	//addToDB("Test", test_structs)
	/*
		test_s := test_struct{}
		test_s.Data.id = 5
		test_s.Data.Title = "tt"
		test_s.Data.Year = 50
		test := reflect.ValueOf(&test_s).Elem()
		for i := 0; i < test.NumField(); i++ {
			fmt.Println("Name: ", test.Field(i).Interface())
		}
	*/

	data := Data{
		Id:           1,
		Type:         "1",
		Episodes:     12,
		Duration:     "23",
		Rating:       "8321",
		Score:        8.7,
		Rank:         300,
		Year:         2016,
		Season:       "2. sezon",
		Licensors:    nil,
		TitleEnglish: "deneme 12 34",
		Title:        "başlıktır bu",
		Producers:    nil,
		Genres:       nil,
		Themes:       nil,
		Relations:    nil,
	}
	andime := Anime{
		data: data,
	}
	addToDB(columnsAnime(), andime)
}
