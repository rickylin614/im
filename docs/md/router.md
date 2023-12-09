

# 路由

| 類別       | 方法  | 路徑                               | 說明                              |
|------------|-------|------------------------------------|----------------------------------|
| Public     | GET   | /ping                              | Ping                             |
| Public     | GET   | /metrics                           | Metrics                          |
| Public     | GET   | /example/:id                       | Get Example by ID                |
| Public     | GET   | /example                           | Get Example List                 |
| Public     | POST  | /example                           | Create Example                   |
| Public     | PUT   | /example                           | Update Example                   |
| Public     | DELETE| /example                           | Delete Example                   |
| Public     | POST  | /users/register                    | 註冊新用戶                        |
| Public     | POST  | /users/login                       | 用戶登錄                         |
| Auth       | POST  | /users/logout                      | 用戶登出                         |
| Auth       | GET   | /users/:id                         | 獲取用戶詳情                      |
| Auth       | PUT   | /users                             | 更新用戶信息                      |
| Auth       | GET   | /users/search                      | 搜索用戶                          |
| Auth       | GET   | /users/{id}/online-status          | 獲取用戶在線狀態                  |
| Auth       | PUT   | /users/{id}/online-status          | 更新用戶在線狀態                  |
| Auth       | GET   | /friend                            | 獲取好友列表                      |
| Auth       | PUT   | /friend                            | 更新好友狀態                      |
| Auth       | DELETE| /friend                            | 刪除好友                          |
| Auth       | GET   | /friend/blocked                    | 獲取封鎖的好友                    |
| Auth       | GET   | /friend/mutual                     | 獲取共同好友                      |
| Auth       | GET   | /friend-requests                   | 獲取好友請求列表                  |
| Auth       | POST  | /friend-requests                   | 發送好友請求                      |
| Auth       | PUT   | /friend-requests                   | 接受或拒絕好友請求                |
| Auth       | GET   | /group/:id                         | 獲取群組詳情                      |
| Auth       | GET   | /group                             | 獲取群組列表                      |
| Auth       | POST  | /group                             | 創建群組                          |
| Auth       | PUT   | /group                             | 更新群組                          |
| Auth       | GET   | /group-members/:id                 | 獲取群組成員                      |
| Auth       | POST  | /group-members                     | 增加群組成員                      |
| Auth       | PUT   | /group-members                     | 更新群組成員                      |
| Auth       | DELETE| /group-members                     | 刪除群組成員                      |
| Auth       | POST  | /group-invitation                  | 邀請用戶進群組                    |

