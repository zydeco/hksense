package sense

import (
	"net/http"

	"github.com/dghubble/sling"
)

type RoomMeasurement struct {
	Temperature  Measurement `json:"temperature"`
	Humidity     Measurement `json:"humidity"`
	Light        Measurement `json:"light"`
	Sound        Measurement `json:"sound"`
	Particulates Measurement `json:"particulates"`
}

type Measurement struct {
	Value           float64   `json:"value"`
	Message         string    `json:"message"`
	IdealConditions string    `json:"ideal_conditions"`
	Condition       string    `json:"condition"`
	LastUpdated     Timestamp `json:"last_updated_utc"`
	Unit            string    `json:"unit"`
}

type RoomService struct {
	sling *sling.Sling
}

func newRoomService(sling *sling.Sling) *RoomService {
	return &RoomService{
		sling: sling.Path("v1/room/"),
	}
}

type CurrentMeasurementParams struct {
	TempUnit string `url:"temp_unit,omitempty"`
}

func (s *RoomService) CurrentMeasurement(tempUnit string) (*RoomMeasurement, *http.Response, error) {
	room := new(RoomMeasurement)
	apiError := new(APIError)
	params := &CurrentMeasurementParams{TempUnit: tempUnit}
	resp, err := s.sling.New().Get("current").QueryStruct(params).Receive(room, apiError)
	return room, resp, relevantError(err, *apiError)
}
