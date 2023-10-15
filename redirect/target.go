package redirect

import "time"

type Target struct {
	id         string
	uuid       string
	redirectTo string
	createdAt  time.Time
	deletedAt  time.Time
}

func (t Target) SetID(id string) Target {
	t.id = id

	return t
}

func (t Target) SetUUID(uuid string) Target {
	t.uuid = uuid

	return t
}

func (t Target) SetTarget(target string) Target {
	t.redirectTo = target

	return t
}

func (t Target) SetCreatedAt(createdAt time.Time) Target {
	t.createdAt = createdAt

	return t
}

func (t Target) SetDeletedAt(deletedAt time.Time) Target {
	t.deletedAt = deletedAt

	return t
}
