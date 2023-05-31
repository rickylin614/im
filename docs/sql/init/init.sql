
-- 1. Users table
CREATE TABLE users (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 2. User_profiles table
CREATE TABLE user_profiles (
    user_id INT UNSIGNED PRIMARY KEY,
    avatar_url VARCHAR(255) DEFAULT NULL,
    display_name VARCHAR(255) DEFAULT NULL,
    bio TEXT DEFAULT NULL,
    location VARCHAR(255) DEFAULT NULL,
    website VARCHAR(255) DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- 3. Friends table
CREATE TABLE friends (
    user_id1 INT UNSIGNED NOT NULL,
    user_id2 INT UNSIGNED NOT NULL,
    friendship_status ENUM('active', 'blocked') NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id1, user_id2),
    -- FOREIGN KEY (user_id1) REFERENCES users (id) ON DELETE CASCADE,
    -- FOREIGN KEY (user_id2) REFERENCES users (id) ON DELETE CASCADE
);

-- 4. Friend_requests table
CREATE TABLE friend_requests (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sender_id INT UNSIGNED NOT NULL,
    receiver_id INT UNSIGNED NOT NULL,
    request_status ENUM('pending', 'accepted', 'rejected', 'canceled') NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (sender_id, receiver_id),
    -- FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE CASCADE,
    -- FOREIGN KEY (receiver_id) REFERENCES users (id) ON DELETE CASCADE
);

-- 5. Groups table
CREATE TABLE groups (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT DEFAULT NULL,
    group_owner_id INT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- FOREIGN KEY (group_owner_id) REFERENCES users (id) ON DELETE CASCADE
);

-- 6. Group_members table
CREATE TABLE group_members (
    group_id INT UNSIGNED NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    role ENUM('owner', 'admin', 'member') NOT NULL,
    joined_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (group_id, user_id),
    -- FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE,
    -- FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- 7. Group_invitations table
CREATE TABLE group_invitations (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    group_id INT UNSIGNED NOT NULL,
    inviter_id INT UNSIGNED NOT NULL,
    invitee_id INT UNSIGNED NOT NULL,
    invitation_status ENUM('pending', 'accepted', 'rejected', 'canceled') NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (group_id, invitee_id),
    -- FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE,
    -- FOREIGN KEY (inviter_id) REFERENCES users (id) ON DELETE CASCADE,
    -- FOREIGN KEY (invitee_id) REFERENCES users (id) ON DELETE CASCADE
);

-- 8. Group_requests table
CREATE TABLE group_requests (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    group_id INT UNSIGNED NOT NULL,
    requester_id INT UNSIGNED NOT NULL,
    request_status ENUM('pending', 'accepted', 'rejected', 'canceled') NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY (group_id, requester_id),
    -- FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE,
    -- FOREIGN KEY (requester_id) REFERENCES users (id) ON DELETE CASCADE
);