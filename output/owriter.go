package output

import (
	"github-fetch/util"

	"github.com/google/go-github/github"
)

type OWriter interface {
	Do(v github.Repository, step int, pager util.Pager)
}
