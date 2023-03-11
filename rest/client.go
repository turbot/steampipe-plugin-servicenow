package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	// baseURL         *url.URL
	// servicenowToken string
	// client          *http.Client
}

type Config struct {
	URL          string
	GrantType    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

type authPayload struct {
	GrantType    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

func New(config Config) *Client {
	u, _ := url.Parse(config.URL)

	// Ensure endpoint ends with a slash
	endpoint := u.Path
	if !strings.HasSuffix(endpoint, "/") {
		endpoint += "/"
	}

	baseURL := u.ResolveReference(&url.URL{Path: endpoint})

	payload := authPayload{
		GrantType:    config.GrantType,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Username:     config.Username,
		Password:     config.Password,
	}

	var r io.Reader
	buf := &bytes.Buffer{}
	r = buf
	err := json.NewEncoder(buf).Encode(payload)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}

	req, err := http.NewRequest("POST", baseURL.ResolveReference(u).String(), r)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}

	var resp interface{}
	_, err = DoRequest(&http.Client{Timeout: 10 * time.Second}, req, resp)
	if err != nil {
		var httpError *HTTPError
		if errors.As(err, &httpError) && httpError.Code == 404 {
			return nil
		}

		return nil
	}

	fmt.Sprintln(resp)

	// client := &Client{
	// 	baseURL: u.ResolveReference(&url.URL{Path: endpoint}),
	// 	client: &http.Client{
	// 		Timeout: 10 * time.Second,
	// 	},
	// }
	// return client
	return nil
}

func DoRequest(c *http.Client, req *http.Request, resp interface{}) (statusCode int, err error) {
	httpResp, err := c.Do(req)
	if err != nil {
		return 0, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode >= 300 {
		httpError := struct {
			Message string `json:"message"`
		}{}
		err = json.NewDecoder(httpResp.Body).Decode(&httpError)
		if err != nil {
			return httpResp.StatusCode, err
		}
		return httpResp.StatusCode, &HTTPError{Code: httpResp.StatusCode, Message: httpError.Message}
	}

	if resp != nil {
		err = json.NewDecoder(httpResp.Body).Decode(resp)
		if err != nil {
			return httpResp.StatusCode, err
		}
	}
	return httpResp.StatusCode, nil
}

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	if e.Code == 0 {
		e.Code = http.StatusInternalServerError
	}
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("response %d (%s)", e.Code, http.StatusText(e.Code))
}
