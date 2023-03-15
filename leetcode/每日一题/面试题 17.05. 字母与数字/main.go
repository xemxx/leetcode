package main
import "fmt"
func main(){
    str := "[[\"[\\\"A\\\",\\\"1\\\",\\\"B\\\",\\\"C\\\",\\\"D\\\",\\\"2\\\",\\\"3\\\",\\\"4\\\",\\\"E\\\",\\\"5\\\",\\\"F\\\",\\\"G\\\",\\\"6\\\",\\\"7\\\",\\\"H\\\",\\\"I\\\",\\\"J\\\",\\\"K\\\",\\\"L\\\",\\\"M\\\"]\"],[\"[\\\"A\\\",\\\"A\\\"]\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseStringArr(unitArgs[0]);
        result := findLongestSubarray(arg0);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}