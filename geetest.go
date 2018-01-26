package geetest

func New(captchaId string, privateKey string, proxy string) GeeTestService {

	// 初始化 config
	Init(captchaId, privateKey, proxy)

	return NewService(getServiceMiddleware(), config)
}

func getServiceMiddleware() (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here

	return
}
