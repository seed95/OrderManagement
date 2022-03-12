package product

import (
	"Product/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProduct_ZeroId(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	p1 := model.Product{
		Id:          0,
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "102",
		Colors:      []string{"آبی", "قرمز"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۲",
	}

	err = repo.CreateProduct(&p1)
	assert.Nil(t, err)

}

func TestCreateProduct_Duplicate(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Duplicate design code
	t.Run("design code", func(t *testing.T) {
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "102",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۲",
		}

		err = repo.CreateProduct(&p1)
		assert.Nil(t, err)

		err = repo.CreateProduct(&p1)
		assert.NotNil(t, err)
	})

	// Duplicate size
	t.Run("size", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "103",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{"6", "9", "6"},
			Description: "توضیحات برای کد ۱۰۳",
		}

		err = repo.CreateProduct(&p1)
		assert.NotNil(t, err)
	})

	// Duplicate color
	t.Run("color", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "105",
			Colors:      []string{"آبی", "آبی"},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		err = repo.CreateProduct(&p1)
		assert.NotNil(t, err)
	})

	//TODO implement read from database is null
	// transaction is failed so not create product
}

func TestCreateProduct_Empty(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Empty company
	t.Run("company", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "",
			CompanyId:   0,
			DesignCode:  "103",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{},
			Description: "توضیحات برای کد ۱۰۳",
		}

		err = repo.CreateProduct(&p1)
		assert.Nil(t, err)
	})

	// Empty design code
	t.Run("design code", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{},
			Description: "توضیحات برای کد خالی",
		}

		err = repo.CreateProduct(&p1)
		assert.Nil(t, err)
	})

	// Empty company and design code
	t.Run("company and design code", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "",
			CompanyId:   0,
			DesignCode:  "",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{},
			Description: "توضیحات برای کد خالی",
		}

		err = repo.CreateProduct(&p1)
		assert.Nil(t, err)
	})

	// Empty size
	t.Run("size", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "103",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{},
			Description: "توضیحات برای کد ۱۰۳",
		}

		err = repo.CreateProduct(&p1)
		assert.Nil(t, err)
	})

	// Empty color
	t.Run("color", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "105",
			Colors:      []string{},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		err = repo.CreateProduct(&p1)
		assert.Nil(t, err)
	})
}

func TestCreateProduct_Nil(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Nil color
	t.Run("color", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "105",
			Colors:      nil,
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		err = repo.CreateProduct(&p1)
		assert.Nil(t, err)
	})

	// Nil size
	t.Run("size", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "103",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  nil,
			Description: "توضیحات برای کد ۱۰۳",
		}

		err = repo.CreateProduct(&p1)
		assert.Nil(t, err)
	})

}

func TestGetAllCarpet_Ok(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)

	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}

	err = repo.CreateProduct(&p1)
	assert.Nil(t, err)

	carpets, err := repo.GetAllCarpet(companyId)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(carpets))

}

func TestGetAllCarpet_EmptyColor(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)

	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}

	err = repo.CreateProduct(&p1)
	assert.Nil(t, err)

	carpets, err := repo.GetAllCarpet(companyId)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(carpets))

}

func TestGetAllCarpet_Empty(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)

	carpets, err := repo.GetAllCarpet(companyId)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(carpets))

}

func TestGetAllCarpet_InvalidCompanyId(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(0)

	carpets, err := repo.GetAllCarpet(companyId)
	assert.NotNil(t, err)
	assert.Nil(t, carpets)

}
