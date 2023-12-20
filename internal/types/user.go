package types

import (
	"matryer/pkg/joker"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Name string

func (n Name) Rules() []validation.Rule {
	return []validation.Rule{validation.Required, validation.Length(1, 10)}
}

func (n Name) Value() string {
	return string(n)
}

type FirstName string

func (fn FirstName) Rules() []validation.Rule {
	return []validation.Rule{validation.Required, validation.Length(1, 10)}
}

func (fn FirstName) Value() string {
	return string(fn)
}

type LastName string

func (ln LastName) Rules() []validation.Rule {
	return []validation.Rule{validation.Required, validation.Length(1, 10)}
}

func (ln LastName) Value() string {
	return string(ln)
}

type Email string

func (e Email) Rules() []validation.Rule {
	return []validation.Rule{validation.Required, is.Email}
}

func (e Email) Value() string {
	return string(e)
}

type Phone string

func (p Phone) Rules() []validation.Rule {
	return []validation.Rule{validation.Required, is.Int, validation.Length(10, 10)}
}

func (p Phone) Value() string {
	return string(p)
}

type Age int32

func (a Age) Rules() []validation.Rule {
	return []validation.Rule{validation.Required, validation.Min(1), validation.Max(100)}
}

func (a Age) Value() int32 {
	return int32(a)
}

type User struct {
	Name      Name      `json:"name"`
	FirstName FirstName `json:"first_name"`
	LastName  LastName  `json:"last_name"`
	Email     Email     `json:"email"`
	Phone     Phone     `json:"phone"`
	Age       Age       `json:"age"`
}

func (u *User) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name, u.Name.Rules()...),
		validation.Field(&u.FirstName, u.FirstName.Rules()...),
		validation.Field(&u.LastName, u.LastName.Rules()...),
		validation.Field(&u.Email, u.Email.Rules()...),
		validation.Field(&u.Age, u.Age.Rules()...),
		validation.Field(&u.Phone, validation.When(u.Age.Value() > 18, u.Phone.Rules()...)),
	)

	if err != nil {
		return joker.WrapErrorf(err, joker.CodeInvalidArgument, "invalid params")
	}

	return nil
}
