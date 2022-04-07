package api

import (
	"fmt"
	"reflect"
	"testing"
)

type PetValidate struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
	Kind string `json:"kind" validate:"required"`
}

type ValidateArgs struct {
	data PetValidate
}

var multipleRequiredNotFilled = ValidateArgs{
	data: PetValidate{
		ID:   123,
		Name: "",
		Kind: "",
	},
}

var multipleRequiredResponse = []*ErrorResponse{
	{
		FailedField: "PetValidate.Name",
		Tag:         "required",
		Value:       "",
	},
	{
		FailedField: "PetValidate.Kind",
		Tag:         "required",
		Value:       "",
	},
}

var singleRequiredNotFilled = ValidateArgs{
	data: PetValidate{
		ID:   123,
		Name: "Toby",
		Kind: "",
	},
}

var singleRequiredResponse = []*ErrorResponse{
	{
		FailedField: "PetValidate.Kind",
		Tag:         "required",
		Value:       "",
	},
}

var requiredFieldsFilled = ValidateArgs{
	data: PetValidate{
		ID:   123,
		Name: "Toby",
		Kind: "Dog",
	},
}

func TestValidateStruct(t *testing.T) {

	tests := []struct {
		description string
		args        ValidateArgs
		want        []*ErrorResponse
	}{
		{
			description: "With multiple required fields not filled",
			args:        multipleRequiredNotFilled,
			want:        multipleRequiredResponse,
		},
		{
			description: "With single required fields not filled",
			args:        singleRequiredNotFilled,
			want:        singleRequiredResponse,
		},
		{
			description: "With all required fields filled",
			args:        requiredFieldsFilled,
			want:        nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if got := ValidateStruct(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("%v", &got)
				fmt.Println(&tt.want)
				t.Errorf("ValidateStruct() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
