package messagebus

import (
	"context"
	"github.com/NordGus/shrtnr/domain/messagebus/url/created"
	"github.com/NordGus/shrtnr/domain/messagebus/url/deleted"
)

func Start(ctx context.Context) {
	created.Start(ctx)
	deleted.Start(ctx)
}
