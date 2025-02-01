package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {

	data_slices := make([][]byte, 0)

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		data_slices = append(data_slices, data)
	}

	return data_slices, nil
}
