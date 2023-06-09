package http

type Middleware = func(next StdHandler) StdHandler
