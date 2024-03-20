package app

type (
	Repositories struct {
		UserRepo         UserRepo
		OrganizationRepo OrganizationRepo
		GroupRepo        GroupRepo
		TransactionRepo  TransactionRepo
		ServicePriceRepo ServicePriceRepo
	}

	Services struct {
		UserService         UserService
		OrganizationService OrganizationService
		GroupService        GroupService
		TransactionService  TransactionService
	}
)
