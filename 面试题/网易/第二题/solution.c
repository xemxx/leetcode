
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int findRoot(int arr[][3],int max,int targe){
    for(int i=0;i<max;i++){
        if (*(arr[i] + 1) ==targe || *(arr[i] + 2)==targe){
            
            int tt=findRoot(arr,max,i);
            if (tt==-2){
                return i;
            }
            return tt;
            
        }
    }
    return -2;
}

void dp(int arr[][3],int l,int root,int (* sum)[]){
    (* sum)[l]+=arr[root][0];
    if (*(arr[root]+1)!=-1){
        dp(arr,l+1,*(arr[root]+1),sum);
    }
    if (*(arr[root]+2)!=-1){
        dp(arr,l+1,*(arr[root]+2),sum);
    }
}

int main(){
    int n;
    scanf("%d", &n);
    char* str= (char *)malloc(512000 * sizeof(char));
    for(int i = 0; i < n; i++){
        int m;
        scanf("%d", &m);
        int nnn[m][3];
        for(int j=0;j<m;j++){
            scanf("%d %d %d",&nnn[i][0],&nnn[i][1],&nnn[i][2]);
        }

        int root=findRoot(nnn,m,nnn[0][1]);
        int (* sum)[m];
        dp(nnn,0,root,sum);
        char *flag="YES";
        for (int i=1;i<m;i++){
            if (sum[i]<sum[i-1]){
                flag="NO";
                break;
            }
        }
        sprintf(str, "%s,%d\n%s", str,root, flag);
    }
    printf("%s\n",str);
    return 0;
}