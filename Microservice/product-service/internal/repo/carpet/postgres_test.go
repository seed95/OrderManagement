package carpet

import (
	"Product/internal/model"
	"Product/pkg/random"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCarpet_Duplicate(t *testing.T) {

	repo, err := NewCarpetRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	c1 := model.Carpet{
		DesignCode: random.StringWithCharset(5, "0123456789"),
		Color:      "قرمز",
		Sizes:      []string{"8"},
	}

	err = repo.NewCarpet(&c1)
	assert.Nil(t, err)

	err = repo.NewCarpet(&c1)
	assert.NotNil(t, err)

}
