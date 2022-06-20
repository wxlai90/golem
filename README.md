# golem

<img style="display: block;width:250px;margin: 0 auto;" src="https://static.wikia.nocookie.net/clashofclans/images/c/c2/Golem_info.png/revision/latest?cb=20170927231256" />

Resembles Express for Node, features are implemented on a need basis.

### To install

```sh
go get -u github.com/wxlai90/golem
```

### Quickstart

```go
package main

import (
	"fmt"

	"github.com/wxlai90/golem"
	"github.com/wxlai90/golem/router"
)

const (
	PORT = "5000"
)

type todo struct {
	Description string
	Done        bool
}

func main() {
	app := golem.New()

	app.GET("/", func(req *golem.Request, res *golem.Response) {
		res.Send("Hello World")
	})

	app.GET("/todos", func(req *golem.Request, res *golem.Response) {
		todos := []todo{
			{
				Description: "Buy groceries",
				Done:        false,
			},
			{
				Description: "Take out trash",
				Done:        true,
			},
		}

		res.JSON(todos)
	})

	app.Listen(PORT, func() {
		fmt.Printf("Listening on %s\n", PORT)
	})
}
```

See the responses at:

http://localhost:5000/

http://localhost:5000/todos

## Cookies

### Set Cookie

```go
res.Cookie("name", "value")
```

### Get Cookie

req.Cookies contains a map[string]string of cookies pairs

```go
if value, ok :=req.Cookies("name"); ok {
	// ... do something with value
}
```

## Status Code

### Set Status

```go
res.Status(http.StatusOK).Send("All good")
```

## Middlewares

### Global Middlewares

```go
func main() {
	app := golem.New()

	app.Use(func(req *golem.Request, res *golem.Response, next router.Next) {
		log.Printf("Incoming [%s] request to %s\n", req.R.Method, req.R.URL)
		next()
	})

	app.GET("/", func(req *golem.Request, res *golem.Response) {
		res.Send("Hello World")
	})

	app.Listen(PORT, func() {
		fmt.Printf("Listening on %s\n", PORT)
	})
}
```

Output

```sh
$ Listening on 5000
$ 2022/06/13 20:14:23 Incoming [GET] request to /
```

### Route specific middlewares

Add route specific middlewares by passing in as last parameter. Slightly different from Express but necessary due to go's variadic parameters requirements.

```go
func main() {
	app := golem.New()

	logger := func(req *golem.Request, res *golem.Response, next router.Next) {
		log.Printf("Incoming [%s] request to %s\n", req.R.Method, req.R.URL)
		next()
	}

	app.GET("/", func(req *golem.Request, res *golem.Response) {
		res.Send("Hello World")
	}, logger)

	app.Listen(PORT, func() {
		fmt.Printf("Listening on %s\n", PORT)
	})
}
```

Output

```sh
$ Listening on 5000
$ 2022/06/13 20:17:19 Incoming [GET] request to /
```

## Bag (key, value) passing

Bag exposes a Put() and Get() interface for adding data into the request object after each middleware call.

In order to cater to any type, Get() returns an empty interface

```go
interface{}
```

A type assertion is required for any meaningful usage. However, if we were simply returning it as JSON, then type assertion is optional.

Example

```go
type Todo struct {
	Description string
	Done        bool
}

func main() {
	app := golem.New()

	createTodo := func(req *golem.Request, res *golem.Response, next router.Next) {
		todo := Todo{
			Description: "Do something",
			Done:        false,
		}
		req.Put("todo", todo)
		next()
	}

	app.GET("/", func(req *golem.Request, res *golem.Response) {
		if todo, ok := req.Get("todo"); ok {
			res.JSON(todo)
			return
		}

		res.Send("Couldn't find todo")

	}, createTodo)
}
```

## Request Body

req.Body exposes RawBytes as an option to handle it manually. Optionally, a convenience function Unmarshal() is provided for unmarshaling the request body into a struct.

```go

type Todo struct {
	Description string
	Done        bool
}

func main() {
	app := golem.New()

	app.POST("/todo/new", func(req *golem.Request, res *golem.Response) {
		todo := &Todo{}
		if err := req.Body.Unmarshal(todo); err == nil {
			res.Send("Created todo")
		}
	})
}
```

## Sub-routes

Sub-routing is a useful way to define routes in separate files. There can be as many sub-routers as needed, with a common prefix to indicate a difference resource.

Create a sub-router like so:

```go
var FruitsRouter *golem.SubRouter

func init() {
	FruitsRouter = golem.NewSubRouter()

	FruitsRouter.GET("/", func(req *golem.Request, res *golem.Response) {
		// returns all fruits
	})

	FruitsRouter.POST("/new", func(req *golem.Request, res *golem.Response) {
		// creates a new fruit
	})

	FruitsRouter.DELETE("/:id", func(req *golem.Request, res *golem.Response) {
		// deletes a new fruit
	}, func(req *golem.Request, res *golem.Response, next golem.Next) {
		// a route-specific middleware
	})

	// register a middleware for all routes in this sub-router
	FruitsRouter.Use(middlewares.Logger)
}
```

Register it into the main router:

```go
func main() {
	app := golem.New()

	app.Routes("/fruits", routes.FruitsRouter)

	app.Listen(PORT, func() {
		fmt.Printf("Listening on %s\n", PORT)
	})
}
```
