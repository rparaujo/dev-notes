[Previous](./names.md) | [Home](../README.md#environment-setup) | [Next](#TODO)

# Declarations

There are four kinds of declarations: [var](#TODO), [const](#TODO), [type](#TODO) and [func](#TODO)

[var](#TODO) and [const](#TODO) can be declared inside [functions](#TODO) are are said to be local to that function.
In any entity is declared outside of a function, i.e at the [package](#TODO) level, it is said to be visible to all the files within the same package.
NAmes that begin with upper-case are _exported_ which means that they can be accessed beyond the scope of the package that defines them.

Go programs are written in one or more text files that have the _.go_ extension.
Each file begins with it's package declaration followed by any import declarations and then any package level declarations that might include, types, variables, constants and functions. These declarations can be in any order.

Example:

```go
// Print's the string "Hello World!!"
package main

import fmt

const sayHello = "Hello World!!"

func main() {
    fmt.Println(sayHello)
}

```