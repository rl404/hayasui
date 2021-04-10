package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Requester contains http request functions.
type Requester interface {
	Get(url string, data interface{}) (int, error)
}

type request struct {
	host   string
	client http.Client
}

func newRequester(host string) Requester {
	return &request{
		host: host,
		client: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// Get is http get request.
func (r *request) Get(url string, data interface{}) (int, error) {
	req, err := http.NewRequest(http.MethodGet, r.host+url, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
