package myFunctions



import (

  "database/sql"
    _"fmt"
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
    
   //database handle
    var db *sql.DB
    
 //returns a db handle   
func connectToDB() *sql.DB{


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
     
     return db

}


//retries a table from the db and returns a slice of structs,each struct contains a row of the table
func FetchData() ([]Fruitmix){


    db := connectToDB()   //invoke function to get db handle
   
       
    // A fruits slice to hold data from returned rows.
    var fruits []Fruitmix  
       
    
    rows, err := db.Query("SELECT * FROM dummyData")
    
    if err != nil {
	log.Fatal(err)
    }
    
    //Defer closing rows so that any resources it holds will be released when the function exits.
    defer rows.Close()
    
    defer db.Close()   //!Attention not sure if it is recommended to close the connection
    
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
    
        var fr Fruitmix
	err := rows.Scan(&fr.id, &fr.a, &fr.b,&fr.c,&fr.d)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(fr.id,fr.a,fr.b,fr.c,fr.d)
	
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
