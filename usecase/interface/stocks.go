package usecase

type Stocks interface {
	GetStockByName(string, string) (any, error)
}
