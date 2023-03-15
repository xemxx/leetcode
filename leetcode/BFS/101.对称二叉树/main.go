package main
import "fmt"
func main(){
    str := "[[\"[1,2,2,3,4,4,3]\"],[\"[1,2,2,null,3,null,3]\"]]"
    arr := parseStringArrArr(str)
    for i:=0;i<len(arr);i++ {
        unitArgs:=arr[i]
        arg0 := deserializeTreeNode(unitArgs[0]);
        result := isSymmetric(arg0);
        resultabc :=serializeInterface(result);
        fmt.Printf("resultabc%d:%sresultend", i,resultabc)
    }
}