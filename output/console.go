package output

import (
	"fmt"
	"github-fetch/util"

	"github.com/google/go-github/github"
)

type Console struct {
	query util.Query
}

func (c Console) Do(v github.Repository, step int, pager util.Pager) {
	util.Log(util.LOG_TYPE_INFO,
		fmt.Sprintf("[total: %d][pages: %d][perPage: %d][pos: %d]\tclone url: %s",
			pager.TotalNum, pager.LastPage, c.query.PerPage, step, *v.CloneURL), false)
}
