package output

import (
	"github.com/google/go-github/github"
)

type OWriter interface {
	Do(v github.Repository, step int, p *github.Response)
}
