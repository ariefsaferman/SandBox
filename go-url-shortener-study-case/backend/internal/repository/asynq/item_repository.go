package asynq

import (
	"context"
	"encoding/json"
	"url-shortener/internal/constants"
	payload "url-shortener/internal/dto/queuepayload"
	"url-shortener/internal/repository"
	cerrors "url-shortener/internal/utils/errors"

	"github.com/hibiken/asynq"
)

type itemAsynqRepository struct {
	taskqueue *asynq.Client
}

func NewItemAsynqRepository(
	taskqueue *asynq.Client,
) repository.ItemQueueRepository {
	return &itemAsynqRepository{
		taskqueue: taskqueue,
	}
}

func (r *itemAsynqRepository) IncrementViewCountById(ctx context.Context, id int64) error {
	payload, err := json.Marshal(payload.ViewCountPayload{ItemId: id})
	if err != nil {
		return cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to create view count task", err)
	}

	task := asynq.NewTask(constants.TaskTypeViewCount, payload)

	_, err = r.taskqueue.Enqueue(task)
	if err != nil {
		return cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to enqueue view count task", err)
	}

	return nil
}
