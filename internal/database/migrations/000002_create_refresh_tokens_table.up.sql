CREATE TABLE IF NOT EXISTS refresh_tokens (
    id bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id bigint NOT NULL,
    token varchar(255) NOT NULL,
    expires_at datetime NOT NULL,
    revoked boolean NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX idx_refresh_tokens_token (token)
)