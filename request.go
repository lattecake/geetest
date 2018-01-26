package geetest

import (
	"net/http"
	"net/url"
	"net"
	"time"
	"crypto/tls"
	"strings"
	"io"
	"github.com/satori/go.uuid"
	"errors"
	"io/ioutil"
	"fmt"
)

func Post(url string, params url.Values) (body []byte, err error) {

	return send("POST", url, strings.NewReader(params.Encode()))
}

func Get(url string, params url.Values) (body []byte, err error) {
	return send("GET", url+"?"+params.Encode(), nil)
}

func send(method string, httpUrl string, params io.Reader) (body []byte, err error) {
	var res *http.Response
	var proxy func(r *http.Request) (*url.URL, error)

	if GetConf().Proxy != "" {
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(GetConf().Proxy)
		}
	}

	dialer := &net.Dialer{
		Timeout:   time.Duration(1 * int64(time.Second)),
		KeepAlive: time.Duration(1 * int64(time.Second)),
	}

	transport := &http.Transport{
		Proxy: proxy, DialContext: dialer.DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(time.Second * 1),
	}
	request, err := http.NewRequest(strings.ToUpper(method), httpUrl, params)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	request.Header.Set("Worker-Id", uuid.NewV4().String())
	request.Header.Set("Token-Id", "X")
	request.Header.Set("Device-Id", "x.lattecake.com")
	request.Header.Set("Client-System", "X")
	request.Header.Set("Client-Time", time.Unix(0, time.Now().UnixNano()).Format("2006-01-02 15:04:05.999999"))

	res, err = client.Do(request)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("http status code %d", res.StatusCode))
	}

	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	return
}
