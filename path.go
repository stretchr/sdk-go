package stretchr

import (
	stewstrings "github.com/stretchrcom/stew/strings"
	"github.com/stretchrcom/stretchr-sdk-go/common"
	"strings"
)

// Path generates a path from the given arguments.
//
// For example:
//
//     Path("people", person.ID(), "books", book.ID())
func Path(items ...string) string {
	return strings.Trim(stewstrings.JoinStrings(common.PathSeparator, items...), common.PathSeparator)
}
