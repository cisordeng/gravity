package xenon

import "context"

type Entity struct {
	Ctx   context.Context
}

type Repository struct {
	Ctx context.Context
}

type Service struct {
	Ctx context.Context
}