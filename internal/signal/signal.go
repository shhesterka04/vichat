package signal

import (
	"context"
	"log"

	"github.com/pion/webrtc/v3"
	"nhooyr.io/websocket"
)

func HandleWebSocket(ctx context.Context, conn *websocket.Conn) {
	defer conn.Close(websocket.StatusInternalError, "internal error")

	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		log.Println("failed to create peer connection:", err)
		return
	}
	defer peerConnection.Close()

	// TODO: handle offer/answer exchange and ICE candidates

	conn.Close(websocket.StatusNormalClosure, "bye")
}
