// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/blakes-api-go/pkg/models/shared"
	"github.com/speakeasy-sdks/blakes-api-go/pkg/utils"
	"net/http"
)

type PatchPets200ApplicationJSONType string

const (
	PatchPets200ApplicationJSONTypePetByAge  PatchPets200ApplicationJSONType = "PetByAge"
	PatchPets200ApplicationJSONTypePetByType PatchPets200ApplicationJSONType = "PetByType"
)

type PatchPets200ApplicationJSON struct {
	PetByAge  *shared.PetByAge
	PetByType *shared.PetByType

	Type PatchPets200ApplicationJSONType
}

func CreatePatchPets200ApplicationJSONPetByAge(petByAge shared.PetByAge) PatchPets200ApplicationJSON {
	typ := PatchPets200ApplicationJSONTypePetByAge

	return PatchPets200ApplicationJSON{
		PetByAge: &petByAge,
		Type:     typ,
	}
}

func CreatePatchPets200ApplicationJSONPetByType(petByType shared.PetByType) PatchPets200ApplicationJSON {
	typ := PatchPets200ApplicationJSONTypePetByType

	return PatchPets200ApplicationJSON{
		PetByType: &petByType,
		Type:      typ,
	}
}

func (u *PatchPets200ApplicationJSON) UnmarshalJSON(data []byte) error {

	petByAge := new(shared.PetByAge)
	if err := utils.UnmarshalJSON(data, &petByAge, "", true, true); err == nil {
		u.PetByAge = petByAge
		u.Type = PatchPets200ApplicationJSONTypePetByAge
		return nil
	}

	petByType := new(shared.PetByType)
	if err := utils.UnmarshalJSON(data, &petByType, "", true, true); err == nil {
		u.PetByType = petByType
		u.Type = PatchPets200ApplicationJSONTypePetByType
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PatchPets200ApplicationJSON) MarshalJSON() ([]byte, error) {
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
	PatchPets200ApplicationJSONOneOf *PatchPets200ApplicationJSON
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

func (o *PatchPetsResponse) GetPatchPets200ApplicationJSONOneOf() *PatchPets200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchPets200ApplicationJSONOneOf
}
