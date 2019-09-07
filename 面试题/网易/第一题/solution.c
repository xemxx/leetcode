
#include <stdio.h>

int is_hui(int arr[],int j){
    int i=0;
    while(i<=j){
        if(arr[i]!=arr[j]){
            return 0;
        }
        i++;j--;
    }
    return 1;
}

int main(){
    int n;
    scanf("%d", &n);
    int arr[n];
    for(int i = 0; i < n; i++){
        scanf("%d", &arr[i]);
    }
    for(int i = 0; i < n; i++){
        int m=0,tmp=arr[i];
        while(tmp>0){
            tmp=tmp>>1;
            m++;
        }

        int temp[m];
        tmp=arr[i];
        for (int j=0;j<m;j++){
            temp[j]=tmp%2;
            tmp=tmp/2;
        }

        if (is_hui(temp,m-1)){
            printf("YES\n");
        }else{
            printf("NO\n");
        }
    }  
    return 0;
}