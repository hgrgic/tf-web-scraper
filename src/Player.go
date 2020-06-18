package src

import (
	"sync"
)

type Player struct {
	PlayerName string
	Age int8
	Height int16
	Position string
	Foot string
	CurrentTeam string
	CurrentLeague string
}


func PlayerWorker(wg *sync.WaitGroup, pc chan Player, playerName string) {
	player := Player{
		PlayerName:  playerName,
		Age:         0,
		Height:      0,
		Position:    "",
		Foot:        "",
		CurrentTeam: "",
		CurrentLeague: "",
	}

	pc <- player
	defer wg.Done()
}