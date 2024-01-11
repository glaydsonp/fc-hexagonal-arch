package application

type ProductService struct {
	Persistance ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistance.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()
	if err != nil {
		return &Product{}, err
	}
	product, err = s.Persistance.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	product.Status = ENABLED
	product, err := s.Persistance.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {

	product.Status = DISABLED
	product, err := s.Persistance.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}