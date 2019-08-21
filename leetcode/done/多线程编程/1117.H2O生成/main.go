package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mo, mh sync.Mutex
	n      int
)

func main() {
	n = 0
	go func() {
		for {
			oxygen()
		}
	}()
	go func() {
		for {
			hydrogen()
		}
	}()
	time.Sleep(1 * time.Second)
}

func oxygen() {
	mo.Lock()
	fmt.Print("O")
	n = 0
	mh.Unlock()
}
func hydrogen() {
	mh.Lock()
	fmt.Print("H")
	n++
	if n >= 2 {
		mo.Unlock()
	} else {
		mh.Unlock()
	}
}

/*
* 实现说明：
* 这个也可以采用channel的方式来实现顺序控制并且资源消耗会更少
* 此外题目本身的意思是输出一个h后可以输出h或者o，但是此处采用类似于信号量的方式，如需要可以采用多个信号量实现题目原意
* golang版本说明：
* 暂时没有实现测试的功能，所以无法运行，但是如果leetcode支持了golang应该是可以通过的
*
* C++代码说明：
* 单纯的加上一个leetcode能过的版本
* 我觉得官方不加上golang的选择可能就是因为golang不是直接操作线程而是有自己的goroutine

#include <mutex>
class H2O {
public:
    std::mutex mo,mh;
    int n;
    H2O() {
        this->n=0;
    }

    void hydrogen(function<void()> releaseHydrogen) {

        // releaseHydrogen() outputs "H". Do not change or remove this line.
        mh.lock();
        releaseHydrogen();
        n++;
        if(n==2){
            mo.unlock();
        }else{
            mh.unlock();
        }
    }

    void oxygen(function<void()> releaseOxygen) {

        // releaseOxygen() outputs "O". Do not change or remove this line.
        mo.lock();
        releaseOxygen();
        n=0;
        mh.unlock();
    }
};

*/
