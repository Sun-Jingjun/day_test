package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"net/http"
	"time"
)

type test struct {
	Person
	add int
}


type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/excel", func(ctx *gin.Context) {
		file := xlsx.NewFile()
		sheet,_ :=file.AddSheet("sheet")
		//设置表格头
		row := sheet.AddRow()
		var headers = []string {"row1","row2"}
		for _,header := range headers{
			row.AddCell().Value = header
		}

		//插入数据，这里只是一个例子
		row = sheet.AddRow()
		row.AddCell().Value = "1"
		row.AddCell().Value = "2"

		ctx.Header("Content-Disposition", "attachment")
		ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")


		var buffer bytes.Buffer
		if err := file.Write(&buffer); err != nil {
			return
		}
		r := bytes.NewReader(buffer.Bytes())
		http.ServeContent(ctx.Writer, ctx.Request, "test", time.Now(), r)
	})
	r.Run()
}
