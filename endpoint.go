package geetest

// WechatRequest collects the request parameters for the Validate method.
type ValidateRequest struct {
	Challenge  string `json:"geetest_challenge"`
	Validate   string `json:"geetest_validate"`
	SecCode    string `json:"geetest_seccode"`
	Success    int    `json:"geetest_success"`
	ClientType string `json:"client_type"`
	Mid        int64  `json:"mid"`
	RemoteAddr string `json:"remote_addr"`
}

// WechatResponse collects the response parameters for the Validate method.
type ValidateResponse struct {
	Ok  bool
	Err error
}

type PreProcessRequest struct {
	RemoteAddr string `json:"remote_addr"`
	Mid        int64  `json:"mid"`
	ClientType string `json:"client_type"`
	NewCaptcha int    `json:"new_captcha"`
}

type PreProcessResponse struct {
	Success    int8   `json:"success"`
	CaptchaID  string `json:"gt"`
	Challenge  string `json:"challenge"`
	NewCaptcha int    `json:"new_captcha"`
}
