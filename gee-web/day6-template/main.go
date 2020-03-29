package main

import (
	"fmt"
	"gee"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func formatAsDate(t time.Time) string{
	year, month,day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d",year,month,day)
}


func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"formatAsDate":formatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets","./static")

	stu1 := &student{"Geektutu",20}
	stu2 := &student{"Jack",22}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK,"css.tmpl",nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK,"arr.tmpl",gee.H{
			"title":"gee",
			"stuArr":[2]*student{stu1,stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK,"custom_func.tmpl",gee.H{
			"title":"gee",
			"now":time.Date(2019,2,17,0,0,0,0,time.UTC),
		})
	})
	r.Run(":80")
}
