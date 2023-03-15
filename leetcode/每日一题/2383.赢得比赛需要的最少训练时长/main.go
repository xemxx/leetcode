package main
import "fmt"
func main(){
    str := "[[\"5\",\"3\",\"[1,4,3,2]\",\"[2,6,3,1]\"],[\"2\",\"4\",\"[1]\",\"[3]\"],[\"1\",\"1\",\"[1,1,1,1]\",\"[1,1,1,50]\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseInteger(unitArgs[0]);
        arg1 := parseInteger(unitArgs[1]);
        arg2 := parseIntegerArr(unitArgs[2]);
        arg3 := parseIntegerArr(unitArgs[3]);
        result := minNumberOfHours(arg0,arg1,arg2,arg3);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}