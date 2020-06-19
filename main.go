package main

import (
	"fmt"
	"regexp"
	"tf-scrapper/src"
	"tf-scrapper/src/entities"
)


func main() {
	doc := src.ReadUrl(src.LEAGUES_URL)
	leagueTable := doc.Find("table", "class", "items").Find("tbody")
	leagues := leagueTable.FindAll("a")

	re, err := regexp.Compile("startseite/wettbewerb/[A-Z]{2}1") //Filtering elements of interest only
	if err != nil{
		fmt.Println("Error building league filtering regex!")
		return
	}

	selectFirst := src.N_LEAGUES
	for _, league := range leagues{
		if selectFirst == 0{
			break
		}
		match := re.MatchString(league.Attrs()["href"])
		if match {
			selectFirst--;
			entities.LeagueWorker(league.Attrs()["href"])
			//fmt.Println("Link :", league.Attrs()["href"])
		}
	}
}
