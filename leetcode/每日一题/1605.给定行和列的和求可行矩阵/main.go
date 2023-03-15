package main
import "fmt"
func main(){
    str := "[[\"[3,8]\",\"[4,7]\"],[\"[5,7,10]\",\"[8,6,8]\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseIntegerArr(unitArgs[0]);
        arg1 := parseIntegerArr(unitArgs[1]);
        result := restoreMatrix(arg0,arg1);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}