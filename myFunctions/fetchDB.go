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


    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", "root:1234_Ken@tcp(127.0.0.1:3306)/fruits")
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
     
    
    var (
    
           id int
           a int
           b string
           c string
           d string
    )
       
    
    rows, err := db.Query("SELECT * FROM dummyData")
    
    if err != nil {
	log.Fatal(err)
    }
    
    defer rows.Close()
    
    for rows.Next() {
	err := rows.Scan(&id, &a, &b,&c,&d)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id,a,b,c,d)
   }
   
   err = rows.Err()
   if err != nil {
	log.Fatal(err)
   }
    
    
    
    
   return "Connected!"




}
