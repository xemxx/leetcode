首先这个是一个topN问题

思路如下：
- 最简单的就是先排序然后直接取
- 然后可以采用部分排序，排出前1000个数即可，具体实现就是优化简单的排序算法，冒泡，插入等，这里不作解析
- 再优化可以采用分治法，类似于快排中的partition操作，每次得到归位的povit值，然后判断前一部分总数是否大于K，如果大于K，即K所代表的值在前一部分中，只需要继续partition前一部分即可，如果小于K，即代表需要的值在后一部分中，只需要继续partition后一部分即可，本题采用此解法，时间复杂度<2n,渐进为O(n)
- 然后其实此题可能存在内存限制，可以采用最小堆，先取出前K个数组成一个大小为K的最小堆，然后读取后面的数，每次读取到的值如果小于堆顶元素则抛弃，如果大于堆顶元素则置换堆顶元素，并调整最小堆，最终，堆顶元素就是最终值