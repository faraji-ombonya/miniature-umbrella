package filters

type Notification struct {
	Sender   string `form:"sender"`
	Receiver string `form:"receiver"`
	Channel  string `form:"channel"`
	Address  string `form:"address"`
}
