package dal

import (
	"testing"
	"time"
	"wash-payment/internal/app"
	"wash-payment/internal/app/entity"

	"github.com/powerman/check"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransaction(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	amount := int64(100)

	organization1 := generateOrganization(10000, 1)
	transaction1 := generateTransaction(entity.DepositOperation, amount, organization1.ID)
	transaction2 := generateTransaction(entity.DebitOperation, amount, organization1.ID)
	transaction3 := generateTransaction(entity.DebitOperation, 100000, organization1.ID)
	transaction4 := generateTransaction(entity.DebitOperation, 100000, uuid.NewV4())

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	res1, err := repositories.TransactionRepo.Create(ctx, transaction1)
	t.Nil(err)
	t.Equal(res1.CreatedAt, transaction1.CreatedAt)
	transaction1.CreatedAt = res1.CreatedAt
	t.DeepEqual(res1, transaction1)

	organization1.Balance += amount
	orgDB, err := repositories.OrganizationRepo.Get(ctx, organization1.ID)
	t.Nil(err)
	t.DeepEqual(orgDB, organization1)

	res2, err := repositories.TransactionRepo.Create(ctx, transaction2)
	t.Nil(err)
	t.Equal(res2.CreatedAt, transaction2.CreatedAt)
	transaction2.CreatedAt = res2.CreatedAt
	t.DeepEqual(res2, transaction2)

	organization1.Balance -= amount
	orgDB, err = repositories.OrganizationRepo.Get(ctx, organization1.ID)
	t.Nil(err)
	t.DeepEqual(orgDB, organization1)

	_, err = repositories.TransactionRepo.Create(ctx, transaction3)
	t.Err(err, app.ErrInsufficientFunds)

	_, err = repositories.TransactionRepo.Create(ctx, transaction4)
	t.Err(err, app.ErrNotFound)

	_, err = repositories.TransactionRepo.Create(ctx, transaction1)
	t.Err(err, app.ErrAlreadyExists)
}

func TestGetTransaction(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var organization1 = generateOrganization(10000, 1)
	var transaction1 = generateTransaction(entity.DepositOperation, 100, organization1.ID)
	var transaction2 = generateTransaction(entity.DepositOperation, 100, organization1.ID)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.TransactionRepo.Create(ctx, transaction1)
	t.Nil(err)

	resGet1, err := repositories.TransactionRepo.Get(ctx, transaction1.ID)
	t.Nil(err)
	t.Equal(resGet1.CreatedAt, transaction1.CreatedAt)
	transaction1.CreatedAt = resGet1.CreatedAt
	t.DeepEqual(resGet1, transaction1)

	_, err = repositories.TransactionRepo.Get(ctx, transaction2.ID)
	t.Err(err, app.ErrNotFound)
}

func TestListTransaction(tt *testing.T) {
	t := check.T(tt)
	err := truncate()
	t.Nil(err)

	var organization1 = generateOrganization(10000, 1)
	var organization2 = generateOrganization(10000, 1)
	var transaction1 = generateTransaction(entity.DepositOperation, 100, organization1.ID)
	var transaction2 = generateTransaction(entity.DepositOperation, 100, organization1.ID)
	transaction2.CreatedAt = transaction2.CreatedAt.Add(time.Second)

	_, err = repositories.OrganizationRepo.Create(ctx, organization1)
	t.Nil(err)

	_, err = repositories.OrganizationRepo.Create(ctx, organization2)
	t.Nil(err)

	_, err = repositories.TransactionRepo.Create(ctx, transaction1)
	t.Nil(err)

	_, err = repositories.TransactionRepo.Create(ctx, transaction2)
	t.Nil(err)

	filter := entity.TransactionFilter{
		OrganizationID: organization1.ID,
		Filter: entity.Filter{
			Page:     1,
			PageSize: 10,
		},
	}
	list, err := repositories.TransactionRepo.List(ctx, filter)
	t.Nil(err)
	t.Equal(list.Page, filter.Page)
	t.Equal(list.PageSize, filter.PageSize)
	t.Equal(list.TotalItems, 2)
	t.Equal(list.TotalPages, 1)
	t.Equal(len(list.Items), 2)
	t.Equal(list.Items[0].CreatedAt, transaction2.CreatedAt)
	t.Equal(list.Items[1].CreatedAt, transaction1.CreatedAt)
	list.Items[0].CreatedAt = transaction2.CreatedAt
	list.Items[1].CreatedAt = transaction1.CreatedAt
	t.DeepEqual(list.Items, []entity.Transaction{transaction2, transaction1})

	filter = entity.TransactionFilter{
		OrganizationID: organization1.ID,
		Filter: entity.Filter{
			Page:     10,
			PageSize: 10,
		},
	}
	list, err = repositories.TransactionRepo.List(ctx, filter)
	t.Nil(err)
	t.Equal(list.Page, filter.Page)
	t.Equal(list.PageSize, filter.PageSize)
	t.Equal(list.TotalItems, 2)
	t.Equal(list.TotalPages, 1)
	t.DeepEqual(list.Items, []entity.Transaction{})

	filter = entity.TransactionFilter{
		OrganizationID: organization1.ID,
		Filter: entity.Filter{
			Page:     1,
			PageSize: 1,
		},
	}
	list, err = repositories.TransactionRepo.List(ctx, filter)
	t.Nil(err)
	t.Equal(list.Page, filter.Page)
	t.Equal(list.PageSize, filter.PageSize)
	t.Equal(list.TotalItems, 2)
	t.Equal(list.TotalPages, 2)
	t.Equal(len(list.Items), 1)
	t.Equal(list.Items[0].CreatedAt, transaction2.CreatedAt)
	list.Items[0].CreatedAt = transaction2.CreatedAt
	t.DeepEqual(list.Items, []entity.Transaction{transaction2})

	filter = entity.TransactionFilter{
		OrganizationID: organization1.ID,
		Filter: entity.Filter{
			Page:     2,
			PageSize: 1,
		},
	}
	list, err = repositories.TransactionRepo.List(ctx, filter)
	t.Nil(err)
	t.Equal(list.Page, filter.Page)
	t.Equal(list.PageSize, filter.PageSize)
	t.Equal(list.TotalItems, 2)
	t.Equal(list.TotalPages, 2)
	t.Equal(len(list.Items), 1)
	t.Equal(list.Items[0].CreatedAt, transaction1.CreatedAt)
	list.Items[0].CreatedAt = transaction1.CreatedAt
	t.DeepEqual(list.Items, []entity.Transaction{transaction1})

	filter = entity.TransactionFilter{
		OrganizationID: organization2.ID,
		Filter: entity.Filter{
			Page:     1,
			PageSize: 10,
		},
	}
	list, err = repositories.TransactionRepo.List(ctx, filter)
	t.Nil(err)
	t.Equal(list.Page, filter.Page)
	t.Equal(list.PageSize, filter.PageSize)
	t.Equal(list.TotalItems, 0)
	t.Equal(list.TotalPages, 0)
	t.DeepEqual(list.Items, []entity.Transaction{})
}
