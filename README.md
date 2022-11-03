## github-fetch 小工具
    方便拉取 github 仓库的数据的一个小工具


## 使用
```
# 按指定日期的文件命名保存到文件
github-fetch cli --language Object-C --created 20220101 --out line
# 直接输出到控制台
github-fetch cli --language Object-C --created 20220101 --out console
# 不指定日期
github-fetch cli --language Object-C --out console
```

## 模式
### CLI
```
Usage:
  github-fetch cli [flags]

Flags:
      --created string    指定日期，默认不指定时间 (default "-1")
  -h, --help              help for cli
      --keyword string    按关键字搜索 https://docs.github.com/cn/search-github/searching-on-github/searching-for-repositories (default "-1")
      --language string   需要制定编程语言的选项 (default "-1")
      --order string      数据排序，默认 desc (default "desc")
      --out string        输出方式; 支持 console (直接输出到控制台) / json (按天保存json格式的文件) (default "console")
      --page int          分页的第几页，默认 1 (default 1)
      --perPage int       分页的每页大小，默认 100 (default 100)
      --sort string       数据的类别排序方式，默认 stars (default "stars")
      --token string      需要github 认证 token 值
```

### Serve
todo

## 建议
如果要取完整的数据，则建议使用按天方式获取，将每天的数据存入文件中
> 可参考 [linux bash](./range.sh)

## 开发
```
go run main.go cli --language Objective-C --token 906a910dxxxxxxxxxxxxxx8882fd988
```

## 编译
```
make
```