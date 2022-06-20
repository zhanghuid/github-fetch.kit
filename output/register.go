package output

import (
	"github-fetch/util"
)

type Register struct {
	members map[string]OWriter
	query   util.Query
}

func New(q util.Query) *Register {
	r := &Register{members: make(map[string]OWriter), query: q}
	r.add("console", &Console{query: q})
	r.add("line", &LineOut{query: q})
	// r.add("jsonout", &JsonOut{})
	return r
}

func (r *Register) add(alias string, obj OWriter) {
	r.members[alias] = obj
}

func (r *Register) Get(alias string) OWriter {
	if _, ok := r.members[alias]; ok {
		return r.members[alias]
	}
	return &Console{query: r.query}
}
