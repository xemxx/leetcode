## 前言：
本题目其实很简单，4个方向依次判断，可以很快的有思路
但是我之前用c写2048没成功过，，，于是准备肝这个题目，然后我强迫症又犯了，觉得分别写4个方向很啰嗦麻烦，然后就考虑用一个通式解决，然后又发现上下和左右方向时`i，j`的位置需要互换甚至循环的先后要变换，所以还是分别写了上下和左右，其实如果把数组颠倒一下就不需要了，但是考虑我不会用go写这个，同时已经快没有时间，只好肝了


## 解法思路：

说向上方向的思路，多个方向则通用：
- 两层遍历，从`[0][0]`开始，第一层遍历` i := e; 0 <= i && i <= 3; i += q`，第二层遍历`j,tail=1,4;tail>0;tail--`，同时用`x=0`表示上一个`j`的值,`tail`表明咱们只遍历四个元素，因为`j`可能出现`j--`的情况，具体往下看
- 每次遍历，应该有三种情况（也可以算做两种）：
  - 判断上一次的值`map[i][x]`是否为0，如果为0，则下标为`x`后面的元素全部前移，并且末尾添0，同时`x,j`均不需要改变
  - 判断本次的值`map[i][j]`是否为0，如果为0，则下标为`j`后面的元素全部前移（注意`map[i][x]`这个时候不为0，所以不变），并且末尾添0，同时`x,j`均不需要改变
  - PS：在上述的两种情况中，这时第二层循环的边界如果为`j<=3`的话，就会出现遍历为了补位而添加的0，所以需要改变开始的边界条件
  - 最后一种情况就是`map[i][x] == map[i][j]`并且均不为0，则相加，并且下标为`j`后面的元素全部前移，并且末尾添0，因为相同的元素只需要相加一次，所以`x++,j++`，此外如果不相等也需要`x++,j++`，因为不相等则代表`x`不需要再考虑，但是`y`及以后的元素需要再此判断

此外就是我的做法了，通过一些通用的做法：
- 外层循环下标`i`
- 内层循环下标`j`
- `i`每次变化步数`q`
- `j`每次变化步数`w`
- `j`的上一个值`x`
- `i,j`的初始值`e,r`
- 当元素前移时的边界`pid`，因为`j`在这种方法下，可能递增也可能递减，然而`pid`必然表示其变化方向上的一个边界值

这些做法在每个方向上都是一样的，唯一不一样的就是每次修改的起点，和修改的方向不同，主要的修改方向分为上下和左右，这里因为涉及到循环的先后变换，然后笔试的时间关系，我还是采用了2份代码，，，其实思路是有的，将数组行列颠倒，在输出的时候再颠倒回来，这样就可以做到只写一个双循环，并且在外部保证的情况下不容易出错。