package url

import "github.com/NordGus/shrtnr/server/storage/url"

type Subscriber func(record url.URL) error
