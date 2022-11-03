package util

type Query struct {
	Keyword  string
	Language string
	Sort     string
	Order    string
	PerPage  int
	Page     int
	Created  string
	Token    string
	Out      string
}
