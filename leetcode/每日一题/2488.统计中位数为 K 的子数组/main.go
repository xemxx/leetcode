package main
import "fmt"
func main(){
    str := "[[\"[3,2,1,4,5]\",\"4\"],[\"[2,3,1]\",\"3\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseIntegerArr(unitArgs[0]);
        arg1 := parseInteger(unitArgs[1]);
        result := countSubarrays(arg0,arg1);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}