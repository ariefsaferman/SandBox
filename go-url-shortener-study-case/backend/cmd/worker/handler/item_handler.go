package handler

import (
	"context"
	"encoding/json"
	"log"
	"url-shortener/internal/constants"
	"url-shortener/internal/dto/queuepayload"
	"url-shortener/internal/usecase"

	cerrors "url-shortener/internal/utils/errors"

	"github.com/hibiken/asynq"
)

type ItemHandler struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemHandler(itemUsecase usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{
		itemUsecase: itemUsecase,
	}
}

func (h *ItemHandler) HandleViewCountTask(ctx context.Context, t *asynq.Task) error {
	var p queuepayload.ViewCountPayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return cerrors.NewErrorWrapper(cerrors.CodeServerError, "json unmarshal failed", err)
	}

	err := h.itemUsecase.IncrementViewCountById(ctx, p.ItemId)
	if err != nil {
		return cerrors.NewErrorWrapper(cerrors.CodeServerError, "increment view count by id usecase error", err)
	}

	log.Printf("message for %s processed", constants.TaskTypeViewCount)

	return nil
}
