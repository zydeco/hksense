package sense

import (
	"net/http"

	"github.com/dghubble/sling"
)

const senseAPI = "https://api.hello.is/"

type Client struct {
	sling *sling.Sling
	Room  *RoomService
}

func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(senseAPI)
	return &Client{
		sling: base,
		Room:  newRoomService(base.New()),
	}
}
