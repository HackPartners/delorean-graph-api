package graph

import (
	"log"
	"strconv"

	. "github.com/northwesternmutual/grammes/model"
)

type station struct {
	Stanox int    `json:"stanox"`
	Name   string `json:"name"`
}

func VertexToStation(vertex Vertex) station {
	id, converted := vertex.ID().(string)
	stanox, err := strconv.Atoi(id)
	name, converted := vertex.PropertyValue(STATION_NAME_KEY, 0).(string)
	if (!converted) || (err != nil) {
		log.Fatalf("Station conversion error %s\n", err.Error())
	}
	return station{
		Stanox: stanox,
		Name:   name,
	}
}

func VerticesToStations(vertices []Vertex) []station {
	var stations []station
	for _, vertex := range vertices {
		stations = append(stations, VertexToStation(vertex))
	}
	return stations
}

func VerticesCollectionToStationsCollection(verticesCollection [][]Vertex) [][]station {
	var stationsCollection [][]station
	for _, vertices := range verticesCollection {
		stationsCollection = append(stationsCollection, VerticesToStations(vertices))
	}
	return stationsCollection
}
