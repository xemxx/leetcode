package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m1, m2 sync.Mutex
	n      int
)

func main() {
	n = 2
	m1.Lock()
	go foo()
	go bar()
	time.Sleep(1 * time.Second)
	//mutex.Lock()
}

func foo() {
	for i := 0; i < n; i++ {
		m2.Lock()
		fmt.Print("foo")
		m1.Unlock()
	}
}
func bar() {
	for i := 0; i < n; i++ {
		m1.Lock()
		fmt.Print("bar")
		m2.Unlock()
	}
}

/*
* 这个也可以采用channel的方式来实现顺序控制并且资源消耗会更少
*
* C++代码说明：
* 单纯的加上一个leetcode能过的版本
* 我觉得官方不加上golang的选择可能就是因为golang不是直接操作线程而是有自己的goroutine
#include <mutex>
class FooBar {
private:
    int n;

public:
    std::mutex m1,m2;
    FooBar(int n) {
        this->n = n;
        m1.lock();
    }

    void foo(function<void()> printFoo) {

        for (int i = 0; i < n; i++) {
            m2.lock();
        	// printFoo() outputs "foo". Do not change or remove this line.
        	printFoo();
            m1.unlock();
        }
    }

    void bar(function<void()> printBar) {

        for (int i = 0; i < n; i++) {
            m1.lock();
        	// printBar() outputs "bar". Do not change or remove this line.
        	printBar();
            m2.unlock();
        }
    }
};

*/
