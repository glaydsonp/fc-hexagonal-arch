package db

import (
	"database/sql"

	"github.com/glaydsonp/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name, status, price from products where id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Status, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("select count(*) from products where id = ?", product.GetID()).Scan(&rows)

	if rows == 0 {
		return p.create(product)
	}

	return p.update(product)
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("insert into products(id, name, status, price) values(?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetStatus(), product.GetPrice())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("update products set name = ?, status = ?, price = ? where id = ?")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetName(), product.GetStatus(), product.GetPrice(), product.GetID())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}
