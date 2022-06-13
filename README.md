# golem

<img src="https://static.wikia.nocookie.net/clashofclans/images/c/c2/Golem_info.png/revision/latest?cb=20170927231256" />

A WIP go http router.

Resembles Express for Node, features are implemented on a need basis. Default state is **broken**.

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

	app.GET("/", func(req router.Request, res router.Response) {
		res.Send("Hello World")
	})

	app.GET("/todos", func(req router.Request, res router.Response) {
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

	app.Use(func(req *router.Request, res *router.Response, next router.Next) {
		log.Printf("Incoming [%s] request to %s\n", req.R.Method, req.R.URL)
		next()
	})

	app.GET("/", func(req *router.Request, res *router.Response) {
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

	logger := func(req *router.Request, res *router.Response, next router.Next) {
		log.Printf("Incoming [%s] request to %s\n", req.R.Method, req.R.URL)
		next()
	}

	app.GET("/", func(req *router.Request, res *router.Response) {
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

	logger := func(req *router.Request, res *router.Response, next router.Next) {
		todo := Todo{
			Description: "Do something",
			Done:        false,
		}
		req.Put("todo", todo)
		next()
	}

	app.GET("/", func(req *router.Request, res *router.Response) {
		if todo, ok := req.Get("todo"); ok {
			res.JSON(todo)
			return
		}

		res.Send("Couldn't find todo")

	}, logger)
}
```
