package main
import "fmt"
func main(){
    str := "[[\"[8,32,31,18,34,20,21,13,1,27,23,22,11,15,30,4,2]\",\"148\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseIntegerArr(unitArgs[0]);
        arg1 := parseInteger(unitArgs[1]);
        result := minSubarray(arg0,arg1);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}