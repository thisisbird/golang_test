* 閱讀文件指令
go doc os OpenFile

* 測試指令

```
go test github.com/headfirstgo/prose
go test github.com/headfirstgo/prose -v
go test github.com/headfirstgo/prose -v -run Two
go test github.com/headfirstgo/prose -v -run Elements
```
*  go run xxx.go  		直接執行
*  go fmt xxx.go		格式化檔案(建議 非必要)
*  go build xxx.go		打包檔案
*  go /.xxx			執行檔案
*  export GOPATH="/Users/zhangkaijun/Desktop/golang/test"		改變路徑
*  export GO111MODULE=off 使用GOPATH
*  每次終端機都要輸入export才會生效