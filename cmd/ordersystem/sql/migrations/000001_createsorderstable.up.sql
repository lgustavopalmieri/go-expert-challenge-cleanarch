-- init_schema.sql

CREATE TABLE orders (
    order_id VARCHAR(50) PRIMARY KEY,
    price NUMERIC,
    tax NUMERIC,
    final_price NUMERIC, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO orders (order_id, price, tax, final_price, created_at) VALUES ('01e68602-9593-4cc6-bdea-30fdbbd8c75c', 333, 33, 366, '2024-04-17T16:04:13');