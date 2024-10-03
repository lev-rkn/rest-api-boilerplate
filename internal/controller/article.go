package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"rest-api-service/internal/domain"
	"rest-api-service/internal/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type createArticleResponse struct {
	Id int `json:"id"`
}

type articleController struct {
	ctx            context.Context
	articleService articleService
}

func InitArticleController(
	ctx context.Context,
	articleService articleService,
	router *gin.RouterGroup,
) *articleController {
	articleController := &articleController{
		ctx:            ctx,
		articleService: articleService,
	}
	router.POST("/create/", articleController.CreateArticle)

	return articleController
}

// CreateArticle godoc
//
//	@Summary	CreateArticle creates new article
//	@Tags		articles
//	@Accept		json
//	@Produce	json
//	@Param		article	body		domain.Article	true "Article"
//	@Success	201		{object}	createArticleResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/article/create/ [post]
func (a *articleController) CreateArticle(c *gin.Context) {
	article := &domain.Article{}
	err := json.NewDecoder(c.Request.Body).Decode(&article)
	if err != nil {
		logger.ErrorLog("unable to decode article", err)
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(article)
	if err != nil {
		logger.ErrorLog("validate article", err)
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	id, err := a.articleService.CreateArticle(article)
	if err != nil {
		// if errors.Is(err, domain.ErrArticleNotFound) {
		// 	logger.ErrorLog("a.articleService.CreateArticle", err)
		// 	errorResponse(c, http.StatusNotFound, err)
		// 	return
		// }
		logger.ErrorLog("article creating by service", err)
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, createArticleResponse{Id: id})
}
