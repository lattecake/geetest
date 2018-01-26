package geetest

type Config struct {
	GeeTestUrl string
	Dial       int
	KeepAlive  int
	CaptchaId  string
	PrivateKey string
	Proxy      string
}

var config Config

func Init(captchaId string, privateKey string, proxy string) {
	config.GeeTestUrl = "http://api.geetest.com"
	config.Dial = 1
	config.KeepAlive = 1
	config.CaptchaId = captchaId
	config.PrivateKey = privateKey
	if proxy != "" {
		//config.Proxy = "http://10.131.30.18:8000"
		config.Proxy = proxy
	}
}

func GetConf() *Config {
	return &config
}
