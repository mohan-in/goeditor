package main

import (
	"go/format"
)

func formatSource(content []byte) ([]byte, error) {
	result, err := format.Source(content)
	if err != nil {
		return nil, err
	}
	return result, nil
}
