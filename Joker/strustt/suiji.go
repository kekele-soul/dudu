package strustt

//随机笑话
type Random struct {
	Content string //数据
	HashId string //哈希值
	Unixtime string //时间戳
}
type Randome struct {
	Reason string
	Result []Random
}