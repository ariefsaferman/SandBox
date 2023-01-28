-- create the tables
CREATE TABLE customers (
	customer_id SERIAL PRIMARY KEY,
	name VARCHAR(256) NOT NULL,
	phone_number VARCHAR(128) NOT NULL,
	email VARCHAR(256) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE purchase_headers (
	purchase_id SERIAL PRIMARY KEY,
	invoice_no VARCHAR(64) UNIQUE NOT NULL,
	purchase_date TIMESTAMP NOT NULL,
	customer_id INTEGER NOT NULL,
	FOREIGN KEY (customer_id) REFERENCES customers(customer_id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE item_categories (
	item_category_id SERIAL PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE items (
	item_id SERIAL PRIMARY KEY,
	item_category_id INTEGER,
	name VARCHAR(128) NOT NULL,
	price DECIMAL(12,2) NOT NULL,
	quantity INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL,
	FOREIGN KEY (item_category_id) REFERENCES item_categories (item_category_id)
);

CREATE TABLE purchase_details (
	purchase_detail_id SERIAL PRIMARY KEY,
	purchase_id INTEGER NOT NULL,
	item_id INTEGER NOT NULL,
	price DECIMAL(12,2) NOT NULL,
	quantity INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL,
	FOREIGN KEY (item_id) REFERENCES items(item_id),
	FOREIGN KEY (purchase_id) REFERENCES purchase_headers(purchase_id)
);

-- populating the data
INSERT INTO item_categories(name)
VALUES
('makanan'), ('minuman'), ('obat'), ('baterai'), ('perlengkapan mandi'), ('voucher');


INSERT INTO items (name, item_category_id, price, quantity)
VALUES
('shampo 150 ml', 5, 15000, 100),
('sabun mandi 100 ml', 5, 5000, 50),
('obat sakit kepala', 5, 5000, 0),
('obat influenza', 3, 8000, 1000),
('obat cacing', 3, 15000, 200),
('susu coklat', 2, 5600, 200),
('susu vanila', 2, 5500, 50),
('permen', 1, 200, 1000),
('keripik kentang 200 gram', 1, 20000, 100),
('mie instan rasa ayam', 1, 2000, 200);

INSERT INTO customers (name, phone_number, email)
VALUES
('Amir', '08123456789', 'amir@ini.example.com'),
('Budi', '08123456790', 'budi@ini.example.com'),
('Caca', '08123456791', 'caca@ini.example.com');

INSERT INTO purchase_headers (invoice_no, purchase_date, customer_id)
VALUES
('ABC/20220101/TOKO_A', '2022-01-01', 1),
('ABC/20220202/TOKO_A', '2022-02-02', 2),
('ABC/20220203/TOKO_B', '2022-02-03', 1);

INSERT INTO purchase_details(purchase_id, item_id, price, quantity)
VALUES
(1, 1, 14000, 5),
(1, 2, 6000, 15),
(2, 5, 6000, 10),
(2, 6, 6000, 20),
(3, 1, 5000, 5),
(3, 5, 6000, 10);


