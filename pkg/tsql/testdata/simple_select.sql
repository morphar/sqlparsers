SELECT pp.LastName FROM Person.Person;

SELECT some_type::static_method (@arg1) FROM Product;

SELECT still_one_type.non_static_method (@but, @with, @params) FROM Product;

SELECT pp.LastName, pp.FirstName, e.JobTitle 
INTO dbo.EmployeeOne
FROM Person.Person AS pp JOIN HumanResources.Employee AS e
ON e.BusinessEntityID = pp.BusinessEntityID
WHERE LastName = 'Johnson';

SELECT p.Name AS ProductName, 
NonDiscountSales = (OrderQty * UnitPrice),
Discounts = ((OrderQty * UnitPrice) * UnitPriceDiscount)
FROM Production.Product AS p 
INNER JOIN Sales.SalesOrderDetail AS sod
ON p.ProductID = sod.ProductID 
ORDER BY ProductName DESC;

USE AdventureWorks2012;
GO
SELECT TOP 10.5 PERCENT *
FROM Production.Product
ORDER BY Name ASC;

SELECT * FROM @x;
