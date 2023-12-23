package msggateway

import (
	"context"
	"im/internal/util/utils"
	"log/slog"
	"sync"
)

type UserMap struct {
	m sync.Map
}

func newUserMap() *UserMap {
	return &UserMap{}
}

func (u *UserMap) GetAll(key string) ([]*Client, bool) {
	allClients, ok := u.m.Load(key)
	if ok {
		return allClients.([]*Client), ok
	}
	return nil, ok
}

func (u *UserMap) Get(key string) ([]*Client, bool) {
	allClients, userExisted := u.m.Load(key)
	if userExisted {
		var clients []*Client
		for _, client := range allClients.([]*Client) {
			clients = append(clients, client)
		}
		return clients, userExisted
	}
	return nil, userExisted
}

func (u *UserMap) Set(key string, v *Client) {
	allClients, existed := u.m.Load(key)
	if existed {
		slog.DebugContext(context.Background(), "Set existed", "user_id", key, "client", v)
		oldClients := allClients.([]*Client)
		oldClients = append(oldClients, v)
		u.m.Store(key, oldClients)
	} else {
		slog.DebugContext(context.Background(), "Set not existed", "user_id", key, "client", v)
		var clients []*Client
		clients = append(clients, v)
		u.m.Store(key, clients)
	}
}

func (u *UserMap) delete(key string, connRemoteAddr string) (isDeleteUser bool) {
	allClients, existed := u.m.Load(key)
	if existed {
		oldClients := allClients.([]*Client)
		var a []*Client
		for _, client := range oldClients {
			if client.ctx.RemoteIP() != connRemoteAddr {
				a = append(a, client)
			}
		}
		if len(a) == 0 {
			u.m.Delete(key)
			return true
		} else {
			u.m.Store(key, a)
			return false
		}
	}
	return existed
}

func (u *UserMap) deleteClients(key string, clients []*Client) (isDeleteUser bool) {
	// 每個user不同的remoteip視為不同連線, 將此轉為map方便判斷
	m := utils.SliceToMapAny(clients, func(c *Client) (string, struct{}) {
		return c.ctx.RemoteIP(), struct{}{}
	})
	allClients, existed := u.m.Load(key)
	if existed {
		oldClients := allClients.([]*Client)
		var a []*Client
		for _, client := range oldClients {
			if _, ok := m[client.ctx.RemoteIP()]; !ok {
				a = append(a, client)
			}
		}
		if len(a) == 0 {
			u.m.Delete(key)
			return true
		} else {
			u.m.Store(key, a)
			return false
		}
	}
	return existed
}

func (u *UserMap) DeleteAll(key string) {
	u.m.Delete(key)
}
