package kiwix_library

import (
	"fmt"
)

func ZimExistsByName(name string) (bool, error) {
	allZims, err := GetAvailableZims(0, 1000, "eng")

	if err != nil {
		fmt.Printf("GetAvailableZims error: %v\n", err)
		return false, err
	}

	for _, zim := range allZims.Entries {
		if zim.Name == name {
			return true, nil
		}
	}

	return false, nil
}
