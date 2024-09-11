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

INSERT INTO products (name, description, image, price, quantity) VALUES
    ('Dell XPS 13', 'A compact and powerful laptop with a sleek design.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 999.00, 15),
    ('Apple MacBook Pro 16', 'High-performance laptop with a stunning Retina display.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 2399.00, 5),
    ('HP Envy 13', 'A stylish and powerful laptop with a long battery life.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 899.00, 20),
    ('Lenovo Yoga C940', 'A versatile 2-in-1 laptop with a 4K touchscreen display.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 1299.00, 8),
    ('Asus ZenBook 14', 'A lightweight and portable laptop with a powerful processor.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 799.00, 25),
    ('Microsoft Surface Laptop 3', 'A premium laptop with a high-resolution PixelSense display.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 999.00, 10),
    ('HP Spectre x360', 'A versatile 2-in-1 laptop with high-end specs.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 1099.00, 18),
    ('Lenovo ThinkPad X1 Carbon', 'Durable and reliable with a high-resolution display.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 1399.00, 10),
    ('Asus ROG Zephyrus G14', 'High-performance gaming laptop with a compact form factor.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 1499.00, 8),
    ('Microsoft Surface Laptop 4', 'Stylish and powerful with a high-resolution touchscreen.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 1299.00, 14),
    ('Acer Predator Helios 300', 'Gaming powerhouse with top-of-the-line graphics.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 1399.00, 11),
    ('Razer Blade 15', 'Sleek design with excellent gaming performance.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 1799.00, 6),
    ('MSI GS66 Stealth', 'High-performance laptop with a stealthy design.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 1599.00, 9),
    ('Samsung Galaxy Book Pro 360', '2-in-1 convertible laptop with AMOLED display.', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 1299.00, 13),
    ('Razer Blade 15866', 'Description for Product 11', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 7.50, 10),
    ('Razer Blade 15dce24', 'Description for Product 12', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 14.00, 20),
    ('Razer Blade dfa15', 'Description for Product 13', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 9.00, 15),
    ('Razer Blade e15', 'Description for Product 14', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 16.00, 12),
    ('Razer Blade ddc15', 'Description for Product 15', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 10.50, 18),
    ('Razer Blade 15ee4657', 'Description for Product 16', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 21.00, 14),
    ('Razer Blade 15rr', 'Description for Product 17', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 19.00, 10),
    ('Razer Blade 15', 'Description for Product 18', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 13.00, 20),
    ('Razer Blade 15ee', 'Description for Product 19', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 11.00, 15),
    ('Razer Blade 15', 'Description for Product 20', 'iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAAC8GO2jAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHx9z8YAAAAASUVORK5CYII=', 17.00, 12);
