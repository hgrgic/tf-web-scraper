package src

import (
	"sync"
)

type Team struct {
	TeamName string
	LeagueName string
	Players []Player
}

func monitorTeamWorker(wg *sync.WaitGroup, pc chan Player) {
	wg.Wait()
	close(pc)
}

func (team Team) TeamWorker(lwg *sync.WaitGroup, tc chan Team) {
	twg := &sync.WaitGroup{}
	pc := make(chan Player)

	players := []string{"player1_"+team.TeamName, "player2_"+team.TeamName, "player3_"+team.TeamName}

	for _, player := range players{
		twg.Add(1)
		go PlayerWorker(twg, pc, player)
	}

	go monitorTeamWorker(twg, pc)

	for player := range pc {
		player.CurrentTeam = team.TeamName
		player.CurrentLeague = team.LeagueName
		team.Players = append(team.Players, player)
	}

	tc <- team
	defer lwg.Done()
}
