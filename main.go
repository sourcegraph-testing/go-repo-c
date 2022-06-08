package gorepoc

import (
	"fmt"

	gorepod "github.com/sourcegraph-testing/go-repo-d"
)

func FunctionInRepoC(name string, a, b int) string {
	result := gorepod.FunctionInRepoD(a, b)
	return fmt.Sprintf("name-%s-%d", name, result)
}
