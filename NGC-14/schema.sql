create table products (
    product_id int primary key,
    product_name varchar(100),
    product_code varchar(20),
    price decimal(10,2)
);

insert into products(product_id, product_name, product_code, price)
values
    (1, 'Laptop', 'LT001', 799.99),
    (2, 'Smartphone', 'SP002', 499.99),
    (3, 'Tablet', 'TB003', 299.99),
    (4, 'Headphone', 'HD004', 149.99),
    (5, 'Monitor', 'MN005', 399.99);


create table product_sales(
    sale_id int primary key,
    product_id int,
    sale_date date,
    quantity_sold int,

    foreign key (product_id) references products(product_id)
);

insert into product_sales(sale_id, product_id, sale_date, quantity_sold)
values
    (1, 1, '2023-08-01', 10),
    (2, 1, '2023-08-02', 5),
    (3, 2, '2023-08-02', 8),
    (4, 3, '2023-08-03', 12),
    (5, 3, '2023-08-03', 6),
    (6, 4, '2023-08-04', 20),
    (7, 5, '2023-08-04', 15);
