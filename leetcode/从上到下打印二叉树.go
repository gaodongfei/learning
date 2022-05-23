package leetcode

/*
题目一：
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。



例如:
给定二叉树: [3,9,20,null,null,15,7],

   3
/ \
9  20
/  \
15   7
返回：

[3,9,20,15,7]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	a := []int{}

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		node := queue[0]
		if node != nil {
			a = append(a, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[1:]
	}

	return a

}

/*
结论：
1. 使用动态队列，能很好的满足需求
2. 随着节点的遍历，而动态添加子树
3. 过程如下：
	1. 插入根节点
	2. 取出节点
	3. 插入节点的左子树和右子树
*/

/*
题目二：

从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	a := [][]int{
		{},
	}
	queue, tmp := []*TreeNode{root}, []*TreeNode{}
	for len(queue) != 0 {
		node := queue[0]
		if node != nil {
			a[len(a)-1] = append(a[len(a)-1], node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = queue[1:]
		if len(queue) == 0 {
			if len(tmp) != 0 {
				queue = tmp
				tmp = []*TreeNode{}
				a = append(a, []int{})
			}
		}
	}

	return a
}

/*
结论：
1. 是题目一的变形，主要考察怎么判断出每一层
2. 额外使用一个临时切片，保存当层的子节点，当该层遍历完毕后，赋值给上一层
*/

/*
func levelOrder2(root *TreeNode) [][]int {
	a := [][]int{}

	if root == nil {
		return a
	}

	queue := []*TreeNode{root}

	for i:=0;0<len(queue);i++{
		a = append(a,[]int{})
		tmpQueue := []*TreeNode{}
		for j:=0;j<len(queue);j++{
			node := queue[j]
			a[i] = append(a[i],node.Val)
			if queue[j].Left != nil{
				tmpQueue = append(tmpQueue,node.Left)
			}
			if queue[j].Right != nil {
				tmpQueue = append(tmpQueue,node.Right)
			}
		}
		queue = tmpQueue
	}
	return a
}
*/

/*
上面的方式是标准答案的解析

结论：
1. 在取指定行赋值时，它这里使用的是外层的循环标识，而我的解题方式用的是直接取最后一个元素，我的方式有点暴力
2. 它没有动态的去改变切片的大小，而使用的是循环递增和长度判断的方式
*/

/*
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]
*/

func levelOrder3(root *TreeNode) [][]int {
	a := [][]int{}

	if root == nil {
		return a
	}

	queue := []*TreeNode{root}

	for i := 0; 0 < len(queue); i++ {
		a = append(a, []int{})
		tmpQueue := []*TreeNode{}
		queue_n := len(queue)
		for j := 0; j < queue_n; j++ {
			node := queue[j]
			if i%2 != 0 {
				node = queue[queue_n-j-1]
			}
			a[i] = append(a[i], node.Val)
			if queue[j].Left != nil {
				tmpQueue = append(tmpQueue, queue[j].Left)
			}
			if queue[j].Right != nil {
				tmpQueue = append(tmpQueue, queue[j].Right)
			}
		}
		queue = tmpQueue
	}
	return a
}

/*
结论：
1. 随便找一循环内（最外层，中间，最内层），判断该层是奇数层，对其翻转


*/
