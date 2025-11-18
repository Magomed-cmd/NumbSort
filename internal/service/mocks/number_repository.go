package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type NumberRepository struct {
	mock.Mock
}

func NewNumberRepository(t mock.TestingT) *NumberRepository {
	m := &NumberRepository{}
	m.Mock.Test(t)
	return m
}

func (m *NumberRepository) Insert(ctx context.Context, value int) error {
	args := m.Called(ctx, value)
	return args.Error(0)
}

func (m *NumberRepository) ListSorted(ctx context.Context) ([]int, error) {
	args := m.Called(ctx)
	var res []int
	if v, ok := args.Get(0).([]int); ok {
		res = v
	}
	return res, args.Error(1)
}
