package utils

import (
	"encoding/json"
	"net/url"
	"strings"
)

func StructToPrettyString(structure interface{}) (*string, error) {
	b, err := json.MarshalIndent(structure, "", "  ")
	if err == nil {
		str := string(b)
		return &str, nil
	} else {
		return nil, err
	}
}

func StructToString(structure interface{}) (*string, error) {
	b, err := json.Marshal(structure)
	if err == nil {
		str := string(b)
		return &str, nil
	} else {
		return nil, err
	}
}

func DecodeStringToParams(str string) ([]string, error) {

	var final []string

	for _, s := range strings.Split(str, "&") {
		value, err := url.QueryUnescape(s)
		if nil != err {
			return nil, err
		}
		final = append(final, value)
	}

	return final, nil
}
