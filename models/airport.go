package models

type JsonAirportSearchResp struct {
	Data    []JsonAirport `json:"data"`
	Context JsonContext   `json:"context"`
}

type JsonAirport struct {
	Name string `json:"name"`
}

type JsonContext []string

func (c *JsonContext) Add(msg string) {
	*c = append(*c, msg)
}

type JsonUser struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Flights  []string `json:"flights"`
}

type JsonFlight struct {
	Name               string  `json:"name"`
	Flight             string  `json:"flight"`
	Equipment          string  `json:"equipment"`
	Utc                string  `json:"utc"`
	SourceAirport      string  `json:"sourceairport"`
	DestinationAirport string  `json:"destinationairport"`
	Price              float64 `json:"price"`
	FlightTime         int     `json:"flighttime"`
}
