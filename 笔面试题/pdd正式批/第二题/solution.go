package main

import "fmt"

func main() {
	var s string 
	fmt.Scan(&s)
	
	fmt.Println(isStr(s))
}

func isStr(s string) int{
	f:=make([]int,len(s))
	max:=0
	for i:=1;i<len(s);i++{
		if f[i-1]==0{
			if s[i-1]=='('{
				if s[i]==')'{
					f[i]=2
				}
			}else{
				f[i]=0
			}
		}else{
			if s[i]==')'{
				tmp:=i-(f[i-1])-1
				if tmp>0{
					if s[tmp]=='('{
						f[i]=f[i-1]+2
					}
					tmp--
					for tmp>=0{
						if f[tmp]>0{
							f[i]+=f[tmp]
							tmp=tmp-f[tmp]
						}else{
							break
						}
					}
				}
			}else{
				f[i]=0
			}
		}
		if f[i]>max{
			max=f[i]
		}
	}
	return max
}