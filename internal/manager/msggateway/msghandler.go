package msggateway

type MessageHandler interface {
	// GetSeq(context context.Context, data *Req) ([]byte, error)
	// SendMessage2Client(context context.Context, data *po.Message) ([]byte, error)
	// SendSignalMessage(context context.Context, data *Req) ([]byte, error)
	// PullMessageBySeqList(context context.Context, data *Req) ([]byte, error)
	// UserLogout(context context.Context, data *Req) ([]byte, error)
	// SetUserDeviceBackground(context context.Context, data *Req) ([]byte, bool, error)
}
