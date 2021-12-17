package main

import "fmt"

/*
Boolean type
Numeric Type:
Integers
Float
Complex Numbers
Text types
*/
func main() {
	var n bool = false
	fmt.Printf("%v,%T \n", n, n)

	x := 1 == 1
	y := 1 == 2
	fmt.Printf("%v,%T \n", x, x)
	fmt.Printf("%v,%T \n", y, y)

	var defaultValue bool
	fmt.Printf("%v,%T \n", defaultValue, defaultValue)

	fmt.Println("====================")

	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))

	fmt.Println("====================")
	//Integers
	s := 8              //2^3
	fmt.Println(s << 3) //2^3 * 2^3 = 2^6
	fmt.Println(s >> 3) //2^3 / 2^3 = 2^0

	fmt.Println("====================")
	//Float
	var f float64 = 3.14
	f = 13.7e72
	f = 2.1e14
	fmt.Printf("%v, %T \n", f, f)
	fmt.Println("====================")
	f1 := 10.2
	f2 := 3.7
	fmt.Println(f1 + f2)
	fmt.Println(f1 - f2)
	fmt.Println(f1 * f2)
	fmt.Println(f1 / f2)

	fmt.Println("====================")
	//Complex
	var c complex64 = 1 + 2i
	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T \n", real(c), real(c))
	fmt.Printf("%v, %T \n", imag(c), imag(c))

	var c2 complex128 = complex(5, 12)
	fmt.Printf("%v, %T \n", c2, c2)

	fmt.Println("====================")
	//Text types
	str := "this is a String"
	fmt.Printf("%v, %T \n", str, str)
	fmt.Printf("%v, %T \n", string(str[2]), str[2])
	str2 := []byte(str)
	fmt.Printf("%v, %T \n", str2, str2)

}

/*
Summary :
Boolean type ->
Values are true or false
Not an alias for other types =
zero value is false
---------------------------
Integers ->
signed int :
	int type has varying size, but min 32 bits
	8 bit (int8) ...64 bit
Unsigned int :
	8 bit (byte and uint8) ..32(unit32)
Arth ->
+
-
*
/
%
operators :
And , or , xor, not
*/
