#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(){
    char *str=(char *)malloc(512000 * sizeof(char));
    scanf("\n%[^\n]", str);
    char first=str[0];
    char last='A';
    int ok=1;
    int res=0;
    int lenn=0;
    //printf("%ld\n",strlen(str));
    for(int i=1;i<strlen(str);i++){
       // printf("%c,%c,%c,%d\n",str[i],first,last,res);
        if (str[i]!=';'){
            if (ok){
                if ((str[i]<='Z' && str[i]>='A') ||(str[i]>='a' && str[i]<='z')){
                    //ok=1;
                    lenn++;
                }else{
                    ok=0;
                }
            }
            last=str[i];
        }else{
            if (ok==1 && lenn!=0){
                if (last>first){
                    res+=last-first;
                }else{
                    res+=first-last;
                }
            }
            ok=1;
            lenn=0;
            first=str[i+1];
        }
    }
    //adsdas12312;a;b1ac;;bc;Zffffz;;
    last=str[strlen(str)-1];
    if (ok==1 && lenn!=0){
        if (last>first){
            res+=last-first;
        }else{
            res+=first-last;
        }
    }
    printf("%d\n",res);
    return 0;
}
