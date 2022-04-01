# Library API Project with PostgreSQL&GORM in GOLANG

This api project reads .json files and do basic database operations with GORM

```[terminal]
go run main.go
```

We have books struct object and the fields are;
```
- Book ID
- Book Title
- Number of Page
- Number of Stock
- Price
- Stock Code
- ISBN
- Author (ID ve Name)
```
We have authors struct object and the fields are;
```
- ID
- Name
```
## Examples

You can check endpoints with curl

```
curl -v localhost:8080/books
```
```
curl -v localhost:8080/authors
```
```
curl localhost:8080/books/1 -GET
```
```
curl localhost:8080/books/withauthorname/1
```

## Third Party packages

* The program is created with **GO main package & GORM & Godotenv & go-swagger & Gorilla-mux**.