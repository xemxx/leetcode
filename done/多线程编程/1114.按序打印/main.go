package main

import (
	"fmt"
	"sync"
	"time"
)

var m1, m2 sync.Mutex

func main() {
	m1.Lock()
	m2.Lock()
	go first()
	go second()
	go third()
	time.Sleep(1 * time.Second)
	//mutex.Lock()
}

func first() {
	fmt.Print("one")
	m1.Unlock()
}
func second() {
	m1.Lock()
	fmt.Print("two")
	m1.Unlock()
	m2.Unlock()
}
func third() {
	m2.Lock()
	fmt.Print("three")
	m2.Unlock()
}

/*
* 这个也可以采用channel的方式来实现顺序控制并且资源消耗会更少
*
* C++代码说明：
* 单纯的加上一个leetcode能过的版本
* 我觉得官方不加上golang的选择可能就是因为golang不是直接操作线程而是有自己的goroutine
#include <mutex>

class Foo {
public:
    std::mutex m1,m2;
    Foo() {
        m1.lock();
        m2.lock();
    }

    void first(function<void()> printFirst) {
        printFirst();
        m1.unlock();
    }

    void second(function<void()> printSecond) {
        m1.lock();
        printSecond();
        m1.unlock();
        m2.unlock();
    }

    void third(function<void()> printThird) {
        m2.lock();
        printThird();
        m2.unlock();
    }
};

*/
