package dal

import (
	"testing"
	"wash-payment/internal/dal/dbmodels"

	"github.com/powerman/check"
)

func TestCreateUser(tt *testing.T) {
	t := check.T(tt)

	res1, err := repositories.UserRepo.Create(ctx, user1)
	t.Nil(err)
	t.DeepEqual(res1, user1)

	res2, err := repositories.UserRepo.Create(ctx, user2)
	t.Nil(err)
	t.DeepEqual(res2, user2)

	_, err = repositories.UserRepo.Create(ctx, user2)
	t.Err(err, dbmodels.ErrAlreadyExists)
}
