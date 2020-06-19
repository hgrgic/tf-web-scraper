package entities

import (
	"sync"
	"tf-scrapper/src"
)

type Team struct {
	TeamUrl   string
	LeagueUrl string
	Players   []Player
}

func NewTeam(teamUrl string, leagueUrl string) Team {
	return Team{
		TeamUrl:   teamUrl,
		LeagueUrl: leagueUrl,
		Players:   []Player{},
	}
}

func monitorTeamWorker(wg *sync.WaitGroup, pc chan Player) {
	wg.Wait()
	close(pc)
}

func (team Team) TeamWorker(lwg *sync.WaitGroup, tc chan Team) {
	twg := &sync.WaitGroup{}
	pc := make(chan Player)

	//Scraping for team urls
	doc := src.ReadUrl(src.BASE_URL + team.TeamUrl)
	playersTable := doc.Find("table", "class", "items")
	players := playersTable.FindAll("a", "class", "spielprofil_tooltip")

	for i := 0;  i<len(players); i+=2 {
		player := players[i]
		twg.Add(1)
		go PlayerWorker(twg, pc, player.Attrs()["href"])
		//fmt.Println(player.Attrs()["href"])
	}

	go monitorTeamWorker(twg, pc)

	for player := range pc {
		player.CurrentTeam = team.TeamUrl
		player.CurrentLeague = team.LeagueUrl
		team.Players = append(team.Players, player)
	}

	tc <- team
	defer lwg.Done()
}
