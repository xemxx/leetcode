package main
import "fmt"
func main(){
    str := "[[\"\\\"WBBWWBBWBW\\\"\",\"7\"],[\"\\\"WBWBBBW\\\"\",\"2\"],[\"\\\"WBBWWWWBBWWBBBBWWBBWWBBBWWBBBWWWBWBWW\\\"\",\"15\"],[\"\\\"WWBBBWBBBBBWWBWWWB\\\"\",\"16\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseString(unitArgs[0]);
        arg1 := parseInteger(unitArgs[1]);
        result := minimumRecolors(arg0,arg1);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}