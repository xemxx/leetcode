package main
import "fmt"
func main(){
    str := "[[\"\\\"xx\\\"\",\"\\\"yy\\\"\"],[\"\\\"xy\\\"\",\"\\\"yx\\\"\"],[\"\\\"xx\\\"\",\"\\\"xy\\\"\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseString(unitArgs[0]);
        arg1 := parseString(unitArgs[1]);
        result := minimumSwap(arg0,arg1);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}