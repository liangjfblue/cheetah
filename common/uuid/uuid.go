package uuid

import (
	"strings"

	guuid "github.com/satori/go.uuid"
)

func UUID() string {
	return strings.Replace(guuid.NewV4().String(), "-", "", -1)
}
