package main

import "fmt"

type Vertex6 struct {
	Lat, Long float64
}

var m map[string]Vertex6

func main() {
	str := "Bell Labs"
	m = make(map[string]Vertex6)
	m[str] = Vertex6{
		Lat:  40.22323,
		Long: -93.22,
	}
	fmt.Println(m[str])
}
