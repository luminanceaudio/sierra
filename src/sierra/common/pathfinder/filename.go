package pathfinder

import (
	uuid "github.com/google/uuid"
	"strings"
)

func RandomFilename() string {
	uuidStr := uuid.New().String()
	return strings.ReplaceAll(uuidStr, "-", "")
}
