# echo-live

echo middleware that provide livereload feature

## Usage

```go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mattn/echo-live"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	e.Use(middleware.Static("assets"))
	e.Use(live.Live())
	e.Logger.Fatal(e.Start(":8989"))
}
```

## Thanks
echo-live is based off the [echo-livereload](https://github.com/mattn/echo-livereload) library by Yasuhiro Matsumoto (a.k.a. mattn)


