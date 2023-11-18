
-- 1. Users table
CREATE TABLE `users` (
    `id` VARCHAR(64) NOT NULL COMMENT '用戶的唯一標識',
    `username` VARCHAR(255) NOT NULL COMMENT '用戶名',
    `email` VARCHAR(255) NOT NULL COMMENT '電子郵件地址',
    `password_hash` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用戶密碼的哈希值',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '賬戶創建的時間',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '賬戶信息最後更新的時間',
    `phone_number` VARCHAR(255) DEFAULT NULL COMMENT '手機號碼',
    `nickname` VARCHAR(50) DEFAULT NULL COMMENT '暱稱',
    `status` INT NOT NULL DEFAULT '0' COMMENT '用戶狀態碼',
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`) COMMENT '用戶名唯一索引',
    UNIQUE KEY `email` (`email`) COMMENT '電子郵件地址唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='儲存用戶基本信息的資料表';


-- 2. User_profiles table
CREATE TABLE `user_profiles` (
    `user_id` VARCHAR(36) NOT NULL PRIMARY KEY COMMENT '用戶的唯一標識符',
    `avatar_url` VARCHAR(255) DEFAULT NULL COMMENT '用戶頭像的URL地址',
    `display_name` VARCHAR(255) DEFAULT NULL COMMENT '用戶展示在個人檔案上的名稱',
    `bio` TEXT DEFAULT NULL COMMENT '用戶的自我介紹或簡短描述',
    `location` VARCHAR(255) DEFAULT NULL COMMENT '用戶的地理位置',
    `website` VARCHAR(255) DEFAULT NULL COMMENT '用戶個人或職業網站的URL',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '創建個人檔案的時間戳記',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '個人檔案最後更新的時間戳記'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='儲存用戶個人信息和社交連結的資料表';


-- 2-1. user_social_accounts table
CREATE TABLE `user_social_accounts` (
    `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '用戶社交賬戶的主鍵ID',
    `user_id` VARCHAR(36) NOT NULL COMMENT '本地用戶表的唯一識別ID',
    `provider` ENUM('facebook', 'google', 'twitter', 'linkedin') NOT NULL COMMENT '指示社交媒體平台提供者',
    `provider_id` VARCHAR(255) NOT NULL COMMENT '社交媒體平台上的用戶唯一識別符',
    `provider_data` TEXT COMMENT '儲存從第三方服務提供者那裡獲得的用戶原始資料',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '記錄創建的時間戳記',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '記錄最後更新的時間戳記'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存儲用戶的社交媒體賬戶信息，支持多種登入方式';


-- 3. Friends table
CREATE TABLE `friends` (
    `id` VARCHAR(36) COMMENT '唯一識別ID' PRIMARY KEY,
    `p_user_id` VARCHAR(36) NOT NULL COMMENT '主要用戶ID，參考users表',       
    `p_user_name` VARCHAR(255) DEFAULT NULL COMMENT '主要用戶的用戶名', 
    `f_user_id` VARCHAR(36) NOT NULL COMMENT '好友的用戶ID，參考users表',        
    `f_user_name` VARCHAR(255) DEFAULT NULL COMMENT '好友的用戶名',
    `message_id` VARCHAR(36) DEFAULT NULL COMMENT '與該好友關係相關的訊息記錄用id',   
    `status` ENUM('active', 'blocked') NOT NULL COMMENT '友情狀態: active=正常, blocked=封鎖',
    `mute` BOOLEAN DEFAULT FALSE COMMENT '靜音狀態: TRUE=已靜音, FALSE=未靜音',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '記錄創建時間',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '記錄最後更新時間',
    UNIQUE KEY `unique_relationship` (`p_user_id`, `f_user_id`) COMMENT '確保用戶間的關係是唯一的'
) COMMENT='儲存用戶之間友誼的資料，包含狀態和相關選項';


-- 4. Friend_requests table
CREATE TABLE `friend_requests` (
    `id` VARCHAR(36) PRIMARY KEY COMMENT '唯一識別符號，代表好友請求的ID',
    `sender_id` VARCHAR(36) NOT NULL COMMENT '發送請求者的用戶ID',
    `sender_name` VARCHAR(255) DEFAULT NULL COMMENT '發送請求者的用戶名',
    `receiver_id` VARCHAR(36) NOT NULL COMMENT '接收請求者的用戶ID',
    `receiver_name` VARCHAR(255) DEFAULT NULL COMMENT '接收請求者的用戶名',
    `request_status` ENUM('pending', 'accepted', 'rejected', 'canceled') NOT NULL COMMENT '請求的當前狀態',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '請求創建的時間戳記',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '請求最後更新的時間戳記',
    UNIQUE KEY `unique_friend_request` (`sender_id`, `receiver_id`) COMMENT '保證每個發送者對每個接收者只有一個活動的請求'
) COMMENT='存儲用戶間發送和接收好友請求的數據';


-- 5. Groups table
CREATE TABLE `groups` (
    `id` VARCHAR(36) PRIMARY KEY COMMENT '唯一識別符號，代表群組的ID',
    `group_name` VARCHAR(255) NOT NULL UNIQUE COMMENT '群組名稱，必須唯一',
    `description` TEXT DEFAULT NULL COMMENT '群組描述，提供群組的詳細信息',
    `group_owner_id` VARCHAR(36) NOT NULL COMMENT '群組擁有者的用戶ID，指向users表的ID',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '群組創建的時間戳記',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '群組最後更新的時間戳記'
) COMMENT='儲存群組資料的表，包括名稱、描述、擁有者以及創建和更新時間';


-- 6. Group_members table
CREATE TABLE `group_members` (
    `group_id` VARCHAR(36) NOT NULL COMMENT '唯一标识群组的ID',
    `user_id` VARCHAR(36) NOT NULL COMMENT '唯一标识用户的ID',
    `role` VARCHAR(10) NOT NULL COMMENT '成员的角色,可以是owner、admin或member',
    `joined_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '成员加入群组的时间',
    PRIMARY KEY (`group_id`, `user_id`)
) COMMENT='儲存群組成員資料和他們的角色，包括加入群組的時間';


-- 7. Group_invitations table
CREATE TABLE `group_invitations` (
    `id` VARCHAR(36) PRIMARY KEY COMMENT '群組邀請的唯一標識符',
    `group_id` VARCHAR(36) NOT NULL COMMENT '群組的唯一標識符',
    `inviter_id` VARCHAR(36) NOT NULL COMMENT '邀請者的唯一標識符',
    `invitee_id` VARCHAR(36) NOT NULL COMMENT '被邀請者的唯一標識符',
    `invitation_status` ENUM('pending', 'accepted', 'rejected', 'canceled') NOT NULL COMMENT '邀請的狀態',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '記錄創建邀請的時間戳',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '記錄邀請最後更新的時間戳',
    UNIQUE KEY `group_invitee_unique` (`group_id`, `invitee_id`) COMMENT '確保每個群組對同一個被邀請者的唯一邀請'
) COMMENT='儲存群組邀請信息，包括邀請狀態以及創建和更新時間';

-- 8. Group_requests table
CREATE TABLE `group_requests` (
    `id` VARCHAR(36) PRIMARY KEY COMMENT '群組請求的唯一標識符',
    `group_id` VARCHAR(36) NOT NULL COMMENT '發出請求的群組的唯一標識符',
    `requester_id` VARCHAR(36) NOT NULL COMMENT '提出請求的用戶的唯一標識符',
    `request_status` ENUM('pending', 'accepted', 'rejected', 'canceled') NOT NULL COMMENT '請求的當前狀態',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '請求創建的時間戳記',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '請求最後更新的時間戳記',
    UNIQUE KEY `group_requester_unique` (`group_id`, `requester_id`) COMMENT '確保每個用戶對同一群組的唯一請求'
) COMMENT='儲存用戶對群組加入請求的資料表，包含請求狀態和時間戳記';


-- 9. messages table
CREATE TABLE `messages` (
    `id` VARCHAR(36)  PRIMARY KEY,
    `sender_id` VARCHAR(36) NOT NULL,
    `content` TEXT NOT NULL,
    `friend_id` VARCHAR(36) DEFAULT NULL,           -- 用於保存好友之間的對話 friend.message_id
    `group_id` VARCHAR(36) DEFAULT NULL,            -- 用於保存群組的對話 group.id
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CHECK (`friend_id` IS NOT NULL OR `group_id` IS NOT NULL) -- 確保每條消息要麼是好友對話，要麼是群組對話
);

CREATE TABLE `login_record` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '唯一識別碼，自動增加',
    `name` VARCHAR(100) DEFAULT '' COMMENT '用戶名稱',
    `user_id` VARCHAR(64) NOT NULL COMMENT '關聯的用戶ID',
    `user_agent` VARCHAR(2000) DEFAULT '' COMMENT '用戶代理信息，包括操作系統、瀏覽器等',
    `ip` VARCHAR(100) DEFAULT '' COMMENT '用戶的IP地址',
    `remote_ip` VARCHAR(200) DEFAULT '' COMMENT '遠程IP地址，可能是經過代理的',
    `login_state` bigint DEFAULT NULL COMMENT '登錄狀態，可以用來表示成功或失敗等',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '紀錄創建的時間',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) COMMENT '用戶ID索引，便於查詢'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用戶登錄紀錄表，包含登錄詳細信息';


-- other demo ddl
CREATE TABLE example (
    id VARCHAR(36)  PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
