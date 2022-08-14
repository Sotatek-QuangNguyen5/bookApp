package handlers

import (
	"bookApp/dto"
	"bookApp/errs"
	"bookApp/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


type ConversationHandler struct {

	services services.ConversationServices
}


func NewConversationHandler(services services.ConversationServices) ConversationHandler {

	return ConversationHandler{

		services: services,
	}
}


func (c ConversationHandler) GetListMess() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var Conversation = new(dto.Conversation)
		e := ctx.ShouldBindJSON(Conversation)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		author_id := ctx.MustGet("author_id").(int)
		err := c.services.GetConversationAuthor(&dto.ConversationAuthor{

			Conversation_id: Conversation.Conversation_id,
			Author_id: author_id,
		})
		if err != nil {

			WriteError(ctx, err)
			return
		}
		res, err := c.services.GetMessByIdConversation(Conversation)
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (c ConversationHandler) CreateRoomConversation() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var roomConversation = new(dto.RoomConversation)
		err := ctx.ShouldBindJSON(roomConversation)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		roomConversation.Name = strings.TrimSpace(roomConversation.Name)
		if len(roomConversation.Name) == 0 {

			WriteError(ctx, errs.BadRequestError("Lost Name Of Room!!!"))
			return
		}
		conver_id, e := c.services.CreateConversation(&dto.Conversation{

			Name: roomConversation.Name,
			IsPrivate: 0,
			Last_update: roomConversation.Last_update,
		})
		if e != nil {

			WriteError(ctx, e)
			return
		}
		var conversation_author = &dto.ConversationAuthor{

			Conversation_id: conver_id,
			Author_id: roomConversation.Author_id,
			IsCreator: 1,
			Last_update: roomConversation.Last_update,
		}
		e = c.services.CreateConversationAuthor(conversation_author)
		if e != nil {

			WriteError(ctx, e)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageCreateSuccess("Conversation Room"))
	}
}

func (c ConversationHandler) AddAuthorToRoomConversation() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var conversationAuthors = new(dto.ConversationAuthors)
		err := ctx.ShouldBindJSON(conversationAuthors)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		var conversation = &dto.Conversation{

			Conversation_id: conversationAuthors.Conversation_id,
		}
		conver, e := c.services.GetConversationById(conversation)
		if e != nil {

			WriteError(ctx, e)
			return
		}
		if conver.IsPrivate != 0 {

			WriteError(ctx, errs.NewUnauthenticatedError("Not Authorization!!!"))
			return
		}
		for _, author := range conversationAuthors.Author_ids {

			e := c.services.CreateConversationAuthor(&dto.ConversationAuthor{

				Conversation_id: conversationAuthors.Conversation_id,
				Author_id: author.Author_id,
				IsCreator: 0,
				Last_update: author.Last_update,
			})
			if e != nil {

				WriteError(ctx, e)
				return
			}
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageAddSuccess("Author To Conversation Room"))
	}
}

func (c ConversationHandler) AddMessageToConversation() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var message = new(dto.ConversationMessage)
		e := ctx.ShouldBindJSON(message)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		err := c.services.GetConversationAuthor(&dto.ConversationAuthor{

			Conversation_id: message.Conversation_id,
			Author_id: message.Author_id,
		})
		if err != nil {

			WriteError(ctx, errs.NewUnauthenticatedError("Not Authorization!!!"))
			return
		}
		err = c.services.AddMessage(message)
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageAddSuccess("Message To Conversation"))
	}
}

func (c ConversationHandler) CreatePrivateConversation() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var privateConversation = new(dto.PrivateConversation)
		e := ctx.ShouldBindJSON(privateConversation)
		if e != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		conversation_id, err := c.services.CreateConversation(&dto.Conversation{

			IsPrivate: 1,
			Last_update: privateConversation.Last_update,
		})
		if err != nil {

			WriteError(ctx, err)
			return
		}
		err = c.services.CreateConversationAuthor(&dto.ConversationAuthor{

			Conversation_id: conversation_id,
			Author_id: privateConversation.Author1_id,
			IsCreator: 1,
			Last_update: privateConversation.Last_update,
		})
		if err != nil {

			WriteError(ctx, err)
			return
		}
		err = c.services.CreateConversationAuthor(&dto.ConversationAuthor{

			Conversation_id: conversation_id,
			Author_id: privateConversation.Author2_id,
			IsCreator: 0,
			Last_update: privateConversation.Last_update,
		})
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageAddSuccess("Message To Conversation"))
	}
}