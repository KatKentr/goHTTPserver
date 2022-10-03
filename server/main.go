package main

import (

"fmt"

"KentrServer.com/myFunctions"

"net/http"

//"os"   although not used yet

"log"

)

//returns the actual HTTP handler function
func Fibonacci(number int) func(w http.ResponseWriter, r *http.Request) {

//error handling sould be normally added: if number ==nil ...

  return func(w http.ResponseWriter, r *http.Request){
  
	// return the 10th Fibonacci number in the response payload
	
	fmt.Fprintf(w, "%d", myFunctions.Fibonacci(number))
  }

}



func FetchDB() func(w http.ResponseWriter, r *http.Request){
   
   //error handling sould be normally added: if number ==nil ...

  return func(w http.ResponseWriter, r *http.Request){
        
        //retrieve table data
        myData := myFunctions.FetchData()
        
        for _,value := range myData {
	// return each record from the table in the response payload
	   fmt.Fprintf(w, "%+v\n", value)
	//w.Write([]byte(myFunctions.FetchData()))
	}
	
  }

}


func main(){

//term:=myFunctions.Fibonacci(10)
//fmt.Println(term)

    //fmt.Println(myFunctions.Hello("Kat"))


    addr := ":4000"
   
    //create a new mux (router)
    //the mux calls different functions for
    //different resource paths
    mux := http.NewServeMux()
    
    //tell it to call the Fibonacci(number) function
    //when someone requests the resource path `/fibonacci10` or `/fibonacci20`  or `/fibonacci30`
    mux.HandleFunc("/fibonacci10",Fibonacci(10))
    mux.HandleFunc("/fibonacci20",Fibonacci(20))
    mux.HandleFunc("/fibonacci30",Fibonacci(30))
    mux.HandleFunc("/fetchDB_test",FetchDB())
     
    
    
    
     //start the web server using the mux as the root handler,
    //and report any errors that occur.
    //the ListenAndServe() function will block so
    //this program will continue to run until killed        
    log.Printf("server is listening at %s...", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
    
    
    
}
