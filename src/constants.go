package src

import "time"

const (
	N_LEAGUES = 1
	LEAGUES_URL = "https://www.transfermarkt.co.uk/wettbewerbe/europa/wettbewerbe"
	BASE_URL = "https://www.transfermarkt.co.uk"
	DELAY_BETWEEN_QUERIES = time.Second / 2
)

var (
	PlayerDetailColumns = map[string] bool {"Age:": true, "Height:": true,"Nationality:": true, "Position:": true, "Foot:": true,}
)
