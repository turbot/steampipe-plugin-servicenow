package goservicenow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client struct {
	baseURL         *url.URL
	servicenowToken string
	// client          *http.Client
}

type Config struct {
	InstanceURL  string
	GrantType    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

type OAuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

func New(config Config) (client *Client, err error) {
	baseURL, _ := url.Parse(config.InstanceURL)

	resp := &OAuthTokenResponse{}
	_, err = authenticate(config, resp)
	if err != nil {
		return nil, err
	}

	return &Client{
		baseURL:         baseURL,
		servicenowToken: resp.AccessToken,
	}, nil
}

func authenticate(config Config, resp interface{}) (statusCode int, err error) {
	endpointUrl, _ := url.Parse(config.InstanceURL)
	endpointUrl = endpointUrl.JoinPath("oauth_token.do")
	method := "POST"

	payloadParameters := url.Values{
		"grant_type":    {config.GrantType},
		"client_id":     {config.ClientID},
		"client_secret": {config.ClientSecret},
		"username":      {config.Username},
		"password":      {config.Password},
	}
	payload := strings.NewReader(payloadParameters.Encode())

	client := &http.Client{}
	req, err := http.NewRequest(method, endpointUrl.String(), payload)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	httpResp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
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
			fmt.Println(err.Error())
			return httpResp.StatusCode, err
		}
	}
	return httpResp.StatusCode, nil
}

func (c *Client) listTable(tableName string, limit int, result interface{}) error {
	endpointUrl := c.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s", tableName))

	queryUrl := endpointUrl.Query()
	queryUrl.Add("sysparm_limit", strconv.Itoa(limit))
	endpointUrl.RawQuery = queryUrl.Encode()

	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.doAPI(*req, result)
}

func (c *Client) getTable(tableName string, sysId string, result interface{}) error {
	endpointUrl := c.baseURL.JoinPath(fmt.Sprintf("api/now/table/%s/%s", tableName, sysId))
	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.doAPI(*req, result)
}

func (c *Client) doAPI(req http.Request, result interface{}) error {
	client := &http.Client{}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.servicenowToken))
	res, err := client.Do(&req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer res.Body.Close()

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON", err.Error())
		return err
	}
	return nil
}

func (c *Client) listArticles(limit, offset int, result interface{}) error {
	endpointUrl := c.baseURL.JoinPath("api/sn_km_api/knowledge/articles")

	queryUrl := endpointUrl.Query()
	queryUrl.Add("limit", strconv.Itoa(limit))
	queryUrl.Add("offset", strconv.Itoa(offset))
	endpointUrl.RawQuery = queryUrl.Encode()

	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.doAPI(*req, result)
}

func (c *Client) getArticle(sysId string, result interface{}) error {
	endpointUrl := c.baseURL.JoinPath("api/sn_km_api/knowledge/articles").JoinPath(sysId)
	method := "GET"
	req, err := http.NewRequest(method, endpointUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.doAPI(*req, result)
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
