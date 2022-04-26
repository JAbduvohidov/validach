package validach

import (
	"testing"
)

type User struct {
	Age            uint8  `validate:"gte=8,lte=130"`
	Email          string `json:"e-mail" validate:"required,email"`
	FavouriteColor string `validate:"hexcolor|rgb|rgba"`
}

func TestValidate(t *testing.T) {
	validUser := &User{
		Age:            30,
		Email:          "you@example.com",
		FavouriteColor: "#000",
	}

	invalidUser := &User{
		Age:            3,
		Email:          "don't have one",
		FavouriteColor: "wooden color",
	}

	errs := Validate(validUser)
	if errs != nil {
		t.Error("found error")
		return
	}

	errs = Validate(invalidUser)
	if errs == nil {
		t.Error("errors not found")
	}
}
