package stretchr

import (
	"github.com/stretchrcom/stew/strings"
	"github.com/stretchrcom/stretchr-sdk-go/common"
)

// Path generates a path from the given arguments.
//
// For example:
//
//     Path("people", person.ID(), "books", book.ID())
func Path(items ...string) string {
	return strings.JoinStrings(common.PathSeparator, items...)
}
