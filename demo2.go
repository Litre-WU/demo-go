package main

import (
    "fmt"
    "github.com/imroc/req"
)

func main() {
    url:="http://127.0.0.1"
    header := req.Header{
        "Accept": "application/json",
        "User-Agent": "Chrome",
    }
    param := req.Param{
        "name": "test",
        "pwd":  "test",
    }
    re,err:= req.Get(url,header,param)
    if err!=nil{
    	fmt.Println(err)
    }
    if re.Response().StatusCode == 200 {
        fmt.Println(re.String())
    }
}
