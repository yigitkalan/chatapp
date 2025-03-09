package message

func GetAllMessages() []Message {
	result, err := GetRepo().GetAll()

	if err != nil {
		panic("Unexpected Error")
	}
	return result

}

func GetMessageById(id uint) Message {
	result, err := GetRepo().GetById(id)

	if err != nil {
		panic("Unexpected Error")
	}
	return result

}

func CreateMessage(message Message) {
	err := GetRepo().Add(message)

	if err != nil {
		panic("Unexpected Error")
	}
}
