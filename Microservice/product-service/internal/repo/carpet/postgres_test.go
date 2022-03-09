package carpet

import (
	"Product/internal/derror"
	"Product/internal/model"
	"Product/pkg/random"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCarpet_Duplicate(t *testing.T) {
	repo, err := NewCarpetRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	c1 := model.Product{
		DesignCode: random.StringWithCharset(5, "0123456789"),
		Color:      "قرمز",
		Sizes:      []string{"8"},
	}

	err = repo.CreateCarpet(&c1)
	assert.Nil(t, err)

	err = repo.CreateCarpet(&c1)
	assert.NotNil(t, err)

	// Empty product
	c1 = model.Product{}

	err = repo.CreateCarpet(&c1)
	assert.Nil(t, err)

	err = repo.CreateCarpet(&c1)
	assert.NotNil(t, err)

}

func TestCreateCarpet_Empty(t *testing.T) {
	repo, err := NewCarpetRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	c1 := model.Product{}

	err = repo.CreateCarpet(&c1)
	assert.Nil(t, err)

}

func TestCreateCarpet_Nil(t *testing.T) {
	repo, err := NewCarpetRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.CreateCarpet(nil)
	assert.Equal(t, err, derror.NilCarpet)

}

func TestGetAllCarpets_Empty(t *testing.T) {
	repo, err := NewCarpetRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	//dCode1 := random.StringWithCharset(5, "0123456789")
	//dCode2 := random.StringWithCharset(5, "0123456789")
	//
	//c1 := model.Product{
	//	DesignCode: dCode1,
	//	Color:      "قرمز",
	//	Sizes:      []string{"8"},
	//}

	carpets, err := repo.GetAllCarpets()
	assert.Nil(t, err)
	assert.Equal(t, carpets, []model.Product{})

}
