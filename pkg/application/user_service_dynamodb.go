package application

import (
	"context"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
)

type userServiceDynamoDb struct {
	UserRepository domain.UserRepositoryCtx
}

//ProvideUserServiceDynamoDb Provide user service via dynamodb
func ProvideUserServiceDynamoDb(u domain.UserRepositoryCtx) domain.UserServiceDynamoDb {
	return userServiceDynamoDb{UserRepository: u}
}

func (u userServiceDynamoDb) Update(ctx context.Context, user domain.User) error {
	err := u.UserRepository.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (u userServiceDynamoDb) Find(ctx context.Context, id string) (domain.User, error) {
	user, err := u.UserRepository.Find(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (u userServiceDynamoDb) Insert(ctx context.Context, user domain.User) error {
	err := u.UserRepository.Insert(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (u userServiceDynamoDb) Delete(ctx context.Context, id string) error {
	err := u.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
