package wiretutorial

type Message string

func NewMessage(phrase string) Message {
	return Message(phrase)
}
