package hw7

import (
	// "github.com/UofSC-Fall-2022-Math-587-001/homework7/library"
)

func MillerRabinTest(N, a int) string {
	// Implement the Miller-Rabin Test found in Table 3.2 of the 
	// text. The pseudo-code is copied below for convenience 
	// Input. Integer n to be tested, integer a as potential 
	// witness. 
	// 1. If n is even or 1 < gcd(a,n) < n, return Composite
	// 2. Write n-1 = 2^k q with q odd
	// 3. Set a = a^q mod n. 
	// 4. If a = 1 mod n, return Test Fails.
	// 5. Loop i = 0,1,2,...,k-1
	//	6. If a = -1 mod n, return Test Fails. 
	//	7. Set a = a^2 mod n 
	// 8. End i loop.
	// 9. Return Composite
	return "Composite"
	//return "Test Fails"
}

