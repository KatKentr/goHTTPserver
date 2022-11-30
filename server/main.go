package main

import (

"fmt"

"KentrServer.com/myFunctions"

"net/http"

//"os"   //although not used yet

"log"

"database/sql"

"flag"
"html/template"

)

//parse command line to check which test case it is(maybe not needed though). TO DO: Consider error handling in case the database server is down

type FibFields struct {

   Num,Term int

}



//returns the actual HTTP handler function
func Fibonacci(number int) func(w http.ResponseWriter, r *http.Request) {

//error handling sould be normally added: if number ==nil ...
  var vars=FibFields{Num: number, Term: myFunctions.Fibonacci(number)}

  return func(w http.ResponseWriter, r *http.Request){
  
        //add headers
         w.Header().Set("Server", "Go/1.19.1 (Ubuntu)")
         w.Header().Set("Connection", "keep-alive")
  
	// return the 10th Fibonacci number in the response payload
	
	//fmt.Fprintf(w, "The %dth term of the fibonacci sequence is: %d\n", number,myFunctions.Fibonacci(number))
	 t, _ := template.ParseFiles("fib_page.html")
	 //Term :=myFunctions.Fibonacci(number)
         t.Execute(w, vars)
  }

}



func FetchDB(db *sql.DB) func(w http.ResponseWriter, r *http.Request){
   
   //error handling sould be normally added: if number ==nil ...

  return func(w http.ResponseWriter, r *http.Request){
        
        //retrieve table data
        myData := myFunctions.FetchData(db)
        
        for _,value := range myData {
	// return each record from the table in the response payload
	   fmt.Fprintf(w, "%+v\n", value)
	//w.Write([]byte(myFunctions.FetchData()))
	}
	
  }

}


func staticFileHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./../serve_image/100kB_image.png")
}


func main(){



    //fmt.Println(myFunctions.Hello("Kat"))

    testcase := flag.String("testcase","","test-case: image or fibonacci or fetchDB")
    //testcase := os.Args[1]
    
    //fmt.Println("Default value of cmd argument is:",*testcase)
    
    flag.Parse()
    
    fmt.Println("value of testcase flag is:",*testcase)

    addr := ":4000"
    
    var db *sql.DB
    
    
    if *testcase=="dbTest" {
        
        //retrieve a reference to db handle
          db = myFunctions.ConnectToDB()
     } else {
         
         db = nil
     
     }
         
        
    
    
  
    //create a new mux (router)
    //the mux calls different functions for
    //different resource paths
    mux := http.NewServeMux()
    
    //tell it to call the Fibonacci(number) function
    //when someone requests the resource path `/fibonacci10` or `/fibonacci20`  or `/fibonacci30`
    mux.HandleFunc("/fibonacci10",Fibonacci(10))
    mux.HandleFunc("/fibonacci20",Fibonacci(20))
    mux.HandleFunc("/fibonacci30",Fibonacci(30))
    mux.HandleFunc("/fetchDB_test",FetchDB(db))
    mux.HandleFunc("/fetch_image",staticFileHandler)
     
    
    
    
     //start the web server using the mux as the root handler,
    //and report any errors that occur.
    //the ListenAndServe() function will block so
    //this program will continue to run until killed        
    log.Printf("server is listening at %s...", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
    
    
    
}
