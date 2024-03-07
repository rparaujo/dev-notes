package main

import (
	"fmt"
	"strconv"
)

var (
	i32 int32
	f32 float32
)

func main() {

	// In go conversions need to be explicit
	// convert float to int
	var f float32 = 3.14
	i := int(f)
	i32 := int32(f)
	i64 := int64(f)

	// string conversions
	i_42 := 42
	fmt.Println(i_42)               // prints '*' because 42 is the decimal value for the glyph *. In go we need to do explicif conversions, even for strings.
	fmt.Println(strconv.Itoa(i_42)) // prints '42'
}
