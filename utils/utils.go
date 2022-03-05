package utils

import (
	"net/url"
	"strings"
)

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
