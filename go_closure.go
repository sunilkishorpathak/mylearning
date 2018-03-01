package main

/*
* What closure is?
*	To understand closure you need to know anonymus function, function, scope of a variable in function
*	1. Anonymus function : As like if you write the following code
*
*	func incr() func() int {
		return func() int {
			return 5
		}
	}
*	Calling incr() function, will create  a anonymus function. If you store the return value in sum variable and call it, you will get 5.
	import "fmt"

	func incr() func() int {
		return func() int {
			return 5
		}
	}

	func main() {
		incrRet := incr()

		fmt.Println(incrRet())
	}

	Now coming to second part, scope of the variable in a function. If there is function having variable, when function is called
	local variables are initialized, used and on return on destroyed. I.t. life of the local variable is tied to the start of the
	function and end of the function.

	Scope, scope of the local variable is tied to the scope of the local block, if variable is not found in local block, upper block
	variable will be used, if not, go in upper block and so on. If the variable is not in any block inside the function it will be error.

	import "fmt"

	func foo() {
		i := 1
		//j := 2 //compilation error, that it's not getting used, because j is used in printf, but local j has been used
		//uncomment the above j to see the error
		{
			j := 5 //this j is getting used
			fmt.Printf("Value of i, j and k is : %d, %d \n", i, j)
		}

	}

	func main() {
		foo()
	}

	Now combine the above two concept 1. that a function can return an anonymus function, and 2. scope of the variable, 3. life of variable.
	In closure, the return function will control the life of the lcoal variable of the function which is returning that i.e.
		closing the scope of the variable


import "fmt"

func nextInt() func() int {
	i := 0

	return func() int {
		i = i + 1
		return i
	}
}

func main() {
	nextIntClosure := nextInt()

	a := nextIntClosure()
	b := nextIntClosure()
	c := nextIntClosure()

	fmt.Printf("a, b, c values should be 1,2 and 3. Actualbe valures are : %d, %d, %d", a, b, c)

}
*/


import "fmt"

func nextInt() func() int {
	i := 0

	return func() int { //1. this anonumus function will be returned on calling of next int, with closed over local variable i = 0
						//2. It has closed over the above, so ccalling it will keep incrementing
		i = i + 1
		return i
	}
}

func main() {
	nextIntClosure := nextInt()

	a := nextIntClosure()
	b := nextIntClosure()
	c := nextIntClosure()

	fmt.Printf("a, b, c values should be 1,2 and 3. Actualbe valures are : %d, %d, %d", a, b, c)

}