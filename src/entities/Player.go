package entities

import (
	"fmt"
	"strings"
	"sync"
	"tf-scrapper/src"
)

type Player struct {
	PlayerName string
	Age string
	Height string
	Position string
	Foot string
	Nationality string
	CurrentTeam string
	CurrentLeague string
}

func NewPlayer() Player {
	return Player{
		PlayerName:  "",
		Age:         "",
		Height:      "",
		Position:    "",
		Foot:        "",
		CurrentTeam: "",
		CurrentLeague: "",
	}
}

func (p *Player) setPlayerFields(fieldName string, fieldValue string) {
	switch fieldName {
	case "Age:":
		p.Age = fieldValue
	case "Height:":
		p.Height = fieldValue
	case "Position:":
		p.Position = fieldValue
	case "Foot:":
		p.Foot = fieldValue
	case "Nationality:":
		p.Nationality = fieldValue
	default:
		fmt.Println("Field name could not be matched")
	}
}

func PlayerWorker(wg *sync.WaitGroup, pc chan Player, playerUrl string) {
	player := NewPlayer()

	//Scraping for player details
	doc := src.ReadUrl(src.BASE_URL + playerUrl)

	player.PlayerName = doc.Find("div", "class", "dataMain").Find("h1").FullText()

	for _, entry := range doc.Find("table", "class", "auflistung").FindAll("th"){
		key := strings.TrimSpace(entry.FullText())
		_, present := src.PlayerDetailColumns[key]
		if present {
			val := strings.TrimSpace(entry.FindNextSibling().FindNextSibling().FullText())
			player.setPlayerFields(key,val)
			//fmt.Println(key, val)
		}
	}
	pc <- player
	defer wg.Done()
}