dll_caller
==========

A windows dll call hellper

### Windows MessageBox Example
```go
package main

import (
    "github.com/igonow/dll_caller"
    "fmt"
)

func main(){
    ShowMessageBox()
}

func ShowMessageBox() {
    var dll *dll_caller.Dll
    if d, e := dll_caller.NewDll("user32.dll"); e != nil {
        fmt.Println(e.Error())
        return
    } else {
        dll = d
    }

    if e := dll.InitalFunctions("MessageBoxW"); e != nil {
        fmt.Println(e.Error())
        return
    }

    ret, err := dll.Call("MessageBoxW", 0, "hello", "world", 3)

    fmt.Println(ret, err)
}
```