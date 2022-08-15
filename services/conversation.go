package services

import (

	"bookApp/dto"
	"bookApp/errs"
	"bookApp/repository"
)

type ConversationServices interface {

	CreateConversation(*dto.Conversation) (int, *errs.AppError)
	CreateConversationAuthor(*dto.ConversationAuthor) (*errs.AppError)
	GetMessByIdConversation(*dto.Conversation) ([]*dto.ConversationMessage, *errs.AppError)
	GetAuthorById(*dto.Conversation) (*dto.ConversationAuthors, *errs.AppError)
	AddMessage(*dto.ConversationMessage) (*errs.AppError)
	GetConversationById(*dto.Conversation) (*dto.Conversation, *errs.AppError)
	GetConversationAuthor(*dto.ConversationAuthor) (*errs.AppError)
	//GetFullConversationByAuthorId()
}

type DefaultConversationServices struct {

	repo repository.ConversationRepository
}

func NewConversationServices(repo repository.ConversationRepository) ConversationServices {

	return DefaultConversationServices{

		repo: repo,
	}
}

func (c DefaultConversationServices) CreateConversation(Conversation *dto.Conversation) (int, *errs.AppError) {

	// err := dto.CheckID(Conversation.Author_id)
	// if err != nil {

	// 	return 0, err
	// }
	return c.repo.Create(dto.ConversationDtoToConversationModel(Conversation))
}

func (c DefaultConversationServices) CreateConversationAuthor(ConversationAuthor *dto.ConversationAuthor) (*errs.AppError) {

	err := dto.CheckID(ConversationAuthor.Conversation_id)
	if err != nil {

		return err
	}
	return c.repo.CreateAuthor(dto.ConversationAuthorDtoToConversationAuthorModel(ConversationAuthor))
}


func (c DefaultConversationServices) GetMessByIdConversation(Conversation *dto.Conversation) ([]*dto.ConversationMessage, *errs.AppError) {

	err := dto.CheckID(Conversation.Conversation_id)
	if err != nil {

		return nil, err
	}
	
	res, e := c.repo.GetMessByIdConversation(Conversation.Conversation_id)
	if e != nil {

		return nil, e
	}
	ConversationMessagesDto := dto.ConversationMessagesModelToConversationMessagesDto(res)
	return ConversationMessagesDto, nil
}

func (c DefaultConversationServices) GetAuthorById(Conversation *dto.Conversation) (*dto.ConversationAuthors, *errs.AppError) {

	err := dto.CheckID(Conversation.Conversation_id)
	if err != nil {

		return nil, err
	}
	res, err := c.repo.GetAuthorById(Conversation.Conversation_id)
	if err != nil {

		return nil, err
	}
	ConversationAuthorsDto := dto.ConversationAuthorsModelToConversationAuthorsDto(res)
	ListConversationAuthor := dto.ConversationAuthorsSameIDToConversationAuthors(ConversationAuthorsDto)
	return ListConversationAuthor, nil
}

func (c DefaultConversationServices) AddMessage(ConversationMess *dto.ConversationMessage) (*errs.AppError) {

	err := dto.CheckID(ConversationMess.Conversation_id)
	if err != nil {

		return err
	}
	return c.repo.AddMessage(dto.ConversationMessageDtoToConversationMessageModel(ConversationMess))
}

func (c DefaultConversationServices) GetConversationById(conversation *dto.Conversation) (*dto.Conversation, *errs.AppError) {

	err := dto.CheckID(conversation.Conversation_id)
	if err != nil {

		return nil, err
	}
	res, err := c.repo.GetConversationById(conversation.Conversation_id)
	if err != nil {

		return nil, err
	}
	return dto.ConversationModelToConversationDto(res), nil
}

func (c DefaultConversationServices) GetConversationAuthor(conver *dto.ConversationAuthor) (*errs.AppError) {

	conversation := dto.ConversationAuthorDtoToConversationAuthorModel(conver)
	return c.repo.GetConversationAuthor(conversation)
}