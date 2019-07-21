package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m1, m2, m3 sync.Mutex
	n          int
)

func main() {
	n = 5
	m2.Lock()
	m3.Lock()
	go zero()
	go even()
	go odd()
	time.Sleep(1 * time.Second)
	//mutex.Lock()
}

func zero() {
	for i := 1; i <= n; i++ {
		m1.Lock()
		fmt.Print(0)
		if i%2 == 0 {
			m3.Unlock()
		} else {
			m2.Unlock()
		}
	}
}
func even() {
	for i := 2; i <= n; i += 2 {
		m3.Lock()
		fmt.Print(i)
		m1.Unlock()
	}
}
func odd() {
	for i := 1; i <= n; i += 2 {
		m2.Lock()
		fmt.Print(i)
		m1.Unlock()
	}
}

/*
* 这个也可以采用channel的方式来实现顺序控制并且资源消耗会更少
*
* C++代码说明：
* 单纯的加上一个leetcode能过的版本
* 我觉得官方不加上golang的选择可能就是因为golang不是直接操作线程而是有自己的goroutine
* 有一个点，如果出现错误“You passed invalid value. Exiting.”那是因为你函数没有按照说明输出正确的奇偶数，直接给你提前停了
#include <mutex>
class ZeroEvenOdd {
private:
    int n;

public:
    std::mutex m1,m2,m3;
    ZeroEvenOdd(int n) {
        this->n = n;
        m2.lock();
        m3.lock();
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(function<void(int)> printNumber) {
        for (int i = 1; i <= n; i++) {
            m1.lock();
            printNumber(0);
            if(i%2==0){
                m2.unlock();
            }else{
                m3.unlock();
            }
        }
    }

    void even(function<void(int)> printNumber) {
        for (int i = 2; i <= n; i=i+2) {
            m2.lock();
            printNumber(i);
            m1.unlock();
        }
    }

    void odd(function<void(int)> printNumber) {
        for (int i = 1; i <= n; i=i+2) {
            m3.lock();
            printNumber(i);
            m1.unlock();
        }
    }
};

*/
