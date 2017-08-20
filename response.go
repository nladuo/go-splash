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
	Script       string   //result of JS Script
	Console      []string //the console.log of JS Script
}

func (this *Response) SavedPng(filename string) error {
	png, err := base64.StdEncoding.DecodeString(this.Base64Png)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, png, 0755)
	return err
}
