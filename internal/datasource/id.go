package datasource

import (
	"strconv"
	"time"
)

// AlwaysUniqueID generates an ID, which is always different
func AlwaysUniqueID() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
