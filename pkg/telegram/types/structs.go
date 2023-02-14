package types

type BotMessage struct {
	Message struct {
		Message_ID int
		From       struct {
			Username string
			Id       int
		}
		Chat struct {
			Id int
		}
		Text string
	}
}
