-- Active: 1693666258757@@127.0.0.1@3306@demo

-- 創建用戶清單
DROP PROCEDURE IF EXISTS CreateRickyUsers;
DELIMITER //
CREATE PROCEDURE CreateRickyUsers()
BEGIN
    DECLARE i INT DEFAULT 1;
    DECLARE userExists INT DEFAULT 1;
    DECLARE userName VARCHAR(255);
    DECLARE userEmail VARCHAR(255);
    DECLARE phonenumber VARCHAR(255);
    DECLARE usernickname VARCHAR(255);

    WHILE i <= 50 DO
        SET userName = CONCAT('ricky', LPAD(i, 3, '0'));
        SET userEmail = CONCAT(userName, '@example.com');
        SET phonenumber = CONCAT('0912345', LPAD(i, 3, '0'));
        SET usernickname = CONCAT(LPAD(i, 3, '0'), '號匿名用戶');

        INSERT IGNORE INTO `users` (`id`, `username`, `email`, `password_hash`, `phone_number`, `nickname`, `status`)
        VALUES (UUID(), userName, userEmail, '1e901a92345b5cff6018316d1cf669f7', phonenumber, usernickname, 1);

        SET i = i + 1;
    END WHILE;
END //
DELIMITER ;
CALL CreateRickyUsers();
DROP PROCEDURE CreateRickyUsers;


-- 創建好友清單
DROP PROCEDURE IF EXISTS CreateFriendshipsForRicky001;
DELIMITER //
CREATE PROCEDURE CreateFriendshipsForRicky001()
BEGIN
    DECLARE i INT DEFAULT 2;  -- 从 ricky002 开始
    DECLARE currentUserID VARCHAR(36);
    DECLARE currentUserUsername VARCHAR(255);
    DECLARE friendUserID VARCHAR(36);
    DECLARE friendUsername VARCHAR(255);
    DECLARE sharedMessageID VARCHAR(36);

    -- 获取 ricky001 的信息
    SELECT `id`, `username` INTO currentUserID, currentUserUsername FROM `users` WHERE `username` = 'ricky001';

    WHILE i <= 50 DO
        -- 生成每对好友关系的共享 message_id
        SET sharedMessageID = UUID();

        -- 获取当前好友的信息
        SET friendUsername = CONCAT('ricky', LPAD(i, 3, '0'));
        SELECT `id` INTO friendUserID FROM `users` WHERE `username` = friendUsername;

        -- 插入 ricky001 -> 当前好友的关系
        INSERT IGNORE INTO `friends` (`id`, `p_user_id`, `p_user_name`, `f_user_id`, `f_user_name`, `message_id`, `status`, `mute`)
        VALUES (UUID(), currentUserID, currentUserUsername, friendUserID, friendUsername, sharedMessageID, 'active', FALSE);

        -- 插入 当前好友 -> ricky001 的关系
        INSERT IGNORE INTO `friends` (`id`, `p_user_id`, `p_user_name`, `f_user_id`, `f_user_name`, `message_id`, `status`, `mute`)
        VALUES (UUID(), friendUserID, friendUsername, currentUserID, currentUserUsername, sharedMessageID, 'active', FALSE);

        SET i = i + 1;
    END WHILE;
END //
DELIMITER ;
CALL CreateFriendshipsForRicky001();
DROP PROCEDURE CreateFriendshipsForRicky001;

-- 創建群組名單:
DROP PROCEDURE IF EXISTS CreateRickyGroupAndMembers;
DELIMITER //
CREATE PROCEDURE CreateRickyGroupAndMembers()
BEGIN
    DECLARE groupID VARCHAR(36);
    DECLARE userID VARCHAR(36);
    DECLARE i INT DEFAULT 2;

    -- 创建群组
    INSERT IGNORE INTO `groups` (`id`, `group_name`, `description`, `group_owner_id`)
    SELECT UUID(), 'ricky_group', 'This is Ricky Group', `id` FROM `users` WHERE `username` = 'ricky001';
    SET groupID = (SELECT id FROM `groups` WHERE `group_name` = 'ricky_group');

    -- 添加 ricky001 作为 owner
    INSERT IGNORE INTO `group_members` (`group_id`, `user_id`, `user_name`, `role`)
    SELECT groupID, `id`, `username`, 'owner' FROM `users` WHERE `username` = 'ricky001';

    -- 添加 ricky002 到 ricky010 作为 admin
    WHILE i <= 10 DO
        INSERT INTO `group_members` (`group_id`, `user_id`, `user_name`, `role`)
        SELECT groupID, `id`, `username`, 'admin' FROM `users` WHERE `username` = CONCAT('ricky', LPAD(i, 3, '0'));
        SET i = i + 1;
    END WHILE;

    -- 添加 ricky011 到 ricky050 作为 member
    WHILE i <= 50 DO
        INSERT INTO `group_members` (`group_id`, `user_id`, `user_name`, `role`)
        SELECT groupID, `id`, `username`, 'member' FROM `users` WHERE `username` = CONCAT('ricky', LPAD(i, 3, '0'));
        SET i = i + 1;
    END WHILE;
END //
DELIMITER ;
CALL CreateRickyGroupAndMembers();
DROP PROCEDURE CreateRickyGroupAndMembers;


select * from `groups`;

select * from `group_members`;