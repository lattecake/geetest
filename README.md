# gee test sdk

`go get github.com/lattecake/geetest`

## useing

```go
geeCaptcha = geetest.New("your captcha Id", "your private Key", "http_proxy or nil")

```

### pre process

```go
var req geetest.PreProcessRequest
req.RemoteAddr = "clientip"
req.Mid = 1
req.ClientType = "web"
req.NewCaptcha = newCaptcha
res, err := geeCaptcha.PreProcess(req)
```


### validate

```go
req := geetest.ValidateRequest{
	Challenge:  "geetest_challenge",
	Validate:   "geetest_validate",
	SecCode:    "geetest_seccode",
	Success:    1, // or 0
	RemoteAddr: "your client ip",
}

if ok, err = geeCaptcha.Validate(req); err != nil {
	fmt.println(err.Error())
	return
}

if !ok {
	return
}
```
