package toy3

// 双向链表
type element struct {
	next, prev *element
	list       *list
	value      interface{}
}

type list struct {
	root element // 根元素
	len  int     // 链表长度
}
