package main

import "fmt"
import "reflect"
import "runtime"

func bacatipe1(){
    var umur = 26
    var reflectValue = reflect.ValueOf(umur)
    
    fmt.Println("Tipe Variable :", reflectValue.Type())
    
    if reflectValue.Kind() == reflect.Int {
        fmt.Println("Nilai Variable :", reflectValue.Int())
        fmt.Println("")
    }
}

func bacatipe2(){
    var nama = "Hasan Suryaman"
    var reflectValue = reflect.ValueOf(nama)
    
    fmt.Println("Tipe Variabel :", reflectValue.Type())
    
    if reflectValue.Kind() == reflect.String {
        fmt.Println("Nilai Variebal :", reflectValue.String())
        fmt.Println("")
    }
}

func main(){
    runtime.GOMAXPROCS(2)
    
    go bacatipe1()
    bacatipe2()
    
    var input string
    fmt.Scanln(&input)
}
