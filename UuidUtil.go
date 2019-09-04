package utils

import (
	"strings"
)

func Get32UUID() string {
	uuids := NewV4()
	id := strings.Replace(uuids.String(), "-", "", -1)
	return id
}
