package main

import (
	"fmt"
	"net/http"
)

/*
    Go encourages the use of short, concise variable names, especially for local variables with a small scope.
	The length of a variable's name should reflect its scope—the shorter the scope, the shorter the variable name.


	Scope: A variable can be declared in three places:
		Function body: Local variable.
		Outside all functions: Global variable.
		In the function signature: Function parameters or return values.
*/

/*
	    Exported Variables:
		When a variable starts with a capital letter, it means the variable is exported (public).
		When it starts with a small letter, it's not exported (private).

	    Exporting a variable (or function, type, or method) means making it accessible from other packages.
		This is crucial when writing libraries or packages that are meant to be used in different contexts
		within your project or by other projects.

		Here are some reasons why you might want to use exported variables:

		Encapsulation:
		Exported variables allow you to control what parts of your package are accessible to the outside world.
		This is a fundamental concept of encapsulation in object-oriented programming. You can keep some variables unexported (private)
		to hide their details and expose only what is necessary via exported (public) variables.

		Code Organization:
		If you're creating a complex application, it might be useful to split your code into different packages.
		Exported variables can be used to share data between these packages.

		Creating Libraries:
		If you're writing a library to be used by other Go programs, you'll need to export the variables,
		functions, types, etc., that should be accessible to users of your library.

		Note: Be careful when using exported variables, as they can be changed by any package that imports them.
		In most cases, it's better to provide exported functions that control access to internal package data,
		rather than exporting the variables directly. This follows the principle of encapsulation and provides better
		control over your code.
*/
var GlobalVar string = "abc" // exported global variable

/*
	    Variable declaration block
		These can also be declared inside functions
*/
var (
	name string = "John"
	age  int    = 20
	city string = "San Francisco"
)

/*
	In Go variables always need to be used, ence all the Println statements
	There is an exception to the variables declared at the package level
*/

func main() {
	// Declaration: You can declare variables in Go using the var keyword followed by the variable name and the type.
	var i int      // variable declaration. You read it like: "declare a variable i of type integer"
	i = 42         // value assignment.
	i = 24         // change the value of i
	fmt.Println(i) // prints 24

	// Initialization: To assign a value to the variable at the time of declaration, you can use the = operator.
	// You should use this form of declaration when you require more control over the type of a variable. i.e floaf32 vs float64
	var j int = 19 // variable declaration and assignment. You read it like: "declare a variable j of type integer and assigned the value 19 to it"
	fmt.Println(j) // prints 19

	/*
		Short declaration: Go also supports short variable declarations using :=. This allows you to declare and initialize a variable without
		explicitly stating its type.
		Go will automatically infer the type based on the initial value
	*/
	k := 33        // short variable declaration. Here go will infer the variable type
	fmt.Println(k) // prints 33

	// Multiple variables: You can declare or initialize multiple variables at once.
	var a, b int = 11, 22
	fmt.Println(a) // prints 11
	fmt.Println(b) // prints 22

	// Zero values: Variables declared without an explicit initial value are given their zero value.
	// The zero value is 0 for numeric types, false for the boolean type, "" (the empty string) for strings,
	// and nil for pointers, functions, interfaces, slices, channels, and maps.
	var zero int
	fmt.Print(zero) // prints 0

	/*
	   Variable declaration block can also be declared inside functions
	*/
	var (
		isTrue bool   = true
		number int32  = 10
		word   string = "cake"
	)
	fmt.Println(isTrue)
	fmt.Println(number)
	fmt.Println(word)

	// Variable shadowing occurs when a variable declared within a certain scope shares the same name as a variable declared in an outer scope.
	// When the duplicate variable is used, it refers to the value in the innermost scope, effectively "shadowing" the outer variable.
	// While shadowing is not necessarily bad and sometimes can be useful, it can lead to bugs that are hard to track down, especially in larger programs.
	// It's generally considered a good practice to avoid shadowing to prevent unintended behavior.
	// Most Go linters will warn you about shadowed variables, and you can usually avoid shadowing by choosing different variable names or by reusing the
	// outer variable, if appropriate.
	var name string = "I'm shadowing the 'name' variable defined at the package level"
	fmt.Println(name)

	// Naming conventions

	// Use MixedCaps or mixedCase rather than underscores: In Go, variable names are usually written in MixedCaps or mixedCase.
	totalCount := 0
	HTTPRequest := ""
	fmt.Println(totalCount)
	fmt.Println(HTTPRequest)

	// Acronyms should be all uppercase: If a name includes an acronym like ID, HTTP, URL, it should be all uppercase.
	var userID string
	var httpURL string
	fmt.Println(userID)
	fmt.Println(httpURL)

	// If you have a variable like HTTPRequest and you want to make it private, it should be httpRequest
	// As the first letter of the identifier determines visibility (exported or unexported). This takes precedence over the convention for acronyms.
	var httpRequest *http.Request
	fmt.Println(httpRequest)

	// Use meaningful names: Variable names should be descriptive enough to indicate what they are used for.
	var customerName string
	var accountBalance float64
	fmt.Println(customerName)
	fmt.Println(accountBalance)
}
