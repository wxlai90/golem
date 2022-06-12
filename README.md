# golem

<img src="https://static.wikia.nocookie.net/clashofclans/images/c/c2/Golem_info.png/revision/latest?cb=20170927231256" />

A WIP go http router.

Resembles Express for Node, features are implemented on a need basis. Default state is **broken**.

### To install

```sh
go get -u github.com/wxlai90/golem
```

Code Example:

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
