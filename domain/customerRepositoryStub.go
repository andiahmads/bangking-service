package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"21", "andi", "indonesia", "123123", "1996-05-05", "1"},
		{"21", "andi", "indonesia", "123123", "1996-05-05", "1"},
	}
	return CustomerRepositoryStub{customers}
}
