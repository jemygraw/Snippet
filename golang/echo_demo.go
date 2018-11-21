package main

import (
	"fmt"
	"io"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/pandora/data/export", func(ctx echo.Context) error {
		for k, _ := range ctx.Request().Header {
			fmt.Println(k, ":", ctx.Request().Header.Get(k))
		}
		io.Copy(os.Stdout, ctx.Request().Body)
		return nil
	})
	e.Start(":9188")
}
