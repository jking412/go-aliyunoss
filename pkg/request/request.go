package request

import (
	"aliyunoss/pkg/logger"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetRequest(url string, obj interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("request", zap.String("url", url),
			zap.String("err", err.Error()))
		return err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(data, obj)
	if err != nil {
		logger.Error("request", zap.String("json", "unmarshal"),
			zap.String("err", err.Error()))
		return err
	}
	return nil
}

func PostRequest(url string, body string, obj interface{}) error {

	ossReq := strings.NewReader(body)
	rep, err := http.Post(url, "application/json", ossReq)

	if err != nil {
		logger.Error("request", zap.String("url", url),
			zap.String("body", body),
			zap.String("err", err.Error()))
		return err
	}

	defer rep.Body.Close()

	data, _ := ioutil.ReadAll(rep.Body)

	err = json.Unmarshal(data, obj)
	if err != nil {
		logger.Error("request", zap.String("json", "unmarshal"),
			zap.String("err", err.Error()))
		return err
	}
	return nil
}
