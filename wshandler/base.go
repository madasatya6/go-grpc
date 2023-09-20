package wshandler

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go_grpc"
	"go_grpc/lib/logger"
	"go_grpc/model"
)

type WSHandler struct {
	Hub   	*Hub
	Backend *backend.Backend
}

func NewWSHandler(backend *backend.Backend, hub *Hub) *WSHandler {
	return &WSHandler{
		Backend: backend,
		Hub:   hub,
	}
}

// serveWs handles websocket requests from the peer.
func (ws *WSHandler) ServeWS(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//check header first
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")

	if len(splitToken) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessToken := splitToken[1]
	session, err := ws.Backend.Usecase.AuthGetSessionByToken(r.Context(), accessToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rooms, err := ws.Backend.Usecase.GetRoomByDeviceScannerID(r.Context(), params["id"])
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if session.ResourceType != "rooms" || session.ResourceID != rooms.RoomID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(r.Context(), err.Error(), map[string]interface{}{
			"error": err,
			"tags":  []string{"websocket"},
		})
		return
	}

	if _, ok := params["id"]; ok {
		client := &Client{hub: ws.Hub, conn: conn, send: make(chan Message, 256), deviceID: params["id"]}
		client.hub.register <- client

		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.writePump()
		go client.readPump()

		meetings, err := ws.Backend.Usecase.WSCheckMeetingIsPublished(r.Context(), rooms.RoomID)

		if err != nil {
			logger.Error(r.Context(), err.Error(), map[string]interface{}{
				"error": err,
				"tags":  []string{"websocket"},
			})
		}

		for _, meeting := range meetings {
			ws.BroadcastMeeting(meeting.MeetingID)
		}
	}
}

type PayloadMeeting struct {
	ID             uint      `json:"id"`    //Meeting ID
	Title          string    `json:"title"` //Meeting title
	StartDateTime  time.Time `json:"start_date_time"`
	EndDateTime    time.Time `json:"end_date_time"`
	RoomID         uint      `json:"room_id"`
	IsCanceledFlag bool      `json:"is_canceled_flag"`
	Participants   []uint    `json:"participants"`
}

func (ws *WSHandler) EventBroadcaster(event <-chan string) {
	for {
		meetingID := <-event

		meetingIDUint, err := strconv.ParseUint(meetingID, 10, 64)
		if err == nil {
			ws.BroadcastMeeting(uint(meetingIDUint))
		}
	}
}

func (ws *WSHandler) BroadcastMeeting(meetingID uint) {
	meeting, err := ws.Backend.Usecase.GetMeetingAndParticipant(context.Background(), uint(meetingID))
	if err == nil {
		var room model.Room
		room = meeting.Room

		if room.DeviceScannerID != nil {
			for {
				message := Message{
					deviceID: *room.DeviceScannerID,
				}

				//create payload
				payload := PayloadMeeting{
					ID:             meeting.MeetingID,
					Title:          meeting.Name,
					StartDateTime:  meeting.StartAt,
					EndDateTime:    meeting.EndAt,
					RoomID:         room.RoomID,
					IsCanceledFlag: meeting.IsCanceledFlag,
				}

				participants := []uint{}

				for _, participant := range meeting.MeetingParticipant {
					participants = append(participants, participant.UserID)
				}

				payload.Participants = participants

				if err == nil {
					message.payload = payload
					ws.Hub.broadcast <- message
				}

				if room.ParentRoomID == 0 {
					break
				}

				room, err = ws.Backend.Usecase.GetRoomByID(context.Background(), room.ParentRoomID)
				if err != nil {
					break
				}
			}
		}
	}
}

func (handler *WSHandler) PanicMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error(r.Context(), "panic occured", map[string]interface{}{
					"error": rec,
				})
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
