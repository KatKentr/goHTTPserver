package myFunctions



import (

  "database/sql"
    //"fmt"
    "log"
    _"os"
    _"github.com/go-sql-driver/mysql"   //import the MySQL driver

)



func FetchData() string{


//database handle
    var db *sql.DB

// Capture connection properties.struct Config
/*
    cfg := mysql.Config{
        
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "fruits",
        AllowNativePasswords: true,
        
    }
    */
    
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", "root:1234_Ken@tcp(127.0.0.1:3306)/test")
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
     return "Connected!"
    
    





}
