package entities

import (
	"sync"
	"tf-scrapper/src"
)

type League struct {
	LeagueUrl string
	Teams     []Team
}

func NewLeague(leagueUrl string) League {
	return League{
		LeagueUrl: leagueUrl,
		Teams:     []Team{},
	}
}

func monitorLeagueWorker(wg *sync.WaitGroup, tc chan Team) {
	wg.Wait()
	close(tc)
}

/*
League scraper function, which takes the relative to base leagueUrl as a parameter.
 */
func LeagueWorker(leagueUrl string) League {
	//Key variables for the execution setup and control.
	league := NewLeague(leagueUrl)
	lwg := &sync.WaitGroup{}
	tc := make(chan Team)

	//Scraping for team urls
	doc := src.ReadUrl(src.BASE_URL + league.LeagueUrl)
	teamsTable := doc.Find("table", "class", "items")
	teams := teamsTable.FindAll("a", "class", "vereinprofil_tooltip")

	for i := 0;  i<len(teams); i+=3 {
		team := teams[i]
		lwg.Add(1)
		t := NewTeam(team.Attrs()["href"], league.LeagueUrl)
		go t.TeamWorker(lwg, tc)
		//fmt.Println(team.Attrs()["href"])
	}

	go monitorLeagueWorker(lwg, tc)

	for team := range tc {
		league.Teams = append(league.Teams, team)
	}

	return league
}