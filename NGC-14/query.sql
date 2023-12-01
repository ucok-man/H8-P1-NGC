-- select product based on category price
SELECT 
	COUNT(price),
    CASE
    	WHEN price < 300 THEN 'Low'
        WHEN price BETWEEN 300 and 600 THEN 'Medium'
        ELSE 'High'
	END AS category
FROM products
GROUP BY category;


-- Write a SQL query to retrieve the top N products based on the total quantity sold. 
-- The result should include the product name and the total quantity sold for each product, 
-- ordered from highest to lowest total quantity sold. 
-- Set the value of N to any positive integer of your choice 
SELECT
	products.product_name,
    SUM(product_sales.quantity_sold) AS "product sold"
FROM 
    products
INNER JOIN 
    product_sales USING (product_id)
GROUP BY 
    product_sales.product_id
ORDER BY 
    "product sold" DESC
LIMIT 10;

-- Write a SQL query to calculate the percentage growth in sales for each product over a specific time period. 
-- The result should include the product name, the total quantity sold for the current period, and the total quantity sold for the previous period. 
-- Assume the previous period is one month ago (from the current date). 
-- If there are no sales in the previous period, consider it as 0. 
-- The percentage growth should be rounded to two decimal places.
SELECT
    p.product_name,
    SUM(
        CASE 
        	WHEN ps.sale_date = CURDATE() THEN ps.quantity_sold 
        	ELSE 0
        END
    ) AS current_period_sale,
    
    SUM(
        CASE 
        	WHEN ps.sale_date = DATE_SUB(CURRENT_DATE(), INTERVAL 1 DAY) AND ps.sale_date < CURDATE() THEN ps.quantity_sold 
        	ELSE 0 
        END
    ) AS previous_period_sale
FROM products p
LEFT JOIN product_sales ps ON p.product_id = ps.product_id
GROUP BY p.product_id, p.product_name;


-- Soal 4
SELECT 
	p1.ProductID,
    p1.ProductName,
    p1.Price,
    AVG(p1.Price - p2.Price) AS avg_different
FROM
	products p1
JOIN
	products p2 ON p1.ProductID = p2.ProductID + 1
GROUP BY
	p1.ProductID,
    p1.ProductName,
    p1.Price;

