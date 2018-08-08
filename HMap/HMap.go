package HMap

import "HashProjectByGo/LNodes"

var buckets = make([]*LNodes.Node, 16)
/**
创建一个长度为16的数组
 */
func InitBuckets() {
	for i := 0; i < 16; i++ {
		buckets[i] = LNodes.CreateHead(LNodes.KValue{"head", "node"})
	}
}

/**
任何长度的字符串，通过运算，散列成0-15整数
1103515245特殊的值
任何值通过此hashCode散列出的0-15的数字的概率是相等的，TODO抽取到utils
 */
func HashCode(key string) int {
	//index := 0
	index := int(key[0])
	for k := 0; k < len(key); k++ {
		index *= (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16 - 1
	return index
}

//向hashmap中保存健值对
func AddKeyValye(key string, value string) {
	//计算出key散列的结果，作为数组的下标
	nIndex := HashCode(key)
	//获取数组中的头节点
	headNode := buckets[nIndex]
	//获取当前链表的尾节点
	tailNode := LNodes.TailNode(headNode)
	//添加节点
	LNodes.AddNode(LNodes.KValue{key, value}, tailNode)
}
/**
获取键值对
 */
func GetValueByKey(key string) string {
	nIndex := HashCode(key)
	headNode := buckets[nIndex]
	//通过链表查询对应key 的value
	return LNodes.FindValueByKey(key, headNode)
}
