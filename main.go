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
	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "pong",
		})
	})

	//测试从url拿到名字
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK,"Hello, %s",name)
	})

	//测试get参数
	r.GET("/param", func(context *gin.Context) {
		firstname := context.Query("firstname")
		lastname:= context.Query("lastname")
		context.String(http.StatusOK,"Hello %s %s",lastname,firstname)
	})
	r.Run()*/

	r.GET("/excel", func(ctx *gin.Context) {
		file := xlsx.NewFile()
		sheet,_ :=file.AddSheet("sheet")
		//设置表格头
		row := sheet.AddRow()
		var headers = []string {"row1","row2"}
		for _,header := range headers{
			row.AddCell().Value = header
		}
		//写入数据
/*		for i, log := range logs {
			row := sheet.AddRow()
			row.AddCell().Value = strconv.Itoa(i)
			row.AddCell().Value = string(log.Operation)
			row.AddCell().Value = strconv.FormatInt(log.FileId, 10)
		}*/
		row = sheet.AddRow()
		row.AddCell().Value = "1"
		row.AddCell().Value = "2"

		ctx.Header("Content-Disposition", "attachment")
		ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		/*ctx.ResponseWriter.Header().Add("Content-Disposition", "attachment")
		//ctx.ResponseWriter.Header().Add("Content-Type", "application/vnd.ms-excel")
		//xlsx
		ctx.ResponseWriter.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")*/

		var buffer bytes.Buffer
		if err := file.Write(&buffer); err != nil {
			return
		}
		r := bytes.NewReader(buffer.Bytes())
		http.ServeContent(ctx.Writer, ctx.Request, "test", time.Now(), r)
	})
	r.Run()
}
