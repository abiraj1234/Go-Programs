package main

import "fmt"

type airlineparts struct {
	id          string
	parts       string
	service     int
	Enginemodel string
	rate        string
}

type chart struct {
	id           string
	engineerName string
	Type         string
	aircarft     airlineparts
}

func Printcharts(c chart) {
	fmt.Println("Chart_id: ", c.id)
	fmt.Println("EngineerName: ", c.engineerName)
	fmt.Println("Type_ :", c.Type)
	fmt.Println("AIRCARFTSDETAILS____________")
	fmt.Println("_id:", c.aircarft.id)
	fmt.Println("_Parts:", c.aircarft.parts)
	fmt.Println("_service:", c.aircarft.service)
	fmt.Println("_Engine_Model:", c.aircarft.Enginemodel)
	fmt.Println("_rate:", c.aircarft.rate)
}

func main() {
	Aircarfts := []chart{
		{
			id:           "SPJ345E1",
			engineerName: "DONALD",
			Type:         "BOEING737",
			aircarft: airlineparts{
				id:          "SPJ345E1",
				parts:       "FULEINJECTORS, FLAPS , HYDRALUICBRAKES,",
				service:     3,
				Enginemodel: "TwinPropeller",
				rate:        "1298USD",
			},
		},

		{
			id:           "INDI345E1",
			engineerName: "JHONNYPINS",
			Type:         "AIRBUS707",
			aircarft: airlineparts{
				id:          "INDI445E2",
				parts:       "TYRES, ROTORS , LANDING GEARS, GENERALS, STATIC TUBES",
				service:     5,
				Enginemodel: "Turbojet",
				rate:        "1112USD",
			},
		},
	}
	for _, a := range Aircarfts {
		Printcharts(a)

	}
}
