package url

import "github.com/NordGus/shrtnr/domain/storage/url"

type Subscriber func(record url.URL) error
