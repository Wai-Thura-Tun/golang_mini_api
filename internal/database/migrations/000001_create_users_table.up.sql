CREATE TABLE IF NOT EXISTS users (
    id bigint NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    email varchar(150) NOT NULL,
    password varchar(255) NOT NULL,
    is_active boolean NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    UNIQUE INDEX idx_users_email (email)
);