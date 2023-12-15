package msggateway

import "net/http"

type LongConnServer interface {
	Run() error
	wsHandler(w http.ResponseWriter, r *http.Request)
	GetUserAllCons(userID string) ([]*Client, bool)
	GetUserPlatformCons(userID string, platform int) ([]*Client, bool, bool)
	Validate(s any) error
	// SetCacheHandler(cache cache.MsgModel)
	// SetDiscoveryRegistry(client discoveryregistry.SvcDiscoveryRegistry)
	KickUserConn(client *Client) error
	UnRegister(c *Client)
	// SetKickHandlerInfo(i *kickHandler)
	Compressor
	// Encoder
	// MessageHandler
}
