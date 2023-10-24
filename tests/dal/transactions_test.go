package dal

import (
	"testing"
	"wash-payment/internal/dal/dbmodels"

	"github.com/powerman/check"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransaction(tt *testing.T) {
	t := check.T(tt)

	amount := int64(100)

	organization1 := generateOrganization(10000, 1)
	transaction1 := generateTransaction(dbmodels.DepositOperation, amount, organization1.ID)
	transaction2 := generateTransaction(dbmodels.DebitOperation, amount, organization1.ID)
	transaction3 := generateTransaction(dbmodels.DebitOperation, 100000, organization1.ID)
	transaction4 := generateTransaction(dbmodels.DebitOperation, 100000, uuid.NewV4())

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	// ---

	res1, err := repositories.TransactionRepo.Create(ctx, transaction1)
	t.Nil(err)
	t.Equal(res1.CreatedAt, transaction1.CreatedAt)
	transaction1.CreatedAt = res1.CreatedAt
	t.DeepEqual(res1, transaction1)

	organization1.Balance += amount
	orgDB, err := repositories.OrganizationRepo.Get(ctx, organization1.ID)
	t.Nil(err)
	t.DeepEqual(orgDB, organization1)

	// ---

	res2, err := repositories.TransactionRepo.Create(ctx, transaction2)
	t.Nil(err)
	t.Equal(res2.CreatedAt, transaction2.CreatedAt)
	transaction2.CreatedAt = res2.CreatedAt
	t.DeepEqual(res2, transaction2)

	organization1.Balance -= amount
	orgDB, err = repositories.OrganizationRepo.Get(ctx, organization1.ID)
	t.Nil(err)
	t.DeepEqual(orgDB, organization1)

	// ---

	_, err = repositories.TransactionRepo.Create(ctx, transaction3)
	t.Err(err, dbmodels.ErrInsufficientFunds)

	// ---

	_, err = repositories.TransactionRepo.Create(ctx, transaction4)
	t.Err(err, dbmodels.ErrNotFound)

	// ---

	_, err = repositories.TransactionRepo.Create(ctx, transaction1)
	t.Err(err, dbmodels.ErrAlreadyExists)
}

func TestGetTransaction(tt *testing.T) {
	t := check.T(tt)

	var organization1 = generateOrganization(10000, 1)
	var transaction1 = generateTransaction(dbmodels.DepositOperation, 100, organization1.ID)
	var transaction2 = generateTransaction(dbmodels.DepositOperation, 100, organization1.ID)

	_, err := repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.TransactionRepo.Create(ctx, transaction1)
	t.Nil(err)

	resGet1, err := repositories.TransactionRepo.Get(ctx, transaction1.ID)
	t.Nil(err)
	t.Equal(resGet1.CreatedAt, transaction1.CreatedAt)
	transaction1.CreatedAt = resGet1.CreatedAt
	t.DeepEqual(resGet1, transaction1)

	_, err = repositories.TransactionRepo.Get(ctx, transaction2.ID)
	t.Err(err, dbmodels.ErrNotFound)
}
