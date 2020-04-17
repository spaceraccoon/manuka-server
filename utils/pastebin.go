package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Based largely off https://github.com/ninedraft/simplepaste/blob/master/simplepaste.go

// Paste struct describes a paste
type Paste struct {
	Text   string
	Name   string
	APIKey string
}

// CreatePaste creates a paste and returns the Paste url
func CreatePaste(paste *Paste) (string, error) {
	values := url.Values{}
	values.Set("api_dev_key", paste.APIKey)
	values.Set("api_paste_code", paste.Text)
	values.Set("api_paste_private", "0")
	values.Set("api_paste_name", paste.Name)
	values.Set("api_paste_expire_date", "N")
	values.Set("api_paste_format", "html5")
	values.Set("api_user_key", "")
	values.Set("api_option", "paste")
	response, err := http.PostForm("https://pastebin.com/api/api_post.php", values)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(string(body), "Bad API request") {
		return "", errors.New(string(body))
	}
	return string(body), nil
}
