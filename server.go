package main

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	"myapp/handler"
)

func main() {
	e := echo.New()

	db, err := mgo.Dial("mongodb://root:1234@localhost:27017")
	if err != nil {
		e.Logger.Fatal(err)
	}
	h := &handler.Handler{DB: db}

	e.GET("/urls", h.FindUrl)
	e.POST("/urls", h.SaveUrl)

	e.Logger.Fatal(e.Start(":8080"))
}
