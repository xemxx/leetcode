## 题目
跳柱子，每个柱子高度为h，从第一个柱子开始跳，每次只能跳到高度不大于自己的柱子，并且每次跳跃最大距离为k，有一次超能力能无视高度跳跃，求能否跳到最后一个柱子

## 思路
其实就是bfs，只是开始我理解错了题意，认为跳跃的范围k里面所有的柱子都要比自己低，就用了简单的贪心- -