package dto

import "bookApp/models"

type Author struct {
	
	Author_id int    `json:"author_id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Authentication struct {

	AccessToken	string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Message		string `json:"message"`
}

func AuthorModelToAuthorDto(authorModel *models.Author) *Author {

	if authorModel == nil {

		return nil
	}
	return &Author{

		Author_id: authorModel.Author_id,
		Name: authorModel.Name,
		Email: authorModel.Email,
		Phone: authorModel.Phone,
	}
}

func AuthorsModelToAuthorsDto(authorsModel []*models.Author) []*Author {

	var authorsDto []*Author
	for _, author := range authorsModel {

		authorsDto = append(authorsDto, AuthorModelToAuthorDto(author))
	}
	return authorsDto
}

func AuthorDtoToAuthorModel(authorDto *Author) *models.Author {

	if authorDto == nil {

		return nil
	}
	return &models.Author{

		Author_id: authorDto.Author_id,
		Name: authorDto.Name,
		Email: authorDto.Email,
		Phone: authorDto.Phone,
		Password: authorDto.Password,
	}
}

func AuthorsDtoToAuthorsModel(authorsDto []*Author) []*models.Author {

	var authors []*models.Author
	for _, author := range authorsDto {

		authors = append(authors, AuthorDtoToAuthorModel(author))
	}
	return authors
}