package geetest

import (
	"net/url"
	"strconv"
	"time"
	"encoding/json"
)

type ServiceGen struct {
}

type ValidateRes struct {
	Seccode string `json:"seccode"`
}

const (
	_register = "/register.php"
	_validate = "/validate.php"
)

func (c *ServiceGen) PreProcess(req PreProcessRequest, CaptchaID string) (challenge string, err error) {
	var (
		bs     []byte
		params url.Values
	)
	params = url.Values{}
	params = url.Values{}
	params.Set("user_id", strconv.FormatInt(req.Mid, 10))
	params.Set("new_captcha", strconv.Itoa(req.NewCaptcha))
	params.Set("client_type", req.ClientType)
	params.Set("ip_address", req.RemoteAddr)
	params.Set("gt", CaptchaID)

	if bs, err = Get(GetConf().GeeTestUrl+_register, params); err != nil {
		return
	}
	if len(bs) != 32 {
		return
	}
	challenge = string(bs)

	return
}

func (c *ServiceGen) Validate(req ValidateRequest, captchaID string) (res *ValidateRes, err error) {
	var (
		bs     []byte
		params url.Values
	)
	params = url.Values{}
	params.Set("seccode", req.SecCode)
	params.Set("challenge", req.Challenge)
	params.Set("captchaid", captchaID)
	params.Set("client_type", req.ClientType)
	params.Set("ip_address", req.RemoteAddr)
	params.Set("json_format", "1")
	params.Set("sdk", "golang_3.0.0")
	params.Set("user_id", strconv.FormatInt(req.Mid, 10))
	params.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	bs, err = Post(GetConf().GeeTestUrl+_validate, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(bs, &res)
	if err != nil {
		return
	}

	return
}
