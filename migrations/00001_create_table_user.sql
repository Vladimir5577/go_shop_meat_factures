-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150),
    password VARCHAR(200) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    address VARCHAR(255),
    is_active BOOlEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT date_trunc('second', now()),
    updated_at TIMESTAMP DEFAULT date_trunc('second', now())
);

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) UNIQUE NOT NULL,
    description TEXT,
    is_active BOOlEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT date_trunc('second', now()),
    updated_at TIMESTAMP DEFAULT date_trunc('second', now())
);

CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price numeric CHECK (price > 0),
    category_id INT,
    in_stock BOOlEAN DEFAULT TRUE,
    is_active BOOlEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT date_trunc('second', now()),
    updated_at TIMESTAMP DEFAULT date_trunc('second', now()),

    FOREIGN KEY (category_id) REFERENCES category(id)
        ON DELETE SET NULL
);


CREATE TABLE order (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    order_products INT,
    total_summ NUMERIC CHECK (amount > 0),
    status VARCHAR(20) CHECK (status IN ('new', 'confirmed', 'processing', 'processed', 'sended', 'received')),
    comment TEXT,
    is_active BOOlEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT date_trunc('second', now()),
    updated_at TIMESTAMP DEFAULT date_trunc('second', now()),

    FOREIGN KEY (user_id) REFERENCES user(id)
        ON DELETE CASCADE
);


CREATE TABLE order_products (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    amount INT CHECK (amount > 0),
    total_summ NUMERIC CHECK (amount > 0),

    created_at TIMESTAMP DEFAULT date_trunc('second', now()),

    FOREIGN KEY (order_id) REFERENCES order(id)
        ON DELETE CASCADE,

    FOREIGN KEY (product_id) REFERENCES product(id)
        ON DELETE CASCADE,
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_products;
DROP TABLE IF EXISTS order;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS category;
DROP TABLE IF EXISTS user;
-- +goose StatementEnd
