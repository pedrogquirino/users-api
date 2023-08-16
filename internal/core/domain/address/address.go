package address

type Address struct {
	Id       string
	District string
	Street   string
	Number   string
	ZipCode  string
}

type AddressRequest struct {
	District string `json:"district" validate:"required"`
	Street   string `json:"street" validate:"required"`
	Number   string `json:"number" validate:"required"`
	ZipCode  string `json:"zipcode" validate:"required"`
}

type AddressResponse struct {
	Id       string `json:"id"`
	District string `json:"district"`
	Street   string `json:"street"`
	Number   string `json:"number"`
	ZipCode  string `json:"zipcode"`
}

func ConvertAddressRequestToAddress(request AddressRequest) *Address {
	return &Address{
		District: request.District,
		Street:   request.Street,
		Number:   request.Number,
		ZipCode:  request.ZipCode,
	}
}
