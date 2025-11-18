package service

import "context"

type NumberService interface {
	AddAndList(ctx context.Context, value int) ([]int, error)
}

type numberService struct {
	repo NumberRepository
}

type NumberRepository interface {
	Insert(ctx context.Context, value int) error
	ListSorted(ctx context.Context) ([]int, error)
}

func NewNumberService(repo NumberRepository) NumberService {
	return &numberService{repo: repo}
}

func (s *numberService) AddAndList(ctx context.Context, value int) ([]int, error) {
	if err := s.repo.Insert(ctx, value); err != nil {
		return nil, err
	}
	return s.repo.ListSorted(ctx)
}
