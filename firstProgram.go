package main

import "fmt"

func main()  {
	/* The first hello world */
	fmt.Println("Hello World")

	/* TIPE DATA */
	/* 1. Int */
	// - int8 (-128 ~ 127)
	// - int16 (-32768 ~ 32767)
	// - int32 (-2147483648 ~ 2147483647)
	// - int64 (-9223372036854775808 ~ 9223372036854775807)

	/* 2. Uint -> Always positif */
	// - uint8 (0 ~ 255)
	// - uint16 (0 ~ 65535)
	// - uint32 (0 ~ 4294967295)
	// - uint64 (0 ~ 18446744073709551615)

	/* 3. float */
	// - float32 
	// - float64 
	// - complex64 
	// - complex128 
	
	/* ALIAS */
	// - byte -> uint8 
	// - rune -> int32 
	// - int -> minimal int32 
	// - uint -> minimal uint32 
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)

	/* 4. Boolean */
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	/* 5. String */
	fmt.Println("Hallo nama saya Arief Rachman Hakim")
	fmt.Println(len("Arief Rachman Hakim"))
	fmt.Println("Arief"[0]) //return byte of the char


	/* VARIABLE */
	// - with var and type
	var name string  

	name = "arief"
	fmt.Println(name)
	
	name = "budi" // var bisa diubah"
	fmt.Println(name)

	// - with var but without type
	var myCatName = "si unyin"
	fmt.Println(myCatName)

	// - without var and type
	birdName := "Jhony"
	fmt.Println(birdName)

	// multiple declration
	var (
		firstName = "Arief Rachman"
		lastName = "Hakim"
	)

	fmt.Println( "My name is ",firstName,lastName)
}