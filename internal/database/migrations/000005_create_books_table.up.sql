CREATE TABLE IF NOT EXISTS books (
    id bigint NOT NULL AUTO_INCREMENT,
    name varchar(150) NOT NULL,
    overview text NOT NULL,
    type varchar(150) NOT NULL,
    cover varchar(255) NOT NULL,
    author_id bigint NOT NULL,
    category_id bigint NOT NULL,
    rating tinyint unsigned NOT NULL,
    price double(10,3) NOT NULL,
    isSpecial boolean NOT NULL DEFAULT 0,

    PRIMARY KEY (id),
    FOREIGN KEY (author_id) REFERENCES authors(id),
    FOREIGN KEY (category_id) REFERENCES category(id)
)