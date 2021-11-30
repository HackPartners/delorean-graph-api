package graph

import (
	"encoding/json"

	"github.com/northwesternmutual/grammes/gremerror"
)

type gStringList struct {
	Type  string   `json:"@type"`
	Value []string `json:"@value"`
}

type gStringListCollection struct {
	Type  string        `json:"@type"`
	Value []gStringList `json:"@value"`
}

type gPathList struct {
	Type  string  `json:"@type"`
	Value []gPath `json:"@value"`
}

type gPath struct {
	Type  string          `json:"@type"`
	Value gPathDescripter `json:"@value"`
}

type gPathDescripter struct {
	Labels  gSetList    `json:"labels"`
	Objects gStringList `json:"objects"`
}

type gSetList struct {
	Type  string        `json:"@type"`
	Value []gStringList `json:"@value"`
}

// UnmarshalList is a utility to unmarshal a list
// or array of IDs properly.
func UnmarshalList(data [][]byte) ([]string, error) {
	var list []string

	for _, res := range data {
		var listPart gStringList
		if err := json.Unmarshal(res, &listPart); err != nil {
			return nil, gremerror.NewUnmarshalError("UnmarshalList", res, err)
		}

		list = append(list, listPart.Value...)
	}

	return list, nil
}

// UnmarshalList is a utility to unmarshal a list
// or array of IDs properly.
func UnmarshalListCollection(data [][]byte) ([][]string, error) {
	var list []gStringList

	for _, res := range data {
		var listPart gStringListCollection
		if err := json.Unmarshal(res, &listPart); err != nil {
			return nil, gremerror.NewUnmarshalError("UnmarshalListCollection", res, err)
		}
		list = append(list, listPart.Value...)
	}
	var output [][]string
	for _, path := range list {
		output = append(output, path.Value)
	}
	return output, nil
}

// Unmarshalpatths is a utility to unmarshal a list
// or array of IDs properly.
func UnmarshalPaths(data [][]byte) ([]gPath, error) {
	var list []gPath

	for _, res := range data {
		var listPart gPathList
		if err := json.Unmarshal(res, &listPart); err != nil {
			return nil, gremerror.NewUnmarshalError("UnmarshalPaths", res, err)
		}

		list = append(list, listPart.Value...)
	}

	return list, nil
}
