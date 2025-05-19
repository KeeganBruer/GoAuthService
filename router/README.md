# KBRouter

A HTTP Router in go

## Usage
#### Define the route handler
```golang
func Login_PostRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
    var body LoginRequest
	req.ParseBodyJSON(&body)

	resVal := &LoginResponse{}
    // Set status code and send response
	res.SetStatusCode(200).SendJSON(resVal)
}
```

#### Route Handlers
Use the route handler when defining the router
```golang
// Setup router
router := kbrouter.NewRouter()
router.AddRoute("POST", "/login", controller_login.Login_PostRequest)

err := router.Listen(port, func() {
    msg := fmt.Sprintf("Listening on port: %d", port)
    fmt.Println(msg)
})

if err != nil {
    panic(err)
}
```

#### Sub-Routers
Define sub routers to section api endpoints by path sections
```golang
// Setup router
router := kbrouter.NewRouter()
router.AddRoute("POST", "/login", controller_login.Login_PostRequest)

tokenRouter := kbrouter.NewRouter()
tokenRouter.AddRoute("POST", "/verify", controller_token.Verify_PostRequest)
router.AddSubRouter("/token", tokenRouter)
```

#### Middleware
Below 
```golang
// Setup router
router := kbrouter.NewRouter()
router.AddRoute("POST", "/login", controller_login.Login_PostRequest)

tokenRouter := kbrouter.NewRouter()
//Apply middleware to the token sub-router
tokenRouter.AddMiddleware(func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
    res.SetHeader("MiddleWare", "custom header value")
})
tokenRouter.AddRoute("POST", "/test-1", controller_token.Verify_PostRequest)
tokenRouter.AddRoute("POST", "/test-2",
    func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
        res.SetHeader("MiddleWare", "second middleware")
    },
    controller_token.Verify_PostRequest
)
//register the sub-router
router.AddSubRouter("/token", tokenRouter)
```
In the above example, /token/test-1 will have a MiddleWare header of "custom header value" and /token/test-2 will have a MiddleWare header of "second middleware". This is because global middleware is run before route handlers. Route handlers are run in order. The endpoint /login will not have the MiddleWare header as the middleware was applied to the token sub-router.


#### Health Check
Add a quick health check route. Returns a 200 status with string "OKAY"
```golang
router.AddHealthRoute("/healthz")
```