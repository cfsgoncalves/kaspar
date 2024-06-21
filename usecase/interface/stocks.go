package usecase

type Stocks interface {
	GetStockByName(string, string) (string, error)
}
