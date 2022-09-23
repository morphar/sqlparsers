-- TSQL allows a trailing comma in table definition as per https://stackoverflow.com/a/20961244/700471
CREATE TABLE dbo.T1 
(
    column_1 int IDENTITY, 
    column_2 uniqueidentifier,
);