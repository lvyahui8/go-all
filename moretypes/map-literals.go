package main

import "fmt"

type Vertex7 struct {
	Lat, Long float64
}

var m = map[string]Vertex7{
	"bell Labs": {
		Lat:  21,
		Long: 1232,
	},
	"google": {23, 23.0},
	"micro": Vertex7{
		Lat:  122,
		Long: 11,
	},
}

func main() {
	fmt.Println(m)
}
