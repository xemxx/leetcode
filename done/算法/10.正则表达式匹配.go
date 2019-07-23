/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

package main

import (
	"fmt"
)

func main() {
   // s := "ab"
    p := ".*c"
	fmt.Println(p[1]=="*")

}

//回溯法 存在s切片越界情况
func isMatch(s string, p string) bool {
    if len(p)==0 {
        return len(s)==0
    }

    h:=len(s)==0 && (s[0:1]==p[0:1] || p[0:1]==".")

    if len(p)>=2 && p[1:2]=="*"{
        return isMatch(s,p[2:]) || (h && isMatch(s[1:],p))
    }

    return h && isMatch(s[1:],p[1:])
}
//回溯法 解决越界情况
func isMatch_solve(s string, p string) bool {
    if len(p)==0 {
        return len(s)==0
    }
    h:=true
    if len(s)==0 {
        h=false
    }else{
        i,j:=s[0:1],p[0:1]
        h=(i==j || j==".")
    }

    if len(p)>=2 && p[1:2]=="*"{
        if h{
            return isMatch(s,p[2:]) || isMatch(s[1:],p)
        }else{
            return isMatch(s,p[2:]) 
        }
    }

    return h && isMatch(s[1:],p[1:])
}

//做备忘录，官方题解自顶往下，但是时间更久，可能是因为我做了太多if
func isMatch_dp(s string,p string) bool{
    memo =make(map[int]map[int]bool,len(s)+1)
    return dp(0,0,s,p);
}

var memo map[int]map[int]bool

func dp(i int ,j int,s string,p string) bool{
    _, ok := memo[i][j]
	if ok {
		return memo[i][j]
	}
    if len(p)== j {
        return len(s)==i
    }
    h:=true
    if len(s)<=i {
        h=false
    }else{
        h= s[i:i+1]==p[j:j+1] || p[j:j+1]=="."
    }

    if len(p)-2>=j && p[j+1:j+2]=="*"{
        if h{
            return dp(i,j+2,s,p) || dp(i+1,j,s,p)
        }else{
            return dp(i,j+2,s,p) 
        }
    }

    if memo[i]==nil {
        memo[i]=make(map[int]bool,len(p)+1)
    }

    memo[i][j]=h && dp(i+1,j+1,s,p)
    return memo[i][j]
    
}

//找到leetcode用时最少代码,自低往上
func isMatch_dp1(s string, p string) bool {
	lens := len(s)
	lenp := len(p)

	match := make([][]bool, lens+1)
	for i := 0; i < lens+1; i++ {
		match[i] = make([]bool, lenp+1)
	}
	match[0][0] = true

	for i := 0; i <= lens; i++ {
		for j := 1; j <= lenp; j++ {
			if j > 1 && p[j-1] == '*' {
				match[i][j] = match[i][j-2] || (i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && match[i-1][j])
			} else {
				match[i][j] = i > 0 && match[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}

	return match[lens][lenp]
}

