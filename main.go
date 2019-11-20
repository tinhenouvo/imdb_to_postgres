package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"imdb_to_postgres/lib"
	"imdb_to_postgres/models"
	"log"
	"os"
	"path"
	"path/filepath"
)

// Main data for downloading from IMDB

var (
	downloadList    []string
	dbAdapter       string
	dbConnectionURL string
	config          models.TomlConfig
)

func init() {
	dbAdapter = "postgres"
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	dbConnectionURL = fmt.Sprintf("host=%s port=%v user=%s dbname=%s sslmode=%s password=%s",
		config.DB.Server,
		config.DB.Port,
		config.DB.User,
		config.DB.DbName,
		config.DB.SslMode,
		config.DB.Password)
}

func main() {
	fmt.Printf("Import IMDB dataset.\n")
	defer importAction()
	for _, item := range config.Imdb.Files {
		if _, err := os.Stat(path.Join(item, "tsv.gz")); os.IsNotExist(err) {
			continue
		}
		url := config.Imdb.BaseURL + item + ".tsv.gz"
		downloadList = append(downloadList, url)
	}
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	if len(downloadList) != 0 {
		lib.DownloadFiles(fmt.Sprintf("%v/imdb/", pwd), downloadList)
		lib.DecompressFiles(fmt.Sprintf("%v/imdb/", pwd), downloadList)
	}
	// TODO :  remove gz file

}

func downloadAction() {
	fmt.Println("1111")

}

func importAction() {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	log.Println("import Files to Database")
	lib.ImportName(filepath.Join(fmt.Sprintf("%v/imdb/", pwd), "name.basics.tsv"), dbConnectionURL)
	log.Println("Name Basic Done")
	lib.ImportTitleAkas(filepath.Join(fmt.Sprintf("%v/imdb/", pwd), "title.akas.tsv"), dbConnectionURL)
	log.Println("Title akas Done")
	lib.ImportTitleBasics(filepath.Join(fmt.Sprintf("%v/imdb/", pwd), "title.basics.tsv"), dbConnectionURL)
	log.Println("Title Basics Done")
	lib.ImportTitleCrew(filepath.Join(fmt.Sprintf("%v/imdb/", pwd), "title.crew.tsv"), dbConnectionURL)
	log.Println("Title crew Done")
	lib.ImportTitlePrincipals(filepath.Join(fmt.Sprintf("%v/imdb/", pwd), "title.principals.tsv"), dbConnectionURL)
	log.Println("Title principals Done")
	lib.ImportTitleRatings(filepath.Join(fmt.Sprintf("%v/imdb/", pwd), "title.ratings.tsv"), dbConnectionURL)
	log.Println("Title ratings Done")
	lib.ImportTitleEpisodes(filepath.Join(fmt.Sprintf("%v/imdb/", pwd), "title.episode.tsv"), dbConnectionURL)
	log.Println("Title episode Done")
	lib.SanityzeDb(dbConnectionURL)
}
