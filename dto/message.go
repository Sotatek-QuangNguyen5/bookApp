package dto

type Message struct {
	Message string `json:"message"`
}

func MessageAddSuccess(obj string) *Message {

	return &Message{

		Message: "Add " + obj + " Success!!!",
	}
}

func MessageCreateSuccess(obj string) *Message {

	return &Message{

		Message: "Create " + obj + " Success!!!",
	}
}

func MessageDeleteSuccess(obj string) *Message {

	return &Message{

		Message: "Delete " + obj + " Success!!!",
	}
}

func MessageUpdateSuccess(obj string) *Message {

	return &Message{

		Message: "Update " + obj + " Success!!!",
	}
}