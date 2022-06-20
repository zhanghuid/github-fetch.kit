package output

import (
	"fmt"
	"github-fetch/util"

	"github.com/google/go-github/github"
)

type Console struct {
	query util.Query
}

func (c Console) Do(v github.Repository, step int, p *github.Response) {
	util.Log(util.LOG_TYPE_INFO,
		fmt.Sprintf("[pages: %d][perPage: %d][pos: %d]\tclone url: %s",
			p.LastPage, c.query.PerPage, step, *v.CloneURL), false)
}
