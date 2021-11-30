package graph

import (
	"log"

	"github.com/northwesternmutual/grammes"
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

func GetAllStations() []string {
	query := g.V().Values(STATION_NAME_KEY)
	result := execute(query)
	list, err := UnmarshalList(result)
	if err != nil {
		log.Fatalf("Unmarshalling error: %s\n", err.Error())
	}
	return list
}

func GetStationsAfter(direction string, station string) []string {
	var filterKey string
	if direction == "UP" {
		filterKey = UP_KEY
	} else if direction == "DOWN" {
		filterKey = DOWN_KEY
	}
	query := g.V().
		Has(STATION_NAME_KEY, "WIMBLEDON").
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
		Dedup().
		Values(STATION_NAME_KEY)
	result := execute(query)
	list, err := UnmarshalList(result)
	if err != nil {
		log.Fatalf("Unmarshalling error: %s\n", err.Error())
	}
	return list
}

func GetPathsAfter(direction string, station string) [][]string {
	var filterKey string
	if direction == "UP" {
		filterKey = UP_KEY
	} else if direction == "DOWN" {
		filterKey = DOWN_KEY
	}
	query := g.V().
		Has(STATION_NAME_KEY, "WIMBLEDON").
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
				Values(STATION_NAME_KEY).
				Fold(),
		)
	result := execute(query)
	pathList, err := UnmarshalListCollection(result)
	if err != nil {
		log.Fatalf("Unmarshalling error: %s\n", err.Error())
	}
	return pathList
}

func GetPathsBetween(direction string, stationA string, stationB string) [][]string {
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
				Values(STATION_NAME_KEY).
				Fold(),
		)
	result := execute(query)
	pathList, err := UnmarshalListCollection(result)
	if err != nil {
		log.Fatalf("Unmarshalling error: %s\n", err.Error())
	}
	return pathList
}
