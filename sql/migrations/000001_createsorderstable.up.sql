-- init_schema.sql

CREATE TABLE orders (
    order_id VARCHAR(50) PRIMARY KEY,
    price NUMERIC,
    tax NUMERIC,
    final_price NUMERIC, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
