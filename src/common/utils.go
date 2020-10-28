package common

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
	"strings"
)

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}



func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func RunInLinux(cmd string) (int,string){
	fmt.Println("Running Linux cmd:" , cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	var resCode int
	if err != nil {
		resCode = -1
		fmt.Println(err.Error())
	}
	return resCode, strings.TrimSpace(string(result))
}
func RunInWindows(cmd string) (int,string){
	fmt.Println("Running Windows cmd:", cmd)
	result, err := exec.Command("cmd", "/c", cmd).Output()
	var resCode int
	if err != nil {
		resCode = -1
		fmt.Println(err.Error())
	}
	return resCode, strings.TrimSpace(string(result))
}