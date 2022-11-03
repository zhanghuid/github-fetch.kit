/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github-fetch/collection"
	"github-fetch/output"
	"github-fetch/util"

	"github.com/spf13/cobra"
)

var query util.Query

// cliCmd represents the cli command
// rootCmd represents the base command when called without any subcommands
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "拉取 github 仓库地址",
	Long: `支持按 语言/日期/排序/分类获取 github 仓库地址. Eg:
github-fetch cli --language Object-C --keyword "jquery in:name" --created 20220101 --sort stars --order desc --perPage 100 --page 1 --out line|console`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化输出
		outputProvider := output.New(query).Get(query.Out)

		p := collection.New(query).Pager()
		if p.LastPage == 0 {
			p.LastPage = 1
		}

		step := 0

		// 定义分页条数
		for page := 1; page <= p.LastPage; page++ {
			query.Page = page
			result := collection.New(query).Get()
			for _, v := range result.Repositories {
				step++
				outputProvider.Do(v, step, p)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVar(&query.Keyword, "keyword", "-1", "按关键字搜索 https://docs.github.com/cn/search-github/searching-on-github/searching-for-repositories")
	cliCmd.Flags().StringVar(&query.Token, "token", "", "需要github 认证 token 值")
	cliCmd.Flags().StringVar(&query.Language, "language", "-1", "需要制定编程语言的选项")
	cliCmd.Flags().StringVar(&query.Created, "created", "-1", "指定日期，默认不指定时间")
	cliCmd.Flags().StringVar(&query.Sort, "sort", "stars", "数据的类别排序方式，默认 stars")
	cliCmd.Flags().StringVar(&query.Order, "order", "desc", "数据排序，默认 desc")
	cliCmd.Flags().IntVar(&query.PerPage, "perPage", 100, "分页的每页大小，默认 100")
	cliCmd.Flags().IntVar(&query.Page, "page", 1, "分页的第几页，默认 1")
	cliCmd.Flags().StringVar(&query.Out, "out", "console", "输出方式; 支持 console (直接输出到控制台) / json (按天保存json格式的文件)")
	cliCmd.MarkFlagRequired("token")
}
