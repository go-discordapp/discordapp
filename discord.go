package discordapp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	baseURL *url.URL
)

func init() {
	baseURL, _ = url.Parse("https://discordapp.com/api/")
}

type Client struct {
	client *http.Client
	token  string

	Guild GuildService
}

type service struct {
	c *Client
}

func NewBotClient(token string, httpClient *http.Client) *Client {
	return newClient("Bot "+token, httpClient)
}

func NewClient(token string, httpClient *http.Client) *Client {
	return newClient("Bearer "+token, httpClient)
}

func newClient(token string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c := &Client{
		client: httpClient,
		token:  token,
	}
	c.Guild = GuildService(service{c})
	return c
}

func (c *Client) newRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	ua := fmt.Sprintf("DiscordBot ($%v, $%v)", "https://github.com/go-discordapp/discordapp", "0.0.1")
	req.Header.Add("User-Agent", ua)
	req.Header.Add("Authorization", c.token)
	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = withContext(ctx, req)
	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()
	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0, 4096))
	io.Copy(buf, resp.Body)
	//fmt.Println(string(buf.Bytes()))
	err = json.NewDecoder(buf).Decode(v)
	if err != nil {
		return nil, err
	}
	return &Response{resp}, nil
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	var errr ErrorResponse
	err := json.NewDecoder(r.Body).Decode(&errr)
	if err != nil {
		return err
	}
	return errr
}

type ErrorResponse struct {
	Code    int
	Message string
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func withContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}

type Response struct {
	Response *http.Response
}
