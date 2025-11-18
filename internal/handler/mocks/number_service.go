package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type NumberService struct {
	mock.Mock
}

func NewNumberService(t mock.TestingT) *NumberService {
	m := &NumberService{}
	m.Mock.Test(t)
	return m
}

func (m *NumberService) AddAndList(ctx context.Context, value int) ([]int, error) {
	args := m.Called(ctx, value)
	var res []int
	if v, ok := args.Get(0).([]int); ok {
		res = v
	}
	return res, args.Error(1)
}
