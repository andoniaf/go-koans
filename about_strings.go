package go_koans

import "fmt"

func aboutStrings() {
	assert("a"+"bc"  == "abc") // string concatenation need not be difficult
	assert(len("abc") == 3)   // and bounds are thoroughly checked

	assert("abc"[0] == 'a') // their contents are reminiscent of C

	// https://golang.org/ref/spec#Slice_expressions
	assert("smith"[2:] == "ith")  // slicing may omit the end point
	assert("smith"[:4] == "smit")  // or the beginning
	assert("smith"[2:4] == "it") // or neither
	assert("smith"[:] == "smith")   // or both

	assert("smith" == __string__) // they can be compared directly
	assert("smith" < __string__)  // i suppose maybe this could be useful.. someday

	bytes := []byte{'a', 'b', 'c'}
	assert(string(bytes) == __string__) // strings can be created from byte-slices

	bytes[0] = 'z'
	assert(string(bytes) == __string__) // byte-slices can be mutated, although strings cannot

	assert(fmt.Sprintf("hello %s", __string__) == "hello world") // our old friend sprintf returns
	assert(fmt.Sprintf("hello \"%s\"", "world") == __string__)   // quoting is familiar
	assert(fmt.Sprintf("hello %q", "world") == __string__)       // although it can be done more easily

	assert(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589) == __string__) // "the root of all evil" is actually a misquotation, by the way
}
