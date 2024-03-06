Names in Go are used to identify `functions`, `variables`, `constants` `types`, `statement labels` and `packages`.
A name begins with a letter and can have any number os additional letters, numbers and underscores.
Names are case sensitive, i.e `count` is different from `Count`.


Go has 25 keywords that cannot be used as names:

[break]()
[case]()
[chan]()
[const]()
[continue]()
[default]()
[defer]()
[else]()
[fallthrough]()
[for]()
[func]()
[go]()
[goto]()
[if]()
[import]()
[interface]()
[map]()
[package]()
[range]()
[return]()
[select]()
[struct]()
[switch]()
[type]()
[var]()

Go has several predeclared names that are built into the language and are always available without the need to import a package. These names include [types](), [constants](), and [functions](). 

### Types
[bool]()
[byte (alias for uint8)]()
[complex64]()
[complex128]()
[error]()
[float32]()
[float64]()
[int]()
[int8]()
[int16]()
[int32]()
[int64]()
[rune (alias for int32)]()
[string]()
[uint]()
[uint8]()
[uint16]()
[uint32]()
[uint64]()
[uintptr]()

### Constants
[true]()
[false]()
[iota]()
[nil]() - The predeclared identifier. Used to represent zero values for pointers, interfaces, maps, slices, channels, and function types.

### Functions
[append]()
[cap]()
[close]()
[complex]()
[copy]()
[delete]()
[imag]()
[len]()
[make]()
[new]()
[panic]()
[print]()
[println]()
[real]()
[recover]()

These predeclared names are **not** reserved and may be used in declarations.
Names in Go are not limited by size, but it is considered idiomatic to adjust the length of the name to the scope of the variable. The bigger the scope the bigger the name.
It is also idiomatic to use `camelCase` when forming names.

[Previous](../01_tooling/make.md) | [Home](../README.md#environment-setup) | [Next](#TODO)