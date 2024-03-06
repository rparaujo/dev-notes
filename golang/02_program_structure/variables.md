# Variables

Variable declarations

A variable declaration has the general form:
```go
var name type = expression
```

- Either the _type_ or the _= expression_ part can be omitted but not both.
- If the _type_ is omitted, it is inferred by the result of the _expression_.
- If the _expression_ is omitted, the variable is initialized to it's [_zero_ value](#zero-value)

This form of variable declaration is enforced in _package_ level variables.
In local variables (within function bodies), it can be used as to express that the variable has a type that differs from that of it's _initializer expression_, or when the variable will be assigned a value later in the execution and it's initial value is of no importance.

Declaring a string to it's zero-value by omitting the _= expression_ part. `s` holds the value `""`:
``` go
var s string
```

Declaring several variables in s single declaration. `i`, `j` and `k` are all of type `int` and their value is `0`:
```go
var i,j,k int
```

Declaring multiple variables of different types, and omitting the _type_ from the declaration:
- `b` has a type of `boolean` with a value of `true`
- `f` has a type of `float64` with a value of `1.1`
- `s` has a type of `string` with a value of `this is a string`
```
var b, f, s = true, 1.1, "this is a string"
```

Declaring multiple variables by calling a function that returns multiple values:
```go
var f, err = os.Open(file) // os.Open return a file and an error
```

Declare variables using [_short variable declaration_](#short-variable-declarations):
```go
s := "this is a string"
i := 1 + 1
f, err := os.Open(file)
```

## Zero value
The _zero-value_ mechanism ensures that a variable is always holds a well-defined value of it's type.

## Short variable declarations
_Short variable declarations_ might be used within a function to declare and initialize local variables.
It takes the form:
```go
name := expression
```
The type of _name_ is determined by the type of _expression_.

Declaring multiple variable already initialized:
```go
i, f := 1, 1.1
```

Declarations with multiple variable should only be used if they help with readability, such as for short and natural groupings, like in loops or functions.

Keep in mind that `:=` is a declaration, and `=` is an assignment. The key difference is that a declaration defines a new variable and an assignment dows not, it only assigns a value to an existing variable.

Keep in mind that variable assignment happens in two steps:
1. the right-hand side is evaluated, creating a temporary tuple with the results
2. the left-hand side is assigned the values of the tuple

This makes swapping variable very easy:
```go
i, j = j, i
```
This is not the case with all programming languages.

Another subtleness of using _short variable declarations_ is that it does not always create new variables. For example in the code:
```go
of, err := os.Open(file) // of and err are two nre variables
// ...
cf, err := os.Create(file) // here cf w variable, but err is not
```

If the second declaration the variable `cf` is a new variable in the scope, so a new variable is created, but `err` was created in the first declaration and so already exists in the scope.
This effect creates what is know as [shadowing](#TODO)

In a _short variable declaration_ at least one variable must be a new variable in the lexical scope.

If none of the variables are new, the code can be fixed with a simple assignment `=`