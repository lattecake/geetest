package geetest

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
	"math/rand"
)

type GeeTestService interface {
	Validate(validate ValidateRequest) (ok bool, err error)
	PreProcess(req PreProcessRequest) (res PreProcessResponse, err error)
}

type basicGeeTestService struct {
	config Config
	gen    ServiceGen
}

func NewBasicGeeTestService(config Config) (s *basicGeeTestService) {
	return &basicGeeTestService{
		config: config,
	}
}

func NewService(middleware []Middleware, config Config) GeeTestService {
	var svc GeeTestService = NewBasicGeeTestService(config)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (c *basicGeeTestService) Validate(req ValidateRequest) (ok bool, err error) {
	if len(req.Validate) != 32 {
		return false, errors.New("validate length is error")
	}
	if req.Success != 1 {
		slice := md5.Sum([]byte(req.Challenge))
		ok = hex.EncodeToString(slice[:]) == req.Validate
		return
	}
	slice := md5.Sum([]byte(c.config.PrivateKey + "geetest" + req.Challenge))
	if hex.EncodeToString(slice[:]) != req.Validate {
		return
	}

	res, err := c.gen.Validate(req, c.config.CaptchaId)

	if err != nil {
		return
	}

	slice = md5.Sum([]byte(req.SecCode))

	ok = hex.EncodeToString(slice[:]) == res.Seccode
	return ok, nil
}

func (c *basicGeeTestService) PreProcess(req PreProcessRequest) (res PreProcessResponse, err error) {
	var pre string

	res.CaptchaID = c.config.CaptchaId
	res.NewCaptcha = req.NewCaptcha

	if pre, err = c.gen.PreProcess(req, c.config.CaptchaId); err != nil {
		randOne := md5.Sum([]byte(strconv.Itoa(rand.Intn(100))))
		randTwo := md5.Sum([]byte(strconv.Itoa(rand.Intn(100))))
		challenge := hex.EncodeToString(randOne[:]) + hex.EncodeToString(randTwo[:])[0:2]
		res.Challenge = challenge
		return
	}

	res.Success = 1
	slice := md5.Sum([]byte(pre + c.config.PrivateKey))
	res.Challenge = hex.EncodeToString(slice[:])

	return
}

