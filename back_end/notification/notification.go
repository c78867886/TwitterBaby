package notification

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/gorilla/websocket"
)

type(
	// Handler : Handler for managing notifications.
	Handler struct {
		upgrader	websocket.Upgrader
		manager 	clientManager
	}

	clientManager struct {
		
	}
	client struct {
		username	string
		socket		*websocket.Conn
	}
)

// NewHandler : Create a Handler instance
func NewHandler() (h *Handler) {
	h = &Handler{websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {return true}}, clientManager{}}
	return h
}

// GetConnection : Upgrade a client connection to web socket.
func (h *Handler) GetConnection(c echo.Context) (err error) {
	/*conn, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "Failed to make web socket connection."}
	}

	return c.NoContent(http.StatusCreated)*/
	return
}




func sendNewTweetNotif() {

}

func sendFollowNotif() {

}