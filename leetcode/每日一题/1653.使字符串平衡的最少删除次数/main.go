package main
import "fmt"
func main(){
    str := "[[\"\\\"aababbab\\\"\"],[\"\\\"bbaaaaabb\\\"\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseString(unitArgs[0]);
        result := minimumDeletions(arg0);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}