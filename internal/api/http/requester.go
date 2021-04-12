package http

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	_http "net/http"
	"time"

	"github.com/rl404/hayasui/internal/model"
)

type requester interface {
	Get(url string, data interface{}) (int, error)
}

type request struct {
	host   string
	client _http.Client
}

func newRequester(host string) requester {
	return &request{
		host: host,
		client: _http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// Get is http get request.
func (r *request) Get(url string, data interface{}) (int, error) {
	req, err := _http.NewRequest(_http.MethodGet, r.host+url, nil)
	if err != nil {
		return _http.StatusInternalServerError, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return _http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return _http.StatusInternalServerError, err
	}

	if resp.StatusCode != _http.StatusOK {
		var respErr model.ResponseError
		if json.Unmarshal(body, &respErr) == nil {
			return resp.StatusCode, errors.New(respErr.Data)
		}
		return resp.StatusCode, errors.New(resp.Status)
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return _http.StatusInternalServerError, err
	}

	return _http.StatusOK, nil
}
