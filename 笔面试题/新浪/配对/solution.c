#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int part(int (* arr)[],int left,int right);

void quickSort(int (* arr)[],int left,int right) {
    if (left<right){
        int p=part(arr,left,right);
        quickSort(arr,left,p-1);
        quickSort(arr,p+1,right);
    }
}

int part(int (* arr)[],int left,int right){
    int temp=(* arr)[left];
    int j=left+1;
	for (int i = j; i <= right; i++) {
		if ((* arr)[i] < temp ){
            int r=(* arr)[i];
            (* arr)[i]=(* arr)[j];
            (* arr)[j]=r;
			j++;
		}
	}
    int r=(* arr)[left];
    (* arr)[left]=(* arr)[j-1];
    (* arr)[j-1]=r;
    return j-1;
}

int main(){
    int n,k;
    scanf("%d %d\n",&n,&k);
    int aaa[n];
    for(int i=0;i<n;i++){
        scanf("%d",&aaa[i]);
    }
    int (* arr)[n]=&aaa;
    quickSort(arr,0,n-1);
    int c=0;
    for(int i=0;i<n-1;i++){
        int num=aaa[i+1]-aaa[i];
        if (num<=k){
            c++;
            i++;
        }

    }
    printf("%d\n",c);
    return 0;
}