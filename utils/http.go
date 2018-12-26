package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/logs"
)

// 发送post请求
func FetchPost(param interface{}, url string) (error, resu interface{}) {
	paramData := []byte{}
	if param != nil {
		//logs.Debug(param)
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false) //取消特殊字符转义
		err := encoder.Encode(param)
		if err != nil {
			return err, nil
		}
		paramData = buffer.Bytes()
	}

	var urlBuf bytes.Buffer
	urlBuf.WriteString(url)

	bodyReader := bytes.NewBuffer(paramData)
	resp, err := http.Post(urlBuf.String(), "application/json;charset=UTF-8", bodyReader)
	if err != nil {
		logs.Error(err)
		return err, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return err, nil
	}

	if resp.StatusCode != http.StatusOK {
		logs.Error(fmt.Sprintf("%d", resp.StatusCode))
		return errors.New(fmt.Sprintf("%d", resp.StatusCode)), nil
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(string(body)), &jsonMap)
	if err != nil {
		logs.Error(err)
		return err, nil
	}

	//logs.Debug(jsonMap)
	return nil, jsonMap
}
