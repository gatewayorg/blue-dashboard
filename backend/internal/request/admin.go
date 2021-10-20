package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gatewayorg/blue-dashboard/internal/model"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type BlueAdmin interface {
	GetUri(ctx context.Context, uriInfo *model.GetUri) (*model.UriRes, error)
	CreateUri(ctx context.Context) error
}

type adminImpl struct {
	url string
	key string
}

func NewAdmin(url, key string) *adminImpl {
	return &adminImpl{
		url: url,
		key: key,
	}
}

func (a *adminImpl) GetUri(_ context.Context, uriInfo *model.GetUri) (*model.UriRes, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", a.url, "/uri/user"), bytes.NewBuffer(uriInfo.Bytes()))
	if err != nil {
		log.Error("get uri list: new request", zap.String("url", a.url), zap.Error(err))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.key)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("get uri list: do", zap.String("url", a.url), zap.Error(err))
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("get uri list: read body", zap.String("url", a.url), zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Error("get uri list: not ok", zap.Int("code", resp.StatusCode))
		return nil, ErrNotStatusOk
	}
	res := model.UriRes{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Error("get uri list: json unmarshal", zap.Int("code", resp.StatusCode))
		return nil, err
	}
	return &res, nil
}

func (a *adminImpl) CreateUri(ctx context.Context, info *model.UriInfo) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", a.url, "/uri"), bytes.NewBuffer(info.Bytes()))
	if err != nil {
		log.Error("create uri: new request", zap.String("url", a.url), zap.Error(err))
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("create uri: do", zap.String("url", a.url), zap.Error(err))
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Error("get uri list: not ok", zap.Int("code", resp.StatusCode))
		return ErrNotStatusOk
	}
	return nil
}
