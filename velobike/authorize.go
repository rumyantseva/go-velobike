package velobike

type AuthorizeService struct {
	client *Client
}

type Authorization struct {
	SessionId *string `json:"SessionId,omitempty"`
}

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
