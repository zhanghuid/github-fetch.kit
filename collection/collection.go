package collection

import (
	"context"
	"fmt"
	"github-fetch/util"
	"time"

	"github.com/golang-module/carbon"
	"github.com/google/go-github/github"
	"github.com/gosuri/uilive"
	"golang.org/x/oauth2"
)

var ctx = context.Background()

type Collection struct {
	query  util.Query
	client *github.Client
}

func IsWithCreated(created string) bool {
	return created != "-1"
}

func http(c *Collection) (*github.RepositoriesSearchResult, *github.Response) {
	queryString := fmt.Sprintf("language:%s", c.query.Language)
	if IsWithCreated(c.query.Created) {
		createdString := carbon.Parse(c.query.Created).Format("Y-m-d")
		queryString = fmt.Sprintf("language:%s created:%s..%s", c.query.Language, createdString, createdString)
	}

	result, repos, err := c.client.Search.Repositories(ctx,
		queryString,
		&github.SearchOptions{
			Sort:        c.query.Sort,
			Order:       c.query.Order,
			ListOptions: github.ListOptions{Page: c.query.Page, PerPage: c.query.PerPage},
		})
	defer repos.Body.Close()

	// 添加重试
	if _, ok := err.(*github.RateLimitError); ok {
		util.Log(util.LOG_TYPE_DANGER, "hit rate limit", false)
		writer := uilive.New()
		// start listening for updates and render
		writer.Start()

		for i := 0; i <= 60; i++ {
			fmt.Fprintf(writer, "trying...,by the 30 requests per minute. %ds\n", i)
			time.Sleep(time.Second)
		}

		util.Log(util.LOG_TYPE_INFO, "success~", false)
		writer.Stop() // flush and stop rendering
		return http(c)
	}
	// fmt.Printf("result: %v\n", result)
	// fmt.Printf("repos: %v\n", repos.LastPage)
	// os.Exit(333)

	if err != nil {
		util.Log(util.LOG_TYPE_DANGER, err.Error(), true)
	}
	return result, repos
}

func New(i util.Query) *Collection {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: i.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &Collection{query: i, client: github.NewClient(tc)}
}

func (c *Collection) Pager() *github.Response {
	_, resp := http(c)
	return resp
}

func (c *Collection) Get() *github.RepositoriesSearchResult {
	result, _ := http(c)
	return result
}
