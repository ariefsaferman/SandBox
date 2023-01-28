package handler

import (
	"net/http"
	"url-shortener/internal/component/config"
	"url-shortener/internal/dto/request"
	"url-shortener/internal/usecase"
	uresponse "url-shortener/internal/utils/response"

	"url-shortener/internal/dto/response"
	cerrors "url-shortener/internal/utils/errors"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	cfg         config.Config
	itemUsecase usecase.ItemUsecase
}

func NewItemHandler(itemUsecase usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{
		itemUsecase: itemUsecase,
	}
}

func (h *ItemHandler) RedirectBySlug(c *gin.Context) {
	ctx := c.Request.Context()

	slug := c.Param("slug")

	m, err := h.itemUsecase.GetOneBySlug(ctx, slug)
	if err != nil {
		uresponse.ResponseError(c, err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, m.LongUrl)
}

func (h *ItemHandler) GetOneItemBySlug(c *gin.Context) {
	ctx := c.Request.Context()

	slug := c.Param("slug")

	var (
		res response.ItemGetOneBySlugResponse
	)

	m, err := h.itemUsecase.GetOneBySlug(ctx, slug)
	if err != nil {
		uresponse.ResponseError(c, err)
		return
	}

	res.ConvertFromModel(h.cfg, m)

	uresponse.ResponseSuccessWithoutMessage(c, res)
}

func (h *ItemHandler) GetOneItemBySlugWithJob(c *gin.Context) {
	ctx := c.Request.Context()

	slug := c.Param("slug")

	var (
		res response.ItemGetOneBySlugResponse
	)

	m, err := h.itemUsecase.GetOneBySlugWithJob(ctx, slug)
	if err != nil {
		uresponse.ResponseError(c, err)
		return
	}

	res.ConvertFromModel(h.cfg, m)

	uresponse.ResponseSuccessWithoutMessage(c, res)
}

func (h *ItemHandler) GetOneItemCachedBySlug(c *gin.Context) {
	ctx := c.Request.Context()

	slug := c.Param("slug")

	var (
		res response.ItemGetOneBySlugResponse
	)

	m, err := h.itemUsecase.GetOneCachedBySlug(ctx, slug)
	if err != nil {
		uresponse.ResponseError(c, err)
		return
	}

	res.ConvertFromModel(h.cfg, m)

	uresponse.ResponseSuccessWithoutMessage(c, res)
}

func (h *ItemHandler) GetOneItemCachedBySlugWithJob(c *gin.Context) {
	ctx := c.Request.Context()

	slug := c.Param("slug")

	var (
		res response.ItemGetOneBySlugResponse
	)

	m, err := h.itemUsecase.GetOneCachedBySlugWithJob(ctx, slug)
	if err != nil {
		uresponse.ResponseError(c, err)
		return
	}

	res.ConvertFromModel(h.cfg, m)

	uresponse.ResponseSuccessWithoutMessage(c, res)
}

func (h *ItemHandler) CreateOneItem(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		req request.ItemCreateOneRequest
		res response.ItemCreateOneResponse
	)

	err := c.ShouldBind(&req)
	if err != nil {
		uresponse.ResponseError(c, cerrors.NewBadFormattedRequest())
		return
	}

	m := req.ConvertToModel()
	cm, err := h.itemUsecase.CreateOne(ctx, m)
	if err != nil {
		uresponse.ResponseError(c, err)
		return
	}

	res.ConvertFromModel(cm)

	uresponse.ResponseSuccessWithoutMessage(c, res)
}
