package main

import (
	"fmt"
)

func main() {
	fmt.Println("running")
	// app := iris.New()
	// app.Logger().SetLevel("debug")

	// // Optionally, add two built'n handlers
	// // that can recover from any http-relative panics
	// // and log the requests to the terminal.
	// app.Use(recover.New())
	// app.Use(logger.New())

	// yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
	// 	On:       true,
	// 	DocTitle: "Iris",
	// 	DocPath:  "apidoc.html",
	// 	BaseUrls: map[string]string{"Production": "", "Staging": ""},
	// })
	// app.Use(irisyaag.New()) // <- IMPORTANT, register the middleware.

	// // Method:    GET
	// // Resource:  http://localhost:8080/service
	// app.Get("/service", func(ctx iris.Context) {
	// 	ctx.Writef("service controller")
	// 	application.BusinessService("First call")
	// })

	// // Start the server using a network address.
	// app.Run(iris.Addr(":8080"))
}
