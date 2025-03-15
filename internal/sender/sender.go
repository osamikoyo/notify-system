package sender

import (
	"context"

	"github.com/koyo-os/notify-system/internal/models"
)

type Sender interface{
	Send(models.Notify, context.Context) error
}