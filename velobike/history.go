package velobike

import "time"

// HistoryService contains history of user's rides.
type HistoryService struct {
	client *Client
}

// History describes the body of ride/history method response.
type History struct {
	Items          []HistoryItem `json:"Items,omitempty"`
	TotalRidesTime *string       `json:"TotalRidesTime,omitempty"`
}

// HistoryItem describes part of body responsible for a single ride.
type HistoryItem struct {
	Type *string `json:"Type,omitempty"` // Possible types: "Ride", "Pay"

	// Fields common for both "Ride" and "Pay" types
	ID        *string    `json:"Id,omitempty"`
	Price     *float64   `json:"Price,omitempty"`
	Rejected  *bool      `json:"Rejected,omitempty"`
	StartDate *time.Time `json:"StartDate,omitempty"` // layout: 2006-01-02T15:04:05

	// Fields available only for "Ride" type
	BikeID                  *string `json:"BikeId,omitempty"`
	BikeType                *string `json:"BikeType,omitempty"`
	EndDate                 *string `json:"EndDate,omitempty"`
	StartBikeParkingNumber  *string `json:"StartBikeParkingNumber,omitempty"`
	StartBikeParkingName    *string `json:"StartBikeParkingName,omitempty"`
	StartBikeParkingAddress *string `json:"StartBikeParkingAddress,omitempty"`
	StartBikeSlotNumber     *string `json:"StartBikeSlotNumber,omitempty"`
	EndBikeParkingNumber    *string `json:"EndBikeParkingNumber,omitempty"`
	EndBikeParkingName      *string `json:"EndBikeParkingName,omitempty"`
	EndBikeParkingAddress   *string `json:"EndBikeParkingAddress,omitempty"`
	EndBikeSlotNumber       *string `json:"EndBikeSlotNumber,omitempty"`
	Time                    *string `json:"Time,omitempty"`
	// Duration is always 5-10 seconds less than Time. I suspect it represents Time minus time needed to lock/unlock the bike.
	Duration *string `json:"Duration"`
	Distance *int    `json:"CoveredDistance,omitempty"`
	Text     *string `json:"Text,omitempty"`

	// Fields available only for "Pay" type
	Contract *string `json:"Contract,omitempty"`
	Status   *string `json:"Status,omitempty"`
	PanMask  *string `json:"PanMask,omitempty"`
}

// Get returns user's history.
// Use this method only for authorized users.
// Please, see an example.
func (s *HistoryService) Get() (*History, *Response, error) {
	u := "ride/history"

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var history = new(History)
	resp, err := s.client.Do(req, history)
	if err != nil {
		return nil, resp, err
	}

	return history, resp, err
}
