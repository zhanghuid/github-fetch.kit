package output

import (
	"fmt"
	"github-fetch/collection"
	"github-fetch/util"
	"os"

	"github.com/google/go-github/github"
)

type LineOut struct {
	query util.Query
}

// 判断所给路径是否为文件夹
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

var dir = "./logs"

func init() {
	if !exists(dir) {
		os.Mkdir(dir, 0777)
	}
}

func (l LineOut) Do(v github.Repository, step int, pager util.Pager) {

	util.Log(util.LOG_TYPE_INFO,
		fmt.Sprintf("[total: %d][pages:%d][perPage:%d][pos:%d]\tclone url: %s",
			pager.TotalNum, pager.LastPage, l.query.PerPage, step, *v.CloneURL), false)
	// 当没有指定日期时
	filename := l.query.Created
	if !collection.IsWithCreated(l.query.Created) {
		filename = "dump"
	}

	filepath := fmt.Sprintf("%s/%s.txt", dir, filename)

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		util.Log(util.LOG_TYPE_DANGER, "open file error :"+err.Error(), true)
		return
	}

	defer f.Close()
	fmt.Fprintln(f, *v.CloneURL)
	err = f.Close()
	if err != nil {
		util.Log(util.LOG_TYPE_DANGER, "close file error :"+err.Error(), true)
		return
	}
}
