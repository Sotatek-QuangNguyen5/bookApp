package models


type Conversation struct {

	Conversation_id int `json:"conversation_id"`
	Name			string `json:"name"`
	IsPrivate		int `json:"isPrivate"`
	Last_update		int `json:"last_update"`
}

type ConversationResponsive struct {

	Conversation_id int `json:"conversation_id"`
	Name			string `json:"name"`
	IsPrivate		int `json:"isPrivate"`
	Last_update		int `json:"last_update"`
	Creator			int `json:"creator"`
	Authors			[]int `json:"authors"`
}


type ConversationAuthor struct {

	Conversation_id int `json:"conversation_id"`
	Author_id		int `json:"author_id"`
	IsCreator		int `json:"iscreator"`
	Last_update		int `json:"last_update"`
}

type ConversationMessage struct {

	Message_id 		int 	`json:"message_id"`
	Conversation_id int 	`json:"conversation_id"`
	Author_id		int 	`json:"author_id"`
	Message			string	`json:"message"`
	File			string	`json:"file"`
	Last_update		int		`json:"last_update"`
}