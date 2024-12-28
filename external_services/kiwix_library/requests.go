package kiwix_library

import (
	"encoding/xml"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

func GetAvailableZims(startPosition int, count int, lang string) (*Feed, error) {
	client := resty.New()

	resp, err := client.R().SetQueryParams(map[string]string{
		"start": strconv.Itoa(startPosition),
		"count": strconv.Itoa(count),
		"lang":  lang,
	}).Get("https://library.kiwix.org/catalog/v2/entries/")

	if err != nil {
		fmt.Println(fmt.Sprintf("Error while getting available zims: %s", err.Error()))
		return nil, err
	}

	var response Feed

	err = xml.Unmarshal(resp.Body(), &response)

	if err != nil {
		fmt.Println(fmt.Sprintf("Unmarshal Error: %s", err.Error()))
		return nil, err
	}

	return &response, nil
}
