package main
import "fmt"
func main(){
    str := "[[\"4\",\"[[0,1],[0,3],[1,2],[1,3]]\"],[\"5\",\"[[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]\"],[\"8\",\"[[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]\"],[\"2\",\"[[1,0]]\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := parseInteger(unitArgs[0]);
        arg1 := parseIntegerArrArr(unitArgs[1]);
        result := maximalNetworkRank(arg0,arg1);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}