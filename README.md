### TEST (MACOS)
```plaintext
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

### dump database
```plaintext
docker exec -it mysql_container mysqldump -uroot -proot1234 --databases shopdevgo --add-drop-database --add-drop-table --add-drop-trigger --add-locks --no-data > migrations/shopdevgo.sql
```