package lab2

import "fmt"

type Node struct {
	value interface{} //值
	next  *Node       //下一节点
}

type LinkTable struct {
	sumOfNode int //节点个数
	pHead     *Node
	pTail     *Node
}

//新建节点的构造函数
func NewNode(v interface{}, nextNode *Node) *Node {
	return &Node{
		value: v,
		next:  nextNode,
	}
}

//创建链表
func createLinkTable() *LinkTable {
	var pLinkTable *LinkTable = new(LinkTable)
	if pLinkTable == nil {
		fmt.Println("CreateLinkTable error.")
		return nil
	}
	pLinkTable.sumOfNode = 0
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	return pLinkTable
}

//在链表尾部添加节点
func (pLinkTable *LinkTable) addLinkTableNode(value interface{}) bool {
	if pLinkTable == nil || value == nil {
		fmt.Println("The input has something wrong.")
		return false
	}

	pNode := NewNode(value, nil)

	if pLinkTable.sumOfNode == 0 {
		//链表为空
		pLinkTable.pHead = pNode
		pLinkTable.pTail = pNode
		pLinkTable.sumOfNode = 1
	} else {
		//链表不为空
		pLinkTable.pTail.next = pNode
		pLinkTable.pTail = pNode
		pLinkTable.sumOfNode++
	}

	return true
}

//获取链表长度
func (pLinkTable *LinkTable) Length() int {
	return pLinkTable.sumOfNode
}

//获取头节点
func (pLinkTable *LinkTable) getLinkTableHead() *Node {
	if pLinkTable == nil {
		fmt.Println("The linktable is empty")
		return nil
	}
	return pLinkTable.pHead
}

//获取尾节点
func (pLinkTable *LinkTable) getLinkTableTail() *Node {
	if pLinkTable == nil {
		fmt.Println("The linktable is empty")
		return nil
	}
	return pLinkTable.pTail
}

//链表按索引位置查找
func (pLinkTable *LinkTable) getLinkTableNodeByIndex(index int) *Node {
	if pLinkTable == nil {
		fmt.Println("The linktable is empty")
		return nil
	}

	if index < 0 || index > pLinkTable.sumOfNode {
		return nil
	}

	tempNode := pLinkTable.pHead
	for j := 0; j < index; j++ {
		tempNode = tempNode.next
	}
	return tempNode

}

//链表按值查找
func (pLinkTable *LinkTable) getLinkTableNodeByValue(v interface{}) *Node {
	if pLinkTable == nil {
		fmt.Println("The linktable is empty")
		return nil
	}
	tempNode := pLinkTable.pHead
	for tempNode != nil {
		if tempNode.value == v {
			return tempNode
		}
		tempNode = tempNode.next
	}
	return nil
}

//链表循环输出
func (pLinkTable *LinkTable) showLinkTable() bool {
	if pLinkTable == nil {
		fmt.Println("The linktable is empty")
		return false
	}
	tempNode := pLinkTable.pHead
	for tempNode != nil {
		fmt.Print(tempNode.value)
		fmt.Print(" ")
		tempNode = tempNode.next
	}
	fmt.Println()
	return true
}

//删除指定节点
func (pLinkTable *LinkTable) delLinkTableNode(pNode *Node) bool {
	if pLinkTable == nil || pNode == nil {
		fmt.Println("The input has something wrong.")
		return false
	}

	if pLinkTable.pHead == pNode {
		pLinkTable.pHead = pLinkTable.pHead.next
		pLinkTable.sumOfNode--
		if pLinkTable.sumOfNode == 0 {
			pLinkTable.pTail = nil
		}
		return true
	}

	tempNode := pLinkTable.pHead
	for tempNode != nil {
		if tempNode.next == pNode {
			pLinkTable.sumOfNode--
			//链表长度为0，或删除尾结点时，尾结点为空
			if pLinkTable.sumOfNode == 0 || pNode == pLinkTable.pTail {
				pLinkTable.pTail = nil
			}
			tempNode.next = tempNode.next.next

			return true
		}
		tempNode = tempNode.next
	}
	fmt.Println("Delete is failed.")
	return false
}

//删除整个链表
func (pLinkTable *LinkTable) deleteLinkTable() bool {
	if pLinkTable == nil {
		fmt.Println("DeleteLinkTable is failed.")
		return false
	}
	for pLinkTable.sumOfNode != 0 {
		tempNode := pLinkTable.pHead
		pLinkTable.pHead = pLinkTable.pHead.next
		tempNode.next = nil
		pLinkTable.sumOfNode--
	}
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	return true
}

func main() {
	LinkTable := createLinkTable()
	LinkTable.addLinkTableNode(1)
	LinkTable.addLinkTableNode(2)
	LinkTable.addLinkTableNode(3)
	LinkTable.addLinkTableNode(4)
	LinkTable.addLinkTableNode(5)
	LinkTable.addLinkTableNode(6)
	fmt.Println("打印现有链表：")
	LinkTable.showLinkTable()
	fmt.Println("删除值为1的节点")
	tempNode := LinkTable.getLinkTableNodeByValue(1)
	LinkTable.delLinkTableNode(tempNode)
	fmt.Println("打印链表：")
	LinkTable.showLinkTable()
	fmt.Println("删除index为2节点")
	tempNode = LinkTable.getLinkTableNodeByIndex(2)
	LinkTable.delLinkTableNode(tempNode)
	fmt.Println("打印现有链表：")
	LinkTable.showLinkTable()
	fmt.Println("输出链表长度：")
	fmt.Println(LinkTable.Length())
	fmt.Println("删除链表尾节点")
	tempNode = LinkTable.getLinkTableTail()
	LinkTable.delLinkTableNode(tempNode)
	fmt.Println("打印链表：")
	LinkTable.showLinkTable()
	fmt.Println("删除整个链表")
	flag := LinkTable.deleteLinkTable()
	if flag == true {
		fmt.Println("删除成功")
	}
}
