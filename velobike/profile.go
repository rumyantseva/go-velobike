package velobike

type ProfileService struct {
	client *Client
}

type Profile struct {
	UserId                *string  `json:"UserId,omitempty"`
	Email                 *string  `json:"Email,omitempty"`
	PhoneNumber           *string  `json:"PhoneNumber,omitempty"`
	RegisterDate          *string  `json:"RegisterDate,omitempty"`
	FirstName             *string  `json:"FirstName,omitempty"`
	LastName              *string  `json:"LastName,omitempty"`
	AvatarUrl             *string  `json:"AvatarUrl,omitempty"`
	TroikaCardNumber      *string  `json:"TroikaCardNumber,omitempty"`
	TroikaPrintCardNumber *string  `json:"TroikaPrintCardNumber,omitempty"`
	Balance               *float64 `json:"Balance,omitempty"`
	Holded                *bool    `json:"Holded,omitempty"`
	HoldedAmount          *float64 `json:"HoldedAmount,omitempty"`
	TariffId              *string  `json:"TariffId,omitempty"`
	TariffStart           *string  `json:"TariffStart,omitempty"`
	TariffEnd             *string  `json:"TariffEnd,omitempty"`
}

// Use this method only for authorized users.
// Please, see an example.
func (s *ProfileService) Get() (*Profile, *Response, error) {
	u := "profile"

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var profile = new(Profile)
	resp, err := s.client.Do(req, profile)
	if err != nil {
		return nil, resp, err
	}

	return profile, resp, err
}
