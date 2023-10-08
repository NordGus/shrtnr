package messagebus

import (
	"context"
	"github.com/NordGus/rom-stack/server/messagebus/url/created"
	"github.com/NordGus/rom-stack/server/messagebus/url/deleted"
)

func Start(ctx context.Context) {
	created.Start(ctx)
	deleted.Start(ctx)
}
