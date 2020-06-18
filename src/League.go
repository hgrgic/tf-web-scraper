package src

import "sync"

type League struct {
	LeagueName string
	Teams []Team
}

func monitorLeagueWorker(wg *sync.WaitGroup, tc chan Team) {
	wg.Wait()
	close(tc)
}

func ScrapeLeague(leagueName string) League{
	league := League{
		LeagueName: leagueName,
		Teams:      []Team{},
	}

	lwg := &sync.WaitGroup{}
	tc := make(chan Team)

	teams := []string{"team1", "team2", "team3"} //TODO: replace with real list of teams

	for _, team := range teams{
		lwg.Add(1)
		t := Team{
			TeamName:   team,
			LeagueName: league.LeagueName,
			Players:    []Player{},
		}
		go t.TeamWorker(lwg, tc)
	}

	go monitorLeagueWorker(lwg, tc)

	for team := range tc {
		league.Teams = append(league.Teams, team)
	}

	return league
}