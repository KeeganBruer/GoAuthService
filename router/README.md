# KBRouter

A HTTP Router in go

## Usage
#### Define A Route Handler
Below is an example route handler skeleton. The LoginRequest and LoginResponse struct definitions are not included.
Function naming convention is not enforced
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
Use the route handler when defining the router. For this example, we assume the previous example's route handler is located in the controller_login package.

```golang
// Setup router
router := kbrouter.NewRouter()
router.AddRoute("POST", "/login", controller_login.Login_PostRequest)

//Have the router listen on a specific port
err := router.Listen(PORT, func(port int) {
    fmt.Println(fmt.Sprintf("Listening on port: %d", port))
})

if err != nil {
    panic(err)
}
```

#### Sub-Routers
Define sub routers to section api endpoints by path sections
```golang
mainRouter := kbrouter.NewRouter()
mainRouter.AddRoute("POST", "/login", controller_login.Login_PostRequest)

tokenRouter := kbrouter.NewRouter()
tokenRouter.AddRoute("POST", "/verify", controller_token.Verify_PostRequest)
// register tokenRouter as a sub-router of mainRouter
mainRouter.AddSubRouter("/token", tokenRouter) 
```

#### Middleware
Below is an example that illustrates how middlewares are handled by the router.
```golang
// Setup router
router := kbrouter.NewRouter()
router.AddRoute("POST", "/login", controller_login.Login_PostRequest)

tokenRouter := kbrouter.NewRouter()
//Apply global middleware to the token sub-router
tokenRouter.AddMiddleware(func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
    res.SetHeader("MiddleWare", "custom header value")
})
tokenRouter.AddRoute("POST", "/test-1", controller_token.Verify_PostRequest)
tokenRouter.AddRoute("POST", "/test-2",
    //This is a "route specific middleware"
    func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
        res.SetHeader("MiddleWare", "second middleware")
    },
    //this is the route handler
    controller_token.Verify_PostRequest
)
//register the sub-router
router.AddSubRouter("/token", tokenRouter)
```
In the above example, /token/test-1 will have a MiddleWare header of "custom header value" and /token/test-2 will have a MiddleWare header of "second middleware". This is because global middleware are run before route handlers. The endpoint /login will not have the MiddleWare header, as the middleware was applied to the token sub-router. Route handlers are run in order, the last route handler is referred to "the route handler" while all but the last are called "route specific middleware". While there are no difference between the signatures of route handlers and middleware, middleware should not write content to the response. 

**To summarize terms:**
- Request Handler => A function with this signature: func(req *kbrouter.KBRequest, res *kbrouter.KBResponse)
    - Middleware => A request handler that handles pre-processing requests
        - Global Middleware => Assigned to router and applied across all routes. \[AddMiddleware()\] 
        - Route Specific Middleware => assigned to a specific route
    - Route Handler => Term for the route specific middlewares and the route handler. \[AddRoute()\]
    - The Route Handler => the last route handler called when processing a request 


#### Health Check
Add a quick health check route. Returns a 200 status with string "OKAY"
```golang
router.AddHealthRoute("/healthz")
```