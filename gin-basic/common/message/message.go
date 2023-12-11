package message

func Success() {

}

func Fail(errorInfo string) {

}

func ErrorServer() *Message {
	return &Message{
		500,
		"服务器异常",
		nil,
		0,
	}
}
