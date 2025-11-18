package service

import (
	"context"
	"errors"
	"testing"

	"numbsort/internal/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNumberService_AddAndList(t *testing.T) {
	repo := mocks.NewNumberRepository(t)
	repo.On("Insert", mock.Anything, 2).Return(nil)
	repo.On("ListSorted", mock.Anything).Return([]int{1, 2, 3}, nil)

	svc := NewNumberService(repo)
	values, err := svc.AddAndList(context.Background(), 2)
	require.NoError(t, err)
	require.Equal(t, []int{1, 2, 3}, values)
	repo.AssertExpectations(t)
}

func TestNumberService_InsertError(t *testing.T) {
	repo := mocks.NewNumberRepository(t)
	repo.On("Insert", mock.Anything, 2).Return(errors.New("fail"))

	svc := NewNumberService(repo)
	_, err := svc.AddAndList(context.Background(), 2)
	require.Error(t, err)
	repo.AssertExpectations(t)
}

func TestNumberService_ListError(t *testing.T) {
	repo := mocks.NewNumberRepository(t)
	repo.On("Insert", mock.Anything, 2).Return(nil)
	repo.On("ListSorted", mock.Anything).Return(nil, errors.New("fail"))

	svc := NewNumberService(repo)
	_, err := svc.AddAndList(context.Background(), 2)
	require.Error(t, err)
	repo.AssertExpectations(t)
}
