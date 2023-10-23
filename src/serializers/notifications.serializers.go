package serializers

type CreateNotificationRequest struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Body     string `json:"body"`
	Channel  string `json:"channel"`
	Address  string `json:"address"`
}
