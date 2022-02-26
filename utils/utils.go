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
	value, err := url.QueryUnescape(str)
	if nil != err {
		return nil, err
	}

	return strings.Split(value, "&"), nil

}
