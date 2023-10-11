package messagebus

import (
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

type Subscriber func(record url.URL) error
