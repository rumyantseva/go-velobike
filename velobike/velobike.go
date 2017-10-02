package velobike

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	baseURL = "http://apivelobike.velobike.ru"
)

// A Client manages communication with the API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// Services used for talking to different parts of the API.
	Parkings      *ParkingsService
	Authorization *AuthorizeService
	Profile       *ProfileService
	History       *HistoryService

	// Session ID for authorized user
	SessionID *string
}

// NewClient returns a new API client.
// If a nil httpClient is provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(baseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.Parkings = &ParkingsService{client: c}
	c.Authorization = &AuthorizeService{client: c}
	c.Profile = &ProfileService{client: c}
	c.History = &HistoryService{client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.SessionID != nil {
		req.Header.Add("SessionID", *c.SessionID)
	}

	return req, nil
}

// Response is an API response.
// This wraps the standard http.Response.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}

// BasicAuthTransport is an http.RoundTripper that authenticates all requests
// using HTTP Basic Authentication with the provided username and password.
type BasicAuthTransport struct {
	Username string // velobike.ru user id
	Password string // velobike.ru password

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.
func (t *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *BasicAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// RoundTrip implements the RoundTripper interface.
func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = cloneRequest(req) // per RoundTrip contract
	req.SetBasicAuth(t.Username, t.Password)

	return t.transport().RoundTrip(req)
}

// cloneRequest returns a clone of the provided *http.Request. The clone is a
// shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
