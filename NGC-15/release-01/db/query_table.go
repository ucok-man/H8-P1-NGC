package db

const (
	table_employee_query = `
		CREATE TABLE IF NOT EXISTS employees(
			emp_id int AUTO_INCREMENT PRIMARY KEY,
			emp_firstname varchar(255) NOT NULL,
			emp_lastname varchar(255) NOT NULL,
			emp_position varchar(255) NOT NULL
		);
	`

	table_menu_items_query = `
		CREATE TABLE IF NOT EXISTS menu_items(
			mi_id int AUTO_INCREMENT PRIMARY KEY,
			mi_name varchar(255) NOT NULL,
			mi_description varchar(255) NOT NULL,
			mi_price decimal(11,2) NOT NULL,
			mi_category varchar(255) NOT NULL,
			
			CHECK(mi_price > 0)
		);
	`

	table_orders_query = `
		CREATE TABLE IF NOT EXISTS orders(
			order_id int AUTO_INCREMENT PRIMARY KEY,
			emp_id int NOT NULL,
			order_table_number int NOT NULL,
			order_date date NOT NULL,
			order_status varchar(255) NOT NULL,
			
			CHECK	(order_table_number > 0),
			FOREIGN KEY (emp_id) REFERENCES employees(emp_id)
		); 
	`

	table_orders_item_query = `
		CREATE TABLE IF NOT EXISTS order_items (
			oi_id int AUTO_INCREMENT PRIMARY KEY,
			order_id int NOT NULL,
			mi_id int NOT NULL,
			oi_quantity int NOT NULL,
			oi_subtotal decimal(14,2) NOT NULL,
			
			CHECK (oi_quantity > 0 AND oi_subtotal > 0),
			FOREIGN KEY (order_id) REFERENCES	orders(order_id),
			FOREIGN KEY (mi_id) REFERENCES menu_items(mi_id)
		);
	`

	table_payment_query = `
		CREATE TABLE IF NOT EXISTS payments(
			payment_id int AUTO_INCREMENT PRIMARY KEY,
			order_id int NOT NULL,
			payment_method varchar(255) NOT NULL,
			payment_date date NOT NULL,
			payment_total_amount int NOT NULL,
			
			CHECK(payment_total_amount > 0),
			FOREIGN KEY (order_id) REFERENCES orders(order_id)
		);
	`
)
