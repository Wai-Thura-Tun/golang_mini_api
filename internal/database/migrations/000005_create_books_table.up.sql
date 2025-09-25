CREATE TABLE IF NOT EXISTS books (
    id bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(150) NOT NULL,
    description varchar(255) NOT NULL,
    type varchar(150) NOT NULL,
    book_cover varchar(255) NOT NULL,
    author_id bigint NOT NULL,
    category_id bigint NOT NULL,
    rating bit(3) NOT NULL,
    price double(10,3) NOT NULL,
    isSpecial boolean NOT NULL DEFAULT 0,

    FOREIGN KEY (author_id) REFERENCES authors(id)
    FOREIGN KEY (category_id) REFERENCES category(id)
)