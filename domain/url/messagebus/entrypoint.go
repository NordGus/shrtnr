package messagebus

import (
	"context"

	"github.com/NordGus/shrtnr/domain/url/messagebus/created"
	"github.com/NordGus/shrtnr/domain/url/messagebus/deleted"
)

func Start(ctx context.Context) {
	created.Start(ctx)
	deleted.Start(ctx)
}
