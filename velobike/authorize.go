package velobike

// AuthorizeService is necessary for the methods which need authorization.
// Velobike uses basic auth for the authorization and gives a 'SessionId' token.
type AuthorizeService struct {
	client *Client
}

// Authorization describes the body of profile/authorize method response.
type Authorization struct {
	SessionId *string `json:"SessionId,omitempty"`
}

// Authorize method authorizes user and get receives token.
func (s *AuthorizeService) Authorize() (*Authorization, *Response, error) {
	u := "profile/authorize"

	req, err := s.client.NewRequest("POST", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var auth = new(Authorization)

	resp, err := s.client.Do(req, auth)
	if err != nil {
		return nil, resp, err
	}

	return auth, resp, err
}
