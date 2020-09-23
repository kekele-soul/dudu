package strustt

//按时间查询的新笑话。
type To struct {
	Reason string
	Result Data
}

type Data struct {
	Data []Content
}

type Content struct {
	Content string //数据
	HashId string	//哈希值
	Unixtime string //时间戳
	Updatetime string//更新时间
}

