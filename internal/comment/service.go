package comment

import (
	"context"
	"errors"
	"log"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type CommentService struct {
	Repository interfaces.IRepository
}

func (c *CommentService) FindById(id int) (*repository.Comment, error) {
	comment, err := c.Repository.FindComment(context.Background(), int64(id))
	if err != nil {
		return nil, customerror.CommentNotFound(err)
	}
	return &comment, nil
}

func (c *CommentService) Create(param repository.CreateCommentParams) (*repository.Comment, error) {
	comment, err := c.Repository.CreateComment(context.Background(), param)
	if err != nil {
		log.Println(err)
		return nil, customerror.CommentCreationFail(err)
	}
	return &comment, nil
}

func (c *CommentService) Update(userId int, param repository.UpdateCommentParams) (*repository.Comment, error) {
	comment, err := c.FindById(int(param.ID))
	if err != nil {
		return nil, customerror.CommentNotFound(err)
	}
	if int(*comment.UserID) != userId {
		return nil, customerror.CommentNotAllowed(errors.New(""))
	}

	updatedComment, err := c.Repository.UpdateComment(context.Background(), param)
	if err != nil {
		return nil, customerror.CommentUpdateFail(err)
	}
	return &updatedComment, nil
}

func (c *CommentService) Delete(userId int, id int) error {
	comment, err := c.FindById(id)
	if err != nil {
		return customerror.CommentNotFound(err)
	}

	if int(*comment.UserID) != userId {
		return customerror.CommentNotAllowed(errors.New(""))
	}
	err = c.Repository.DeleteComment(context.Background(), int64(id))
	if err != nil {
		return customerror.CommentDeleteFail(err)
	}
	return nil
}

func InitCommentService() *CommentService {
	return &CommentService{}
}

func (c CommentService) WithRepository(repo interfaces.IRepository) CommentService {
	c.Repository = repo
	return c
}
