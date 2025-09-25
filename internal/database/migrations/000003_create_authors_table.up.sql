CREATE TABLE IF NOT EXISTS authors (
    id bigint NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    biography varchar(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
)