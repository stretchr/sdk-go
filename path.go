package stretchr

import (
	"github.com/stretchrcom/sdk-go/common"
	stewstrings "github.com/stretchrcom/stew/strings"
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
