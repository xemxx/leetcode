#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int max(int i,int j){
    if (j>i){
        i=j;
    }
    return i;
}

int main(){
    int m,n,k;
    scanf("%d %d %d\n",&n,&m,&k);
    int aaa[n];
    for(int i=0;i<n;i++){
        scanf("%d",&aaa[i]);
    }
    int amax=0,imax=aaa[0];
    int bbb[m];
    bbb[0]=aaa[0];
    int c=1;
    for(int i=1;i<n;i++){
        if(c==m){
            for(int j=0;i<m;j++){
                bbb[j]=bbb[j+1];
            }
            bbb[c-1]=0;
            c--;
        }

        bbb[c]=aaa[i];
        c++;
        int tmp=c;
        if (i<n-k){
            for(int j=i;j<i+k;j++){
                imax=max(imax+aaa[j],imax+aaa[j+1]);
            }
        }else{
            imax+=aaa[i];
        }
        amax=max(amax,imax);
    }
    printf("%d\n",amax);
    return 0;
}