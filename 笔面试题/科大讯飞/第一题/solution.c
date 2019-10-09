#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(){
    int n;
    scanf("%d", &n);
    if (n>35 || n<1){
        printf("0\n");
        return 0;
    }
    int arr[n];
    arr[0]=3;arr[1]=3;
    for (int i=2;i<n;i++){
        arr[i]=arr[i-1]+arr[i-2];
    }
    long int res=0;
    for(int i=n-1;i>=0;i--){
        res+=arr[i];
        res*=2;
    }
    printf("%ld\n",res);
    return 0;
}