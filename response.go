package splash

import (
	"encoding/base64"
	"io/ioutil"
)

type Response struct {
	Title        string
	Url          string
	RequestedUrl string
	Base64Png    string
	Html         string
}

func (this *Response) SavedPng(filename string) error {
	png, err := base64.StdEncoding.DecodeString(this.Base64Png)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, png, 0755)
	return err
}
