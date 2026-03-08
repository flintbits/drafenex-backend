package dto

type CreateOrganizerInput struct {
	FullName    string
	PhoneNumber string

	CompanyName *string
	AvatarUrl   *string
	Website     *string
	Bio         *string

	AddressLine1 string
	AddressLine2 *string
	City         string
	State        string
	Country      string
	PostalCode   string
}
