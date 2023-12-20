package types_test

import (
	"matryer/internal/types"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserName(t *testing.T) {
	table := []struct {
		title     string
		wantError bool
		err       validation.Error
		name      types.Name
	}{
		{
			title:     "cannot be blank",
			wantError: true,
			err:       validation.ErrRequired,
			name:      types.Name(""),
		},
		{
			title:     "the length must be between 1 and 10",
			wantError: true,
			err:       validation.ErrLengthOutOfRange,
			name:      types.Name("Jorge Luis Alonso Hdez"),
		},
		{
			title:     "success",
			wantError: false,
			err:       nil,
			name:      types.Name("Jorge Luis"),
		},
	}
	var verr validation.Error
	for _, tt := range table {
		t.Run(tt.title, func(t *testing.T) {
			err := validation.Validate(tt.name, tt.name.Rules()...)
			if !tt.wantError {
				assert.Equal(t, tt.err, err)
			} else {
				assert.ErrorAs(t, err, &verr)
				assert.Equal(t, tt.err.Code(), verr.Code())
			}
		})
	}
}
