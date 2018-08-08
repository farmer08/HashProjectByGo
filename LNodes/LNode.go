package LNodes

import "fmt"

type KValue struct {
	Key   string
	Value string
}

type Node struct {
	Data     KValue
	NextNode *Node
}

/**
创建头结点
 */
func CreateHead(data KValue) *Node {
	var head = &Node{data, nil}
	return head
}

/**
添加节点
 */
func AddNode(data KValue, node *Node) *Node {
	var newNode = &Node{data, nil}
	node.NextNode = newNode
	return newNode
}

/**
节点的遍历
 */
func ShowNodes(head *Node) {
	node := head
	for {
		if node.NextNode != nil {
			fmt.Println(node.Data)
			node = node.NextNode
		} else {
			break
		}
	}
	fmt.Println(node.Data)

}

//获得当前链表的尾节点
func TailNode(head *Node) *Node {
	node := head
	for {
		if node.NextNode == nil {
			return node
		} else {
			node = node.NextNode
		}
	}
}
/**
根据key获取值
 */
func FindValueByKey(key string, head *Node) string {
	node := head
	for {
		if node.NextNode != nil {
			if node.Data.Key == key {
				return node.Data.Value
			}
			node = node.NextNode
		} else {
			break
		}
	}
	return node.Data.Value
}
