#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(){
    int n;
    scanf("%d", &n);
    int arr[n][30];
    int l[n][2];
    for(int i = 0; i < n; i++){
        scanf("%d %d", &l[i][0],&l[i][1]);
        for (int j=0;j<l[i][1];j++){
            scanf("%d",&arr[i][j]);
        }
    }
    for(int i = 0; i < n; i++){
        if (l[i][0]==0){
            printf("30\n");
            continue;
        }
        int num=0;
        int last=0;
        int k=l[i][0];
        for (int j=0;j<l[i][1];j++){
            num+=1;
            int tmp=arr[i][j];
            tmp-=last+1;
            while (tmp>k){
                num+=1;
                tmp-=k;
            }
            last=arr[i][j];
        }
        
        if (30-last>k){
            int tmp=30-last-1;
            while (tmp>k){
            num+=1;
            tmp-=k;
        }
        }
       
        printf("%d\n",num);
    }
}