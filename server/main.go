package main

import (

"fmt"

"KentrServer.com/myFunctions"

"net/http"

//"os"   although not used yet

"log"

)

func Fibonacci10(w http.ResponseWriter, r *http.Request) {
	// return the 10th Fibonacci number in the response payload
	fmt.Fprintf(w, "%d", myFunctions.Fibonacci(10))
}


func main(){

//term:=myFunctions.Fibonacci(10)
//fmt.Println(term)

    addr := ":4000"
   
    //create a new mux (router)
    //the mux calls different functions for
    //different resource paths
    mux := http.NewServeMux()
    
    //tell it to call the Fibonacci10() function
    //when someone requests the resource path `/fibonacci10`
    
    mux.HandleFunc("/fibonacci10",Fibonacci10)
    
    
    
     //start the web server using the mux as the root handler,
    //and report any errors that occur.
    //the ListenAndServe() function will block so
    //this program will continue to run until killed
        
    log.Printf("server is listening at %s...", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
    
    
    
}
