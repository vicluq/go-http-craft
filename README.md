# HTTP Boost

Golang 1.22 introduced an improved HTTP native module with many new features, which allows us to create servers without the need of extensive libraries easily.

Although it came with greate improvements, some features, such as **routers** and **middlewares** are still important but we must implement it by ourselves.

This package aims to help those who are in the need for such features but could not find it anywhere or do not want to maintain it.

Beware that this is still being tested and is still under development.

- [Routers](#routers)
- [Middlewares](#middlewares)

## Routers

Routers are very usefull when it comes to keeping our code organized and well structured. Therefore, I have created a new type, `Router` with the following features.

### Creating a Router

We must use the `NewRouter` factory which will return a new router. We must pass in the base path for the router.

```golang
AuthRouter := router.NewRouter("/auth")
```

### Methods

#### `Router.AddRoute(path string, handler http.Handler, middlewares ...Middleware)`

This method adds a HTTP Handler to the given `path` and it can receive route specific middlewares to be executed only for this route.

```golang
AuthRouter.AddRoute("POST /register", h.AuthRegisterHandler, m.Profile())
```

#### `Router.AddMiddleware(middleware Middleware)`

Adds a Router level middleware to be executed for every route from the router.

```golang
AuthRouter.AddMiddleware(m.EnableCORS(*GetCORSConfig()))
```

#### `Router.Register(root *http.ServeMux)`

This will be responsible for registering the router on the ServeMux, whic is the server path multiplexer, basically, it means adding the router's routes to the server.

In the example bellow, I have created a `GetAuthRouter` function which returns a `Router` whit registered routes and middlewares. After retrieving it, we must register it in our `serverMux`.

```golang
// Server config
serverMux := http.NewServeMux()
server := &http.Server{
    Handler: serverMux,
    Addr:    fmt.Sprintf(":%v", 8080),
}

// Declare Routers
authRouter := routes.GetAuthRouter()
authRouter.Register(&serverMux)
```

#### `Router.GetBasePath()`

This function returns the base path for the router.

## Middlewares

Middlewares are useful to create request pipelines, in order to perform authentication, enabling;configuring CORS, etc.

I have added the Middleware type, which will guide us through it.

### Creating a Middleware

The structure bellow represents a function that returns middleware to be chained along with the handlers and other middlewares.

We perform our middleware logic and call the `next.ServeHTTP(res, req)`, which will propagate - serve - the request to the following mids/handlers in the chain.

```golang
func AuthMiddleware() m.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	            // Verify auth token
	
	            if is_verified {
	                next.ServeHTTP(res, req) // call the request handler
	            }
	            else {
	                http.Error(res, "Access not allowed.", http.StatusForbidden)
	            }
		})
	}
}

ProetectedRouter.AddRoute("GET /protected/books", protectedBooksHandler, AuthMiddleware())
```

When the request comes in, it will first pass through the middleware, if the auth token is verified correctly, it will then allow the request to continue.

### Pre-made Middlewares: CORS

```golang
corsConfig := m.CORS{
    AccessControlAllowOrigin: "*",
}

AuthRouter.AddMiddleware(m.EnableCORS(corsConfig))
```

Bellow, there are all of the cors configuration available in the struct.

```golang
type CORS struct {
	AccessControlAllowOrigin string
	AllowMethods             []string
	AllowHeaders             []string
	AllowCredentials         bool
	ExposeHeaders            []string
	MaxAge                   time.Duration
}
```
