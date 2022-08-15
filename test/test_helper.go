package test

import (
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type RequestParams struct {
	ParamKey   string
	ParamValue string
}

// RandSeq генерирует случайную строку длины n
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func MakePOSTRequest(paramKey, paramValue string) (*http.Request, error) {
	param := url.Values{}
	param.Add(paramKey, paramValue)
	req, err := http.NewRequest(http.MethodPost, paramValue, strings.NewReader(param.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.PostForm = param

	if err != nil {
		err = errors.Errorf("MakePOSTRequest err: %v", err)
	}

	return req, err

}

func MakePOSTRequestSet(p []RequestParams) (rs []*http.Request, err error) {
	for _, param := range p {
		r, err := MakePOSTRequest(param.ParamKey, param.ParamValue)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}

	return rs, nil
}
