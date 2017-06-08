package velobike

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
	Id                     *string  `json:"Id,omitempty"`
	Type                   *string  `json:"Type,omitempty"` // Possible types: "Ride", "Pay"
	Price                  *float64 `json:"Price,omitempty"`
	Rejected               *bool    `json:"Rejected,omitempty"`
	StartDate              *string  `json:"StartDate,omitempty"`
	StartBikeParkingNumber *string  `json:"StartBikeParkingNumber,omitempty"` // Only for type "Ride"
	EndBikeParkingNumber   *string  `json:"EndBikeParkingNumber,omitempty"`   // Only for type "Ride"
	Time                   *string  `json:"Time,omitempty"`                   // Only for type "Ride"
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
