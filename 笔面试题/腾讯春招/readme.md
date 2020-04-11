## 题目

时间：1小时

1. 判断数上是否存在一条从根节点到叶子节点的路径，使得路径上所有节点的和为给定值
struct TreeNode
{
    TreeNode* left;
    TreeNode* right;
    int val;
};
bool IsTreeMeetTarget(TreeNode* root, int target);

2 合并两个有序链表
struct ListNode
{
    ListNode* next;
    int    val;
};
ListNode* MergeTwoSortedList(ListNode* l1, ListNode* l2);


3. 给定一个数组，找出数组中所有3个数和为0的序列，序列不可重复。
例如：输入[1, 2, -3, -2, 0]，输出[[1, 2, -3], [2,  -2, 0]]
std::vector<std::vector<int> > threeSum(std::vector<int>& nums);

4. 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。返回一个表示每个字符串片段
输入: S = "ababcbacadefegdehijhklij"
输出: ["ababcbaca", "defegde", "hijhklij"]
vector<string> partitionLabels(string S)