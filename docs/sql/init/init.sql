
-- 1. Users table
CREATE TABLE
    `users` (
        `id` varchar(64) NOT NULL,
        `username` varchar(255) NOT NULL,
        `email` varchar(255) NOT NULL,
        `password_hash` varchar(255) NOT NULL,
        `phone_number` varchar(255) DEFAULT NULL COMMENT '手機號碼',
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`),
        UNIQUE KEY `username` (`username`),
        UNIQUE KEY `email` (`email`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci

select * from users;

-- 2. User_profiles table
CREATE TABLE `user_profiles` (
    `user_id` VARCHAR(36) PRIMARY KEY,
    `avatar_url` VARCHAR(255) DEFAULT NULL,
    `display_name` VARCHAR(255) DEFAULT NULL,
    `bio` TEXT DEFAULT NULL COMMENT 'User self-introduction or brief description',
    `location` VARCHAR(255) DEFAULT NULL,
    `website` VARCHAR(255) DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 3. Friends table
CREATE TABLE `friends` (
    `primary_user_id` VARCHAR(36) NOT NULL,   -- 主要用戶ID
    `friend_user_id` VARCHAR(36) NOT NULL,    -- 好友ID
    `friendship_status` ENUM('active', 'blocked') NOT NULL COMMENT '友情狀態: active=正常, blocked=封鎖',
    `mute_status` BOOLEAN DEFAULT FALSE COMMENT '靜音狀態: TRUE=已靜音, FALSE=未靜音',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (primary_user_id, friend_user_id)
);

-- 4. Friend_requests table
CREATE TABLE `friend_requests` (
    `id` VARCHAR(36)  PRIMARY KEY,
    `sender_id` VARCHAR(36) NOT NULL,
    `receiver_id` VARCHAR(36) NOT NULL,
    `request_status` ENUM('pending', 'accepted', 'rejected', 'canceled') NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (sender_id, receiver_id)
);

-- 5. Groups table
CREATE TABLE `groups` (
    `id` VARCHAR(36)  PRIMARY KEY,
    `group_name` VARCHAR(255) NOT NULL UNIQUE,
    `description` TEXT DEFAULT NULL,
    `group_owner_id` VARCHAR(36) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 6. Group_members table
CREATE TABLE group_members (
    `group_id` VARCHAR(36) NOT NULL,
    `user_id` VARCHAR(36) NOT NULL,
    `role` ENUM('owner', 'admin', 'member') NOT NULL,
    `joined_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`group_id`, `user_id`)
);

-- 7. Group_invitations table
CREATE TABLE `group_invitations` (
    `id` VARCHAR(36)  PRIMARY KEY,  -- 群組邀請的唯一標識符
    `group_id` VARCHAR(36) NOT NULL,  -- 群組的唯一標識符，不允許為NULL
    `inviter_id` VARCHAR(36) NOT NULL,  -- 邀請者的唯一標識符，不允許為NULL
    `invitee_id` VARCHAR(36) NOT NULL,  -- 被邀請者的唯一標識符，不允許為NULL
    `invitation_status` ENUM('pending', 'accepted', 'rejected', 'canceled') NOT NULL,  -- 邀請的狀態，只能是指定的值，不允許為NULL
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- 創建時間，不允許為NULL，默認為當前時間
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- 更新時間，不允許為NULL，當行被更新時自動更新為當前時間
    UNIQUE KEY (`group_id`, `invitee_id`)  -- 確保每個群組只能向同一人發送一次邀請
);

-- 8. Group_requests table
CREATE TABLE group_requests (
    `id` VARCHAR(36)  PRIMARY KEY,
    `group_id` VARCHAR(36) NOT NULL,
    `requester_id` VARCHAR(36) NOT NULL,
    `request_status` ENUM('pending', 'accepted', 'rejected', 'canceled') NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (`group_id`, `requester_id`)
);

-- 9. messages table
CREATE TABLE `messages` (
    `id` VARCHAR(36)  PRIMARY KEY,
    `sender_id` VARCHAR(36) NOT NULL,
    `content` TEXT NOT NULL,
    `friend_id` VARCHAR(36) DEFAULT NULL,           -- 用於保存好友之間的對話
    `group_id` VARCHAR(36) DEFAULT NULL,            -- 用於保存群組的對話
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CHECK (`friend_id` IS NOT NULL OR `group_id` IS NOT NULL) -- 確保每條消息要麼是好友對話，要麼是群組對話
);

CREATE TABLE
    `login_record` (
        `id` bigint NOT NULL AUTO_INCREMENT,
        `name` varchar(100) DEFAULT '',
        `user_id` varchar(64) NOT NULL,
        `user_agent` varchar(2000) DEFAULT '',
        `ip` varchar(100) DEFAULT '',
        `remote_ip` varchar(200) DEFAULT '',
        `login_state` bigint DEFAULT NULL,
        `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`),
        KEY `idx_user_id` (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci

-- other demo ddl
CREATE TABLE example (
    id VARCHAR(36)  PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
