package velobike

import "time"

// ProfileService  is a service to deal with user's profile.
type ProfileService struct {
	client *Client
}

// Profile describes profile method response body.
type Profile struct {
	UserID                *string    `json:"UserId,omitempty"`
	Email                 *string    `json:"Email,omitempty"`
	PhoneNumber           *string    `json:"PhoneNumber,omitempty"`
	RegisterDate          *string    `json:"RegisterDate,omitempty"`
	FirstName             *string    `json:"FirstName,omitempty"`
	LastName              *string    `json:"LastName,omitempty"`
	AvatarURL             *string    `json:"AvatarUrl,omitempty"`
	TroikaCardNumber      *string    `json:"TroikaCardNumber,omitempty"`
	TroikaPrintCardNumber *string    `json:"TroikaPrintCardNumber,omitempty"`
	Balance               *float64   `json:"Balance,omitempty"`
	Holded                *bool      `json:"Holded,omitempty"`
	HoldedAmount          *float64   `json:"HoldedAmount,omitempty"`
	TariffID              *string    `json:"TariffId,omitempty"`
	TariffStart           *time.Time `json:"TariffStart,omitempty"`
	TariffEnd             *time.Time `json:"TariffEnd,omitempty"`
}

// Get returns a profile of the current user.
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
