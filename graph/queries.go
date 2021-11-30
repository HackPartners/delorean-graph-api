package graph

import (
	"log"

	"github.com/northwesternmutual/grammes"
	. "github.com/northwesternmutual/grammes/model"
	"github.com/northwesternmutual/grammes/query/traversal"
)

const STATION_NAME_KEY = "d0"
const UP_KEY = "d1"
const DOWN_KEY = "d2"
const STATION_LABEL = "node"

func execute(query traversal.String) [][]byte {
	result, err := client.ExecuteQuery(query)
	if err != nil {
		log.Fatalf("Querying error: %s\n", err.Error())
	}
	return result
}

func GetAllStations() []station {
	query := g.V()
	result := execute(query)
	vertList, err := UnmarshalVertexList(result)
	if err != nil {
		log.Fatalf("Unmarshalling error: %s\n", err.Error())
	}
	return VerticesToStations(vertList)
}

func GetStationsAfter(direction string, station string) []station {
	var filterKey string
	if direction == "UP" {
		filterKey = UP_KEY
	} else if direction == "DOWN" {
		filterKey = DOWN_KEY
	}
	query := g.V().
		Has(STATION_NAME_KEY, station).
		Repeat(
			grammes.Traversal().
				OutE().
				Has(filterKey, true).
				InV().
				SimplePath(),
		).
		Until(
			grammes.Traversal().
				OutE().
				Has(filterKey, true).
				Count().Is(0).
				Or().Loops().Is(50),
		).
		Path().
		Unfold().
		HasLabel(STATION_LABEL).
		Dedup()
	result := execute(query)
	vertList, err := UnmarshalVertexList(result)
	if err != nil {
		log.Fatalf("Unmarshalling error: %s\n", err.Error())
	}
	return VerticesToStations(vertList)
}

func GetPathsAfter(direction string, station string) [][]station {
	var filterKey string
	if direction == "UP" {
		filterKey = UP_KEY
	} else if direction == "DOWN" {
		filterKey = DOWN_KEY
	}
	query := g.V().
		Has(STATION_NAME_KEY, station).
		Repeat(
			grammes.Traversal().
				OutE().
				Has(filterKey, true).
				InV().
				SimplePath(),
		).
		Until(
			grammes.Traversal().
				OutE().
				Has(filterKey, true).
				Count().Is(0).
				Or().
				Loops().Is(50),
		).
		Path().
		Local(
			grammes.Traversal().
				Unfold().
				HasLabel(STATION_LABEL).
				Fold(),
		)
	result := execute(query)
	pathList, err := UnmarshalVertexCollection(result)
	if err != nil {
		log.Fatalf("Unmarshalling error: %s\n", err.Error())
	}
	return VerticesCollectionToStationsCollection(pathList)
}

func GetPathsBetween(direction string, stationA string, stationB string) [][]station {
	var filterKey string
	if direction == "UP" {
		filterKey = UP_KEY
	} else if direction == "DOWN" {
		filterKey = DOWN_KEY
	}
	query := g.V().
		Has(STATION_NAME_KEY, stationA).
		Repeat(
			grammes.Traversal().
				OutE().
				Has(filterKey, true).
				InV().
				SimplePath(),
		).
		Until(
			grammes.Traversal().
				Has(STATION_NAME_KEY, stationB).
				Or().
				Loops().Is(50),
		).
		Path().
		Local(
			grammes.Traversal().
				Unfold().
				HasLabel(STATION_LABEL).
				Fold(),
		)
	result := execute(query)
	pathList, err := UnmarshalVertexCollection(result)
	if err != nil {
		log.Fatalf("Unmarshalling error: %s\n", err.Error())
	}
	return VerticesCollectionToStationsCollection(pathList)
}
