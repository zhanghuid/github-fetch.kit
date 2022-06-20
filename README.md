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