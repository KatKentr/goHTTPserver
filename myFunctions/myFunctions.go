package myFunctions

// O(2^n) Fibonacci
func Fibonacci(n int) int {
	// F(0) = 0
	if n == 0 {
	
		return 0
	// F(1) = 1	
	} else if n==1 {
	
	        return 1
	
	} else {
	        // F(n) = F(n-1) + F(n-2) - return the n-th Fibonacci number
	        return Fibonacci(n-1) + Fibonacci(n-2)
	
	}
	

}
