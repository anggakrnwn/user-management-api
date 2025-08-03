package helpers

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestTranslateErrorMessage(t *testing.T) {
	t.Run("Validasi Error", func(t *testing.T) {
		type TestStruct struct {
			Name string `validate:"required"`
		}
		validator := validator.New()
		err := validator.Struct(TestStruct{})

		errors := TranslateErrorMessage(err)
		assert.Equal(t, "Name is required", errors["Name"])
	})

	t.Run("Duplikasi Data", func(t *testing.T) {
		// Simulasi error MySQL karena duplikasi data
		err := &mysql.MySQLError{
			Number:  1062,
			Message: "Duplicate entry 'testuser' for key 'uni_users_username'",
		}

		errors := TranslateErrorMessage(err)
		assert.Equal(t, "username already exists", errors["username"])
	})
}
