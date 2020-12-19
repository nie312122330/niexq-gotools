package niexqext

import (
	"strings"

	"github.com/gofrs/uuid"
)

//UUID ...
func UUID() string {
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()
}

// UUIDUperStr ...
func UUIDUperStr() string {
	uidStr := UUID()
	uidStr = strings.Replace(uidStr, "-", "", -1)
	return strings.ToUpper(uidStr)
}
