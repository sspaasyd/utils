package utils

import (
	. "github.com/satori/go.uuid"
	"strings"
)

func Get32UUID() string {
	uuids, _ := NewV4()
	id := strings.Replace(uuids.String(), "-", "", -1)
	return id
}
