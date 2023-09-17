# im
instant message project

## API List

### User

1. `POST` `/api/users/register`: 註冊新用戶。
1. `POST` `/api/users/login`: 用戶登錄並返回授權令牌。
1. `GET` `/api/users/logout`: 用戶登出。
1. `GET` `/api/users/{id}`: 獲取指定ID的用戶詳細信息。
1. `PUT` `/api/users/{id}`: 更新指定ID的用戶信息。
1. `DELETE` `/api/users/{id}`: 刪除指定ID的用戶。
1. `GET` `/api/users/search?query={query}`: 根據查詢條件搜索用戶。
1. `GET` `/api/users/{id}/online-status`：獲取指定用戶ID的在線狀態。
1. `PUT` `/api/users/{id}/online-status`：更新指定用戶ID的在線狀態（例如，上線、離線、隱身等）。

### Friends

1. `GET` `/api/users/friends`: 獲取用戶的好友列表。
1. `POST` `/api/users/friends`: 向指定用戶發送好友請求。
1. `PUT` `/api/users/friends`: 更新與指定用戶的好友關係（接受/拒絕/阻止）。
1. `DELETE` `/api/users/friends`: 刪除與指定用戶的好友關係。
1. `GET` `/api/users/friend-requests`：獲取指定用戶ID收到的好友請求列表。
1. `POST` `/api/users/friend-requests`：讓指定的requester-id向指定用戶ID發送好友請求。
1. `PUT` `/api/users/friend-requests`：指定用戶ID接受或拒絕來自requester-id的好友請求。
1. `DELETE` `/api/users/friend-requests`：指定用戶ID刪除來自requester-id的好友請求。
1. `GET` `/api/users/blocked-friends`：獲取指定用戶ID的已封鎖好友列表。
1. `PUT` `/api/users/blocked-friends`：指定用戶ID封鎖或取消封鎖指定好友ID。
1. `GET` `/api/users/mutual-friends`：獲取指定用戶ID與另一指定用戶ID的共同好友列表。

### Group
1. `GET` `/api/users/{id}/groups`: 獲取用戶所屬的群組列表。
1. `POST` `/api/groups`：創建新群組。
1. `GET` `/api/groups/{id}`：根據群組ID獲取群組資訊。
1. `PUT` `/api/groups/{id}`：更新指定群組ID的群組資訊。
1. `DELETE` `/api/groups/{id}`：刪除指定群組ID的群組。
1. `GET` `/api/groups/{id}/members`：獲取指定群組ID的成員列表。
1. `POST` `/api/groups/{id}/members/{user-id}`：將指定用戶ID添加到指定群組ID。
1. `DELETE` `/api/groups/{id}/members/{user-id}`：將指定用戶ID從指定群組ID中移除。
1. `GET` `/api/groups/{id}/invitations`：獲取指定群組ID的未處理邀請列表。
1. `POST` `/api/groups/{id}/invitations/{user-id}`：向指定用戶ID發送指定群組ID的邀請。
1. `PUT` `/api/groups/{id}/invitations/{user-id}`：指定用戶ID接受或拒絕來自指定群組ID的邀請。
1. `DELETE` `/api/groups/{id}/invitations/{user-id}`：刪除來自指定群組ID的邀請。
1. `GET` `/api/groups/{id}/requests`：獲取指定群組ID的加入請求列表。
1. `POST` `/api/groups/{id}/requests/{user-id}`：指定用戶ID向指定群組ID發送加入請求。
1. `PUT` `/api/groups/{id}/requests/{user-id}`：指定群組ID接受或拒絕來自指定用戶ID的加入請求。
1. `DELETE` `/api/groups/{id}/requests/{user-id}`：刪除來自指定用戶ID的加入請求。
1. `GET` `/api/groups/search?query={search-query}`：根據搜索條件（例如群組名稱）查找群組。
1. `POST` `/api/private-messages/{message-id}/reactions`：為指定私人訊息ID添加表情符號或反應。
1. `DELETE` `/api/private-messages/{message-id}/reactions/{reaction-id}`：刪除指定私人訊息ID的表情符號或反應。
### Private Message

1. `GET` `/api/users/{id}/private-messages`: 獲取用戶的私人訊息列表。
1. `GET` `/api/users/{id}/private-messages`：獲取指定用戶ID的所有私人訊息列表。
1. `GET` `/api/users/{id}/private-messages/{friend-id}`：獲取指定用戶ID與指定好友ID之間的私人訊息列表。
1. `POST` `/api/users/{id}/private-messages/{friend-id}`：指定用戶ID向指定好友ID發送私人訊息。
1. `PUT` `/api/private-messages/{message-id}`：更新指定私人訊息ID的訊息內容（例如，標記為已讀）。
1. `DELETE` `/api/private-messages/{message-id}`：刪除指定私人訊息ID的訊息。
1. `GET` `/api/private-messages/{message-id}`：根據私人訊息ID獲取訊息詳情。
1. `POST` `/api/private-messages/{message-id}/reactions`：為指定私人訊息ID添加表情符號或反應。
1. `DELETE` `/api/private-messages/{message-id}/reactions/{reaction-id}`：刪除指定私人訊息ID的表情符號或反應。

### Group Message

1. `GET` `/api/users/{id}/group-messages`: 獲取用戶的私人訊息列表。
1. `GET` `/api/users/{id}/group-messages`：獲取指定用戶ID的所有私人訊息列表。
1. `GET` `/api/groups/{id}/messages`：獲取指定群組ID的所有訊息列表。
1. `POST` `/api/groups/{id}/messages`：向指定群組ID發送訊息。
1. `PUT` `/api/group-messages/{message-id}`：更新指定群組訊息ID的訊息內容（例如，標記為已讀）。
1. `DELETE` `/api/group-messages/{message-id}`：刪除指定群組訊息ID的訊息。
1. `GET` `/api/group-messages/{message-id}`：根據群組訊息ID獲取訊息詳情。
1. `POST` `/api/group-messages/{message-id}/reactions`：為指定群組訊息ID添加表情符號或反應。
1. `DELETE` `/api/group-messages/{message-id}/reactions/{reaction-id}`：刪除指定群組訊息ID的表情符號或反應。

## note

- 後台與前台API共用, 中間件檢測API權限(如A的好友列表只能由A確認)


## project structure

```
├─api
├─cmd
├─deployments
├─docs
│  └─sql
│      └─init
├─internal
│  ├─consts
│  ├─handler
│  ├─models
│  ├─repo
│  └─service
├─pkg
│  ├─config
│  ├─helper
│  ├─https
│  └─log
└─scripts
```

# template快速創建

- 安裝套件 ``

```
go install github.com/rickylin614/nunu@v1.0.3-rc.4
```