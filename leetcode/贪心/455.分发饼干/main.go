package main

import "fmt"
func main(){
    str := "[[\"[1,2,3]\",\"[1,1]\"],[\"[1,2]\",\"[1,2,3]\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseIntegerArr(unitArgs[0]);
        arg1 := parseIntegerArr(unitArgs[1]);
        result := findContentChildren(arg0,arg1);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}