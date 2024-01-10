package main

import (
	"fmt"

	"github.com/iqbalpradipta/micro-feature-with-go/config"
	"github.com/labstack/echo/v4"
)

func main()  {
	e := echo.New()
	cfg := config.DBConfig()
	if cfg != nil {
		fmt.Println("Lapor Error Pak")
	} else {
		fmt.Println("Masok DATABASE NYA")
	}
	e.Logger.Fatal(e.Start(":8000"))
}