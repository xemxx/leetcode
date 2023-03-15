package main
import "fmt"
func main(){
    str := "[[\"\\\"{a,b}{c,{d,e}}\\\"\"],[\"\\\"{{a,z},a{b,c},{ab,z}}\\\"\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseString(unitArgs[0]);
        result := braceExpansionII(arg0);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}