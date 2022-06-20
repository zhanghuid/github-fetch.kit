#!/bin/bash
start=`date -d $1 +%Y%m%d`
end=`date -d $2 +%Y%m%d`

while [ ${start} -le ${end} ]
do
  # 假设下面有一个脚本需要2016-01-01格式的日期
  echo `date -d "${start}" +%Y-%m-%d`
  /work/tool/github-fetch/github-fetch-0.0.1-linux-amd64 cli  --language Objective-C  --token 906a910d6409307xxxxx88882fd988 --created `date -d "${start}" +%Y-%m-%d` --out line
  start=`date -d "1 day ${start}" +%Y%m%d`	# 日期自增
done