CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
    quantity DECIMAL(10, 2) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    price DECIMAL(10, 2) NOT NULL
);
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    userId INTEGER NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    status VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    orderId INTEGER NOT NULL,
    productId INTEGER NOT NULL,
    quanity DECIMAL(10, 2) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    status VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

INSERT INTO products (name, description, image, quantity,price)
VALUES 
('Product 1', 'Description for Product 1', 'image1.jpg', 10.00, 10.00),
('Product 2', 'Description for Product 2', 'image2.jpg', 20.00, 20.00),
('Product 3', 'Description for Product 3', 'image3.jpg', 15.00, 15.00),
('Product 4', 'Description for Product 4', 'image4.jpg', 8.00, 8.00),
('Product 5', 'Description for Product 5', 'image5.jpg', 5.00, 5.00),
('Product 6', 'Description for Product 6', 'image6.jpg', 12.50, 12.50),
('Product 7', 'Description for Product 7', 'image7.jpg', 30.00, 30.00),
('Product 8', 'Description for Product 8', 'image8.jpg', 25.00, 25.00),
('Product 9', 'Description for Product 9', 'image9.jpg', 18.00, 18.00),
('Product 10', 'Description for Product 10', 'image10.jpg', 22.00, 22.00),
('Product 11', 'Description for Product 11', 'image11.jpg', 7.50, 7.50),
('Product 12', 'Description for Product 12', 'image12.jpg', 14.00, 14.00),
('Product 13', 'Description for Product 13', 'image13.jpg', 9.00, 9.00),
('Product 14', 'Description for Product 14', 'image14.jpg', 16.00, 16.00),
('Product 15', 'Description for Product 15', 'image15.jpg', 10.50, 10.50),
('Product 16', 'Description for Product 16', 'image16.jpg', 21.00, 21.00),
('Product 17', 'Description for Product 17', 'image17.jpg', 19.00, 19.00),
('Product 18', 'Description for Product 18', 'image18.jpg', 13.00, 13.00),
('Product 19', 'Description for Product 19', 'image19.jpg', 11.00, 11.00),
('Product 20', 'Description for Product 20', 'image20.jpg', 17.00, 17.00);