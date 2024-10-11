-- +goose Up
CREATE TABLE expenses (
    id          SERIAL        PRIMARY KEY,
    user_id     INT           NOT NULL,
    amount      NUMERIC(12,2) NOT NULL,
    category    VARCHAR(100)  NOT NULL,
    description VARCHAR       NOT NULL,
    added_at    TIMESTAMP     NOT NULL DEFAULT NOW(),
    FOREIGN KEY(user_id)   REFERENCES users(id),
    FOREIGN KEY(category)  REFERENCES categories(name)
);

-- +goose Down
DROP TABLE expenses;

