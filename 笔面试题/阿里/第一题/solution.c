#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

/** 请完成下面这个函数，实现题目要求的功能 **/
 /** 当然，你也可以不按照这个模板来作答，完全按照自己的想法来 ^-^  **/
int getRelationCount(int takeBoatList_size, char** takeBoatList, double threshold) {
    char* student[takeBoatList_size][2];
    int one,two=0;
    for(int i=0;i<takeBoatList_size;i++){
        char* s=strtok(takeBoatList[i],",");
        student[i][0]=s;
        s=strtok(takeBoatList[i],",");
        student[i][1]=s;
        if (two==0 && s!=student[one][0]){
            two=i;
        }else if (s!=student[one][0] && s!=student[two][0]){
            int tmp=i;
            for 
        }
    }
    return 1;
}

int main() {
    int res;

    int _takeBoatList_size = 0;
    int _takeBoatList_i;
    scanf("%d\n", &_takeBoatList_size);
    char* _takeBoatList[_takeBoatList_size];
    for(_takeBoatList_i = 0; _takeBoatList_i < _takeBoatList_size; _takeBoatList_i++) {
        char* _takeBoatList_item;
        _takeBoatList_item = (char *)malloc(512000 * sizeof(char));
        scanf("\n%[^\n]", _takeBoatList_item);
        
        _takeBoatList[_takeBoatList_i] = _takeBoatList_item;
    }
    double _threshold;
    scanf("%lf", &_threshold);
    
    res = getRelationCount(_takeBoatList_size, _takeBoatList, _threshold);
    printf("%d\n", res);
    
    return 0;

}