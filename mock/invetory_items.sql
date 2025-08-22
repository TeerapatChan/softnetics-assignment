CREATE TABLE IF NOT EXISTS inventory_items (
    id SERIAL PRIMARY KEY,
    product_name TEXT,
    status TEXT,
    price NUMERIC(10,2),
    amount INT,
    at TIMESTAMP,
);
