package models

import "time"

type OrganizerStatus string

const (
	OrganizerPending   OrganizerStatus = "PENDING"
	OrganizerActive    OrganizerStatus = "ACTIVE"
	OrganizerSuspended OrganizerStatus = "SUSPENDED"
)

// func (s OrganizerStatus) IsValid() bool {
// 	switch s {
// 	case OrganizerPending, OrganizerActive, OrganizerSuspended:
// 		return true
// 	default:
// 		return false
// 	}
// }

type Organizer struct {
	ID          int64  `json:"id" db:"id"`
	UserID      int64  `json:"user_id" db:"user_id"`
	FullName    string `json:"full_name" db:"full_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`

	CompanyName *string `json:"company_name,omitempty" db:"company_name"`
	AvatarUrl   *string `json:"avatar_url,omitempty" db:"avatar_url"`
	Website     *string `json:"website,omitempty" db:"website"`
	Bio         *string `json:"bio,omitempty" db:"bio"`

	AddressLine1 string  `json:"address_line1" db:"address_line1"`
	AddressLine2 *string `json:"address_line2,omitempty" db:"address_line2"`
	City         string  `json:"city" db:"city"`
	State        string  `json:"state" db:"state"`
	Country      string  `json:"country" db:"country"`
	PostalCode   string  `json:"postal_code" db:"postal_code"`

	Status     OrganizerStatus `json:"status" db:"status"`
	IsVerified bool            `json:"is_verified" db:"is_verified"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

/*
TODO: call /organizer/me
*/

/*
ID          int64     `json:"id" db:"id"`
UserID      int64     `json:"user_id" db:"user_id"`
Email       string    `json:"email" db:"email"`
FullName    string    `json:"full_name" db:"full_name"`
PhoneNumber string    `json:"phone_number" db:"phone_number"`
Status      string    `json:"status" db:"status"`
IsVerified  bool      `json:"is_verified" db:"is_verified"`


CompanyName string `json:"company_name" db:"company_name"`
AvatarURL   string `json:"avatar_url" db:"avatar_url"`
Website     string `json:"website" db:"website"`
Bio         string `json:"bio" db:"bio"`

AddressLine1 string `json:"address_line1" db:"address_line1"`
AddressLine2 string `json:"address_line2" db:"address_line2"`
City         string `json:"city" db:"city"`
State        string `json:"state" db:"state"`
Country      string `json:"country" db:"country"`
PostalCode   string `json:"postal_code" db:"postal_code"`

KYCStatus        string     `json:"kyc_status" db:"kyc_status"`
KYCVerifiedAt    *time.Time `json:"kyc_verified_at" db:"kyc_verified_at"`
IdentityDocURL   string     `json:"identity_doc_url" db:"identity_doc_url"`
TaxID            string     `json:"tax_id" db:"tax_id"`


MaxActiveAuctions int  `json:"max_active_auctions" db:"max_active_auctions"`
CanCreateAuction  bool `json:"can_create_auction" db:"can_create_auction"`
IsSuspended       bool `json:"is_suspended" db:"is_suspended"`
*/
