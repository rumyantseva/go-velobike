package velobike

type ParkingsService struct {
	client *Client
}

type Parkings struct {
	Items []Parking `json:"Items,omitempty"`
}

type Parking struct {
	Address     *string   `json:"Address,omitempty"`
	FreePlaces  *int      `json:"FreePlaces,omitempty"`
	Id          *string   `json:"Id,omitempty"`
	IsLocked    *bool     `json:"IsLocked,omitempty"`
	Position    *Position `json:"Position,omitempty"`
	TotalPlaces *int      `json:"TotalPlaces,omitempty"`
}

type Position struct {
	Lat *float64 `json:"Lat,omitempty"`
	Lon *float64 `json:"Lon,omitempty"`
}

func (s *ParkingsService) List() (*Parkings, *Response, error) {
	u := "ride/parkings"

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var parkings = new(Parkings)
	resp, err := s.client.Do(req, parkings)
	if err != nil {
		return nil, resp, err
	}

	return parkings, resp, err
}
