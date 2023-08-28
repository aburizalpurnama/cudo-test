SELECT 
	DATE_FORMAT(o.OrderDate, "%Y-%m") Bulan, 
	c.CategoryName CategoryName,
    SUM(od.Quantity) QtyItem
FROM
	Orders o
    JOIN OrderDetails od ON o.OrderID = od.OrderID
    JOIN Products p ON od.ProductID = p.ProductID
    JOIN Categories c ON p.CategoryID = c.CategoryID
GROUP BY
	DATE_FORMAT(o.OrderDate, "%Y-%m"),
	c.CategoryName