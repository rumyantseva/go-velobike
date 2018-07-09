package velobike

// ParkingsService is a service to deal with parkings.
type ParkingsService struct {
	client *Client
}

// Parkings describes ride/parkings method response.
type Parkings struct {
	Items []Parking `json:"Items,omitempty"`
}

// Parking describes part of body responsible for a single parking.
type Parking struct {
	Address             *string   `json:"Address,omitempty"`
	FreeElectricPlaces  *int      `json:"FreeElectricPlaces,omitempty"`
	FreeOrdinaryPlaces  *int      `json:"FreeOrdinaryPlaces,omitempty"`
	FreePlaces          *int      `json:"FreePlaces,omitempty"`
	HasTerminal         *bool     `json:"HasTerminal,omitempty"`
	ID                  *string   `json:"Id,omitempty"`
	IsFavourite         *bool     `json:"IsFavourite,omitempty"`
	IsLocked            *bool     `json:"IsLocked,omitempty"`
	Name                *string   `json:"Name,omitempty"`
	Position            *Position `json:"Position,omitempty"`
	StationTypes        []string  `json:"StationTypes,omitempty"`
	TotalElectricPlaces *int      `json:"TotalElectricPlaces,omitempty"`
	TotalOrdinaryPlaces *int      `json:"TotalOrdinaryPlaces,omitempty"`
	TotalPlaces         *int      `json:"TotalPlaces,omitempty"`
}

// Position describes parking's geo location.
type Position struct {
	Lat *float64 `json:"Lat,omitempty"`
	Lon *float64 `json:"Lon,omitempty"`
}

// List returns list of existed parkings.
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
