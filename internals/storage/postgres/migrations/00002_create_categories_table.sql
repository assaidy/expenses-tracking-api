-- +goose Up
create table categories(
    name VARCHAR(100) PRIMARY KEY
);

INSERT INTO categories (name)
VALUES 
    ('Groceries'),
    ('Leisure'),
    ('Electronics'),
    ('Utilities'),
    ('Clothing'),
    ('Health'),
    ('Others');

-- +goose Down
DROP TABLE categories;
