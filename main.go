package main

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/upload", func(c echo.Context) error {
		var r request
		if err := c.Bind(&r); err != nil {
			log.Printf("Error ->> %+v\n\n\n", err)
			return err
		}
		return c.String(http.StatusOK, "Is it okay?!")
	})

	e.Logger.Fatal(e.Start(":9999"))
}

type request struct {
	ID       uint                 `form:"id"`
	File     multipart.FileHeader `from:"file"`
	Products []product            `form:"products"`
}

type product struct {
	Name     string  `form:"name"`
	Quantity int     `form:"quantity"`
	Price    float32 `form:"price"`
}
