package myFunctions



import (

  "database/sql"
    "fmt"
    "log"
    _"os"
    _"github.com/go-sql-driver/mysql"   //import the MySQL driver

)

 //we will use this struct to hold row data returned from the query
    type Fruitmix struct {
    
           id int
           a int
           b string
           c string
           d string
    }



func FetchData() ([]Fruitmix){


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
     
   
       
    // A fruits slice to hold data from returned rows.
    var fruits []Fruitmix  
       
    
    rows, err := db.Query("SELECT * FROM dummyData")
    
    if err != nil {
	log.Fatal(err)
    }
    
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
    
        var fr Fruitmix
	err := rows.Scan(&fr.id, &fr.a, &fr.b,&fr.c,&fr.d)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(id,a,b,c,d)
	
	fruits=append(fruits,fr)
   }
   
   err = rows.Err()
   if err != nil {
	log.Fatal(err)
   }
   
   //myVar:=fruits
   //fmt.Printf("Data: %v\n", myVar)
        
   return fruits




}
