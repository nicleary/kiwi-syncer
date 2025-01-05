package kiwix_library

import (
	"fmt"
)

func GetZimByName(name string) (*Entry, error) {
	allZims, err := GetAvailableZims(0, 1000, "eng")

	if err != nil {
		fmt.Printf("GetAvailableZims error: %v\n", err)
		return nil, err
	}

	for _, zim := range allZims.Entries {
		if zim.Name == name {
			return &zim, nil
		}
	}

	return nil, nil
}
