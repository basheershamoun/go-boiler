CREATE TABLE IF NOT EXISTS user(
    id int PRIMARY KEY AUTO_INCREMENT,
    created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    uuid varchar(180) NOT NULL,
    roles json NOT NULL,
    password varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    username varchar(55)
);
