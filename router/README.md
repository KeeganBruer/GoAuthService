# KBRouter

A HTTP Router in go

## Usage
Define the route handler
```golang
func Login_PostRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
    var body LoginRequest
	req.ParseBodyJSON(&body)

	resVal := &LoginResponse{}
    // Set status code and send response
	res.SetStatusCode(200).SendJSON(resVal)
}
```

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