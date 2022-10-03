module KentrServer.com/server

go 1.19

replace KentrServer.com/myFunctions => ../myFunctions

require KentrServer.com/myFunctions v0.0.0-00010101000000-000000000000

require github.com/go-sql-driver/mysql v1.6.0 // indirect
