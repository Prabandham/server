### Overview
This is a custom server that internally uses Gin Melody and Cors libriaries


Instead of repeating the configuration of the server over for every project
this has been extracted and base configured the way I need it.

This also comes with a basic cors policy which accepts all domains
and allows the Authorization header that is required for Jwt.



### Usage

`go get -u github.com/Prabandham/server`

`import "github.com/Prabandham/server"`

```go
// main.go or server.go


router := server.Router

app := router.Group("/")
app.GET("/", CustomHandler) // CustomHandler is some thing that you define.
api := app.Group("/api")
authorized.Use(CheckAuth()) // CheckAuth is a custom middleware that you define.


server.Start(":8080") // NOTE : is important this libriary does not hadle it for you.
```

### TODO

Add how you can overwrite cors and websockets.


