package dto 

import (

	"bookApp/models"
)

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

type RoomConversation struct {

	Conversation_id int `json:"conversation_id"`
	Name			string `json:"name"`
	Author_id		int `json:"author_id"`
	IsPrivate		int `json:"isPrivate"`
	Last_update		int `json:"last_update"`
}

type PrivateConversation struct {

	Author1_id		int `json:"author1_id"`
	Author2_id		int `json:"author2_id"`
	Last_update		int `json:"last_update"`
}

type ConversationAuthor struct {

	Conversation_id int `json:"conversation_id"`
	Author_id		int `json:"author_id"`
	IsCreator		int `json:"iscreator"`
	Last_update		int `json:"last_update"`
}

type ConversationAuthors struct {

	Conversation_id int 					`json:"conversation_id"`
	Author_ids		[]*ConversationAuthor 	`json:"author_ids"`
}

type ConversationMessage struct {

	Message_id 		int 	`json:"message_id"`
	Conversation_id int 	`json:"conversation_id"`
	Author_id		int 	`json:"author_id"`
	Message			string	`json:"message"`
	File			string	`json:"file"`
	Last_update		int		`json:"last_update"`
}

func ConversationModelToConversationDto(ConversationModel *models.Conversation) *Conversation {

	if ConversationModel == nil {

		return nil
	}
	return &Conversation{

		Conversation_id: ConversationModel.Conversation_id,
		Name: ConversationModel.Name,
		IsPrivate: ConversationModel.IsPrivate,
		Last_update: ConversationModel.Last_update,
	}
}

func ConversationsModelToConversationsDto(ConversationsModel []*models.Conversation) []*Conversation {

	var ConversationsDto []*Conversation
	for _, Conversation := range ConversationsModel {

		ConversationsDto = append(ConversationsDto, ConversationModelToConversationDto(Conversation))
	}
	return ConversationsDto
}

func ConversationAuthorModelToConversationAuthorDto(ConversationAuthorModel *models.ConversationAuthor) *ConversationAuthor {

	if ConversationAuthorModel == nil {

		return nil
	}
	return &ConversationAuthor{

		Conversation_id: ConversationAuthorModel.Conversation_id,
		Author_id: ConversationAuthorModel.Author_id,
		Last_update: ConversationAuthorModel.Last_update,
	}
}
func ConversationAuthorsModelToConversationAuthorsDto(ConversationAuthorsModel []*models.ConversationAuthor) []*ConversationAuthor {

	var ConversationAuthors []*ConversationAuthor
	for _, Conversation := range ConversationAuthorsModel {

		ConversationAuthors = append(ConversationAuthors, ConversationAuthorModelToConversationAuthorDto(Conversation))
	}
	return ConversationAuthors
}


func ConversationDtoToConversationModel(ConversationDto *Conversation) *models.Conversation {

	if ConversationDto == nil {

		return nil
	}
	return &models.Conversation{

		Conversation_id: ConversationDto.Conversation_id,
		Name: ConversationDto.Name,
		IsPrivate: ConversationDto.IsPrivate,
		Last_update: ConversationDto.Last_update,
	}
}

func ConversationsDtoToConversationsModel(ConversationsDto []*Conversation) []*models.Conversation {

	var ConversationsModel []*models.Conversation
	for _, Conversation := range ConversationsDto {

		ConversationsModel = append(ConversationsModel, ConversationDtoToConversationModel(Conversation))
	}
	return ConversationsModel
}

func ConversationAuthorDtoToConversationAuthorModel(ConversationAuthorDto *ConversationAuthor) *models.ConversationAuthor {

	if ConversationAuthorDto == nil {

		return nil
	}
	return &models.ConversationAuthor{

		Conversation_id: ConversationAuthorDto.Conversation_id,
		Author_id: ConversationAuthorDto.Author_id,
		Last_update: ConversationAuthorDto.Last_update,
	}
}

func ConversationAuthorsDtoToConversationAuthorsModel(ConversationAuthorsDto []*ConversationAuthor) []*models.ConversationAuthor {

	var ConversationAuthors []*models.ConversationAuthor
	for _, Conversation := range ConversationAuthorsDto {

		ConversationAuthors = append(ConversationAuthors, ConversationAuthorDtoToConversationAuthorModel(Conversation))
	}
	return ConversationAuthors
}

func ConversationMessageDtoToConversationMessageModel(ConversationMess *ConversationMessage) *models.ConversationMessage {

	if ConversationMess == nil {

		return nil
	}
	return &models.ConversationMessage{

		Message_id: ConversationMess.Message_id,
		Conversation_id: ConversationMess.Conversation_id,
		Author_id: ConversationMess.Author_id,
		Message: ConversationMess.Message,
		File: ConversationMess.File,
		Last_update: ConversationMess.Last_update,
	}
}

func ConversationMessagesDtoToConversationMessagesModel(ConversationMessages []*ConversationMessage) []*models.ConversationMessage {

	var ConversationMessagesModel []*models.ConversationMessage
	for _, val := range ConversationMessages {

		ConversationMessagesModel = append(ConversationMessagesModel, ConversationMessageDtoToConversationMessageModel(val))
	}
	return ConversationMessagesModel
}

func ConversationMessageModelToConversationMessageDto(ConversationMess *models.ConversationMessage) *ConversationMessage {

	if ConversationMess == nil {

		return nil
	}
	return &ConversationMessage{

		Message_id: ConversationMess.Message_id,
		Conversation_id: ConversationMess.Conversation_id,
		Author_id: ConversationMess.Author_id,
		Message: ConversationMess.Message,
		File: ConversationMess.File,
		Last_update: ConversationMess.Last_update,
	}
}

func ConversationMessagesModelToConversationMessagesDto(ConversationMessages []*models.ConversationMessage) []*ConversationMessage {

	var ConversationMessagesModel []*ConversationMessage
	for _, val := range ConversationMessages {

		ConversationMessagesModel = append(ConversationMessagesModel, ConversationMessageModelToConversationMessageDto(val))
	}
	return ConversationMessagesModel
}

func ConversationAuthorsSameIDToConversationAuthors(ListConversationAuthor []*ConversationAuthor) *ConversationAuthors {

	for i := 1; i < len(ListConversationAuthor); i++ {

		if ListConversationAuthor[i] == nil || ListConversationAuthor[i - 1] == nil {

			return nil
		}
		if ListConversationAuthor[i].Conversation_id != ListConversationAuthor[i - 1].Conversation_id {

			return nil
		}
	}

	var conversationAuthors = new(ConversationAuthors)
	conversationAuthors.Conversation_id = ListConversationAuthor[0].Conversation_id
	conversationAuthors.Author_ids = append(conversationAuthors.Author_ids, ListConversationAuthor...)
	return conversationAuthors
}

func ConversationReponsiveModelToDto(cr *models.ConversationResponsive) *ConversationResponsive {

	if cr == nil {

		return nil
	}

	return &ConversationResponsive{

		Conversation_id: cr.Conversation_id,
		IsPrivate: cr.IsPrivate,
		Name: cr.Name,
		Creator: cr.Creator,
		Authors: cr.Authors,
		Last_update: cr.Last_update,
	}
}

func ConversationResponsivesModelToDto(crs []*models.ConversationResponsive) []*ConversationResponsive {

	var crsDto []*ConversationResponsive
	for i := 0; i < len(crs); i++ {

		crsDto = append(crsDto, ConversationReponsiveModelToDto(crs[i]))
	}
	return crsDto
}