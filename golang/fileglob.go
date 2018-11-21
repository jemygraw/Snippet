package main

import(
    "fmt"
    "path/filepath"
)

func main(){
    matches,err:=filepath.Glob("*.go")
    if err!=nil{
        fmt.Println(err)
        return 
    }

    for _,name:=range matches{
        fmt.Println(name)
    }
}
