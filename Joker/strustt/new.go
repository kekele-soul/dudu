package strustt


//最新笑话
type Tos struct {
	Reason string
	Result Datas
}

type Datas struct {
	Data []Joker
}

type Joker struct {
	Content string //数据
	HashId string	//哈希值
	Unixtime string //时间戳
	Updatetime string//更新时间
}
