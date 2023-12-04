package main

import (
	"fmt"
	datatype "golang-dasar/01-data-type"
)


func main(){
    //! Print Hello World
    fmt.Println("Hello World")

    //! Data Type
    fmt.Println(datatype.IntValue) //integer
    fmt.Println(datatype.FloatValue) // float
    fmt.Println(datatype.BooleanValue) // Boolean
    fmt.Println(datatype.StringValue) // String

    //! String Function
    var myLengthString = len(datatype.StringValue)
    fmt.Println("Length of a String :", myLengthString) // Lenght of A String
    var getValueOfString = string(datatype.StringValue[0])
    fmt.Println("Value of a String :", getValueOfString) // Lenght of A String

    // ! Type Declration
    type NoKtp string
    var ktpArief NoKtp = "1111111"
    fmt.Println(ktpArief)


    // ! Boolean
    var nilaiAkhir float64 = 90
    var absensi float64 = 80

    var lulusNilaiAkhir bool =  nilaiAkhir > 80
    var lulusAbsensi bool = absensi > 80

    var lulus bool = lulusNilaiAkhir && lulusAbsensi

    fmt.Println("Apakah Anda lulus? ", lulus)

    // ! Array
    // * make a template
    var names[3] string
    names[0] ="arief"
    names[1] ="budi"
    names[2] ="alex"
    // names[3] ="this make an error" // it will err index of bounds

    // * make an array direct way
    var values = [3]int{
        90,
        80,
        21,
    }
    fmt.Println(values)


}