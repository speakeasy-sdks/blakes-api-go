// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/blakes-api-go/pkg/models/shared"
	"github.com/speakeasy-sdks/blakes-api-go/pkg/utils"
	"net/http"
)

type PatchPetsResponseBodyType string

const (
	PatchPetsResponseBodyTypePetByAge  PatchPetsResponseBodyType = "PetByAge"
	PatchPetsResponseBodyTypePetByType PatchPetsResponseBodyType = "PetByType"
)

// PatchPetsResponseBody - Updated
type PatchPetsResponseBody struct {
	PetByAge  *shared.PetByAge
	PetByType *shared.PetByType

	Type PatchPetsResponseBodyType
}

func CreatePatchPetsResponseBodyPetByAge(petByAge shared.PetByAge) PatchPetsResponseBody {
	typ := PatchPetsResponseBodyTypePetByAge

	return PatchPetsResponseBody{
		PetByAge: &petByAge,
		Type:     typ,
	}
}

func CreatePatchPetsResponseBodyPetByType(petByType shared.PetByType) PatchPetsResponseBody {
	typ := PatchPetsResponseBodyTypePetByType

	return PatchPetsResponseBody{
		PetByType: &petByType,
		Type:      typ,
	}
}

func (u *PatchPetsResponseBody) UnmarshalJSON(data []byte) error {

	petByAge := shared.PetByAge{}
	if err := utils.UnmarshalJSON(data, &petByAge, "", true, true); err == nil {
		u.PetByAge = &petByAge
		u.Type = PatchPetsResponseBodyTypePetByAge
		return nil
	}

	petByType := shared.PetByType{}
	if err := utils.UnmarshalJSON(data, &petByType, "", true, true); err == nil {
		u.PetByType = &petByType
		u.Type = PatchPetsResponseBodyTypePetByType
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PatchPetsResponseBody) MarshalJSON() ([]byte, error) {
	if u.PetByAge != nil {
		return utils.MarshalJSON(u.PetByAge, "", true)
	}

	if u.PetByType != nil {
		return utils.MarshalJSON(u.PetByType, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type PatchPetsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updated
	OneOf *PatchPetsResponseBody
}

func (o *PatchPetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchPetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchPetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchPetsResponse) GetOneOf() *PatchPetsResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
