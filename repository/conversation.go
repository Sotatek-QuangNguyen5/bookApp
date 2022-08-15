package repository

import (
	"bookApp/errs"
	"bookApp/models"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

type ConversationRepository interface {

	Create(*models.Conversation) (int, *errs.AppError)
	CreateAuthor(*models.ConversationAuthor) (*errs.AppError)
	GetMessByIdConversation(int) ([]*models.ConversationMessage, *errs.AppError)
	AddMessage(*models.ConversationMessage) (*errs.AppError)
	GetAuthorById(int) ([]*models.ConversationAuthor, *errs.AppError)
	GetConversationById(int) (*models.Conversation, *errs.AppError)
	GetConversationAuthor(*models.ConversationAuthor) (*errs.AppError)
	GetFullConversationByAuthorId(int) ([]*models.ConversationResponsive, *errs.AppError)
}

type DefaultConversationRepository struct {

	db *sql.DB
}

func NewConversationRepository(db *sql.DB) ConversationRepository {

	return DefaultConversationRepository{

		db: db,
	}
}

func (c DefaultConversationRepository) Create(Conversation *models.Conversation) (int, *errs.AppError) {

	query := "INSERT INTO conversation(name, isPrivate, last_update) VALUES(?, ?, ?)";
	stmt, e := c.db.Prepare(query)
	if e != nil {

		return 0, errs.ErrorInsertData()
	}
	res, e := stmt.Exec(Conversation.Name, Conversation.IsPrivate, time.Now().Unix())
	if e != nil {

		return 0, errs.ErrorInsertData()
	}
	lastInsertID, e := res.LastInsertId()
	if e != nil {

		return 0, errs.ErrorInsertData()
	}
	return int(lastInsertID), nil
}

func (c DefaultConversationRepository) CreateAuthor(conversationAuthor *models.ConversationAuthor) (*errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM conversation_author WHERE conversation_id = %d AND author_id = %d", conversationAuthor.Conversation_id, conversationAuthor.Author_id)
	res, e := c.db.Query(query)
	if e != nil {

		return errs.ErrorGetData()
	}
	cnt := 0
	for res.Next() {

		cnt += 1
	}
	if cnt != 0 {

		return errs.BadRequestError("Data survived!!!")
	}
	query = fmt.Sprintf("INSERT INTO conversation_author(conversation_id, author_id, iscreator, last_update) VALUES(%d, %d, %d, %d)",
				conversationAuthor.Conversation_id, conversationAuthor.Author_id, conversationAuthor.IsCreator, conversationAuthor.Last_update)
	_, err := c.db.Query(query)

	if err != nil {

		return errs.ErrorInsertData()
	}

	return nil
}

func (c DefaultConversationRepository) AddMessage(ConversationMess *models.ConversationMessage) (*errs.AppError) {

	query := fmt.Sprintf("INSERT INTO  messages(conversation_id, author_id, message, file, last_update) VALUES(%d, %d, '%s', '%s', %d)", 
				ConversationMess.Conversation_id, ConversationMess.Author_id, ConversationMess.Message, ConversationMess.File, ConversationMess.Last_update)
	_, err := c.db.Query(query)

	if err != nil {

		return errs.ErrorInsertData()
	}

	return nil
}

func (c DefaultConversationRepository) GetMessByIdConversation(Conversation_id int) ([]*models.ConversationMessage, *errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM message WHERE conversation_id = %d", Conversation_id)
	res, err := c.db.Query(query)
	
	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var ConversationMessages []*models.ConversationMessage
	for res.Next() {

		var ConversationMessage = new(models.ConversationMessage)
		e := res.Scan(&ConversationMessage.Message_id, &ConversationMessage.Conversation_id,&ConversationMessage.Author_id, &ConversationMessage.Message, &ConversationMessage.File, &ConversationMessage.Last_update)
		if e != nil {

			return nil, errs.ErrorReadData()
		}
		ConversationMessages = append(ConversationMessages, ConversationMessage)
	}
	if len(ConversationMessages) == 0 {

		return nil, errs.ErrorDataNotSurvive()
	}

	return ConversationMessages, nil
}

func (c DefaultConversationRepository) GetAuthorById(Conversation_id int) ([]*models.ConversationAuthor, *errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM conversation_author WHERE conversation_id = %d", Conversation_id)
	res, err := c.db.Query(query)
	
	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var ConversationAuthors []*models.ConversationAuthor
	for res.Next() {

		var ConversationAuthor = new(models.ConversationAuthor)
		e := res.Scan(&ConversationAuthor.Conversation_id, &ConversationAuthor.Author_id, &ConversationAuthor.IsCreator, &ConversationAuthor.Last_update)
		if e != nil {

			return nil, errs.ErrorReadData()
		}
		ConversationAuthors = append(ConversationAuthors, ConversationAuthor)
	}
	if len(ConversationAuthors) == 0 {

		return nil, errs.ErrorDataNotSurvive()
	}

	return ConversationAuthors, nil
}

func (c DefaultConversationRepository) GetConversationById(conversation_id int) (*models.Conversation, *errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM conversation WHERE conversation_id = %d", conversation_id)
	res, e := c.db.Query(query)
	if e != nil {

		return nil, errs.ErrorReadData()
	}

	var conversation = new(models.Conversation)
	cnt := 0
	for res.Next() {

		e := res.Scan(&conversation.Conversation_id, &conversation.Name, &conversation.IsPrivate, &conversation.Last_update)
		if e != nil {

			return nil, errs.ErrorReadData()
		}
		cnt += 1
	}
	if cnt == 0 {

		return nil, errs.ErrorDataNotSurvive()
	}
	if cnt > 1 {

		return nil, errs.InternalServerError("Error conversation!!!")
	}
	return conversation, nil
}

func (c DefaultConversationRepository) GetConversationAuthor(conversationAuthor *models.ConversationAuthor) (*errs.AppError) {

	query := fmt.Sprintf("SELECT * FROM conversation_author WHERE conversation_id = %d AND author_id = %d", conversationAuthor.Conversation_id, conversationAuthor.Author_id)
	res, e := c.db.Query(query)
	if e != nil {

		return errs.ErrorGetData()
	}
	cnt := 0
	for res.Next() {

		cnt += 1
	}
	if cnt == 0 {

		return errs.ErrorDataNotSurvive()
	}
	if cnt > 1 {

		return errs.InternalServerError("Error conversation!!!")
	}
	return nil
}

func (c DefaultConversationRepository) GetFullConversationByAuthorId(author_id int) ([]*models.ConversationResponsive, *errs.AppError) {

	query := fmt.Sprintf("SELECT c.* FROM conversation c JOIN conversation_author ca ON c.conversation_id = ca.conversation_id AND ca.author_id = %d", author_id)
	res, e := c.db.Query(query)
	if e != nil {

		return nil, errs.ErrorGetData()
	}
	var conversations []*models.ConversationResponsive
	var conver_ids = "("
	for res.Next() {

		var conversation = new(models.ConversationResponsive)
		e = res.Scan(&conversation.Conversation_id, &conversation.Name, &conversation.IsPrivate, &conversation.Last_update)
		if e != nil {

			return nil, errs.ErrorReadData()
		}
		conversations = append(conversations, conversation)
		conver_ids = conver_ids + strconv.Itoa(conversation.Conversation_id) + ","
	}
	conver_ids = conver_ids[:len(conver_ids) - 1] + ")"
	query = "SELECT * FROM conversation_author WHERE conversation_id IN " + conver_ids
	res, e = c.db.Query(query)
	if e != nil {

		return nil, errs.ErrorGetData()
	}
	var conversationAuthors []*models.ConversationAuthor
	for res.Next() {

		var conversation_author = new(models.ConversationAuthor)
		e = res.Scan(&conversation_author.Conversation_id, &conversation_author.Author_id, &conversation_author.IsCreator, &conversation_author.Last_update)
		if e != nil {

			return nil, errs.ErrorReadData()
		}
		conversationAuthors = append(conversationAuthors, conversation_author)
	}
	for i := 0; i < len(conversations); i++ {

		for j := 0; j < len(conversationAuthors); j++ {

			if conversations[i].Conversation_id == conversationAuthors[j].Conversation_id {

				if conversationAuthors[j].IsCreator == 1 {

					conversations[i].Creator = conversationAuthors[j].Author_id
				}
				conversations[i].Authors = append(conversations[i].Authors, conversationAuthors[j].Author_id)
			}
		}
	}
	return conversations, nil
}