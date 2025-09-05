package models

import "time"

// Address represents a user's address information
type Address struct {
	Street     string `json:"street" xml:"street" example:"123 Main St"`
	City       string `json:"city" xml:"city" example:"New York"`
	State      string `json:"state" xml:"state" example:"NY"`
	PostalCode string `json:"postal_code" xml:"postal_code" example:"10001"`
	Country    string `json:"country" xml:"country" example:"USA"`
} // @name Address

// ContactInfo represents a user's contact information
type ContactInfo struct {
	Phone          string `json:"phone" xml:"phone" example:"+1-555-123-4567"`
	AlternateEmail string `json:"alternate_email,omitempty" xml:"alternate_email,omitempty" example:"john.doe.alt@example.com"`
	Website        string `json:"website,omitempty" xml:"website,omitempty" example:"https://johndoe.com"`
} // @name ContactInfo

// Preferences represents a user's preferences
type Preferences struct {
	Theme         string               `json:"theme" xml:"theme" example:"dark"`
	Language      string               `json:"language" xml:"language" example:"en"`
	Notifications NotificationSettings `json:"notifications" xml:"notifications"`
	Privacy       PrivacySettings      `json:"privacy" xml:"privacy"`
} // @name Preferences

// NotificationSettings represents notification preferences
type NotificationSettings struct {
	Email     bool `json:"email" xml:"email" example:"true"`
	SMS       bool `json:"sms" xml:"sms" example:"false"`
	Push      bool `json:"push" xml:"push" example:"true"`
	Marketing bool `json:"marketing" xml:"marketing" example:"false"`
} // @name NotificationSettings

// PrivacySettings represents privacy preferences
type PrivacySettings struct {
	ProfileVisible bool `json:"profile_visible" xml:"profile_visible" example:"true"`
	ShowEmail      bool `json:"show_email" xml:"show_email" example:"false"`
	ShowPhone      bool `json:"show_phone" xml:"show_phone" example:"false"`
} // @name PrivacySettings

// Profile represents additional user profile information
type Profile struct {
	Bio        string   `json:"bio,omitempty" xml:"bio,omitempty" example:"Software engineer with 5 years of experience"`
	Avatar     string   `json:"avatar,omitempty" xml:"avatar,omitempty" example:"https://example.com/avatar.jpg"`
	Skills     []string `json:"skills,omitempty" xml:"skills>skill,omitempty" example:"Go,JavaScript,Python"`
	Experience int      `json:"experience" xml:"experience" example:"5"`
} // @name Profile

// User represents a user in the system with nested structures
type User struct {
	ID          int         `json:"id" xml:"id" example:"1"`
	Name        string      `json:"name" xml:"name" example:"John Doe"`
	Email       string      `json:"email" xml:"email" example:"john@example.com"`
	Address     Address     `json:"address" xml:"address"`
	Contact     ContactInfo `json:"contact" xml:"contact"`
	Profile     Profile     `json:"profile" xml:"profile"`
	Preferences Preferences `json:"preferences" xml:"preferences"`
	CreatedAt   time.Time   `json:"created_at" xml:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt   time.Time   `json:"updated_at" xml:"updated_at" example:"2023-01-01T00:00:00Z"`
} // @name User

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Name        string      `json:"name" xml:"name" binding:"required" example:"John Doe"`
	Email       string      `json:"email" xml:"email" binding:"required,email" example:"john@example.com"`
	Address     Address     `json:"address" xml:"address" binding:"required"`
	Contact     ContactInfo `json:"contact" xml:"contact"`
	Profile     Profile     `json:"profile" xml:"profile"`
	Preferences Preferences `json:"preferences" xml:"preferences"`
} // @name CreateUserRequest

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Name        string       `json:"name,omitempty" xml:"name,omitempty" example:"Jane Doe"`
	Email       string       `json:"email,omitempty" xml:"email,omitempty" example:"jane@example.com"`
	Address     *Address     `json:"address,omitempty" xml:"address,omitempty"`
	Contact     *ContactInfo `json:"contact,omitempty" xml:"contact,omitempty"`
	Profile     *Profile     `json:"profile,omitempty" xml:"profile,omitempty"`
	Preferences *Preferences `json:"preferences,omitempty" xml:"preferences,omitempty"`
} // @name UpdateUserRequest

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" xml:"error" example:"Something went wrong"`
} // @name ErrorResponse

// SuccessResponse represents a success response
type SuccessResponse struct {
	Message string `json:"message" xml:"message" example:"Operation successful"`
} // @name SuccessResponse

// UserListResponse represents a list of users for XML response
type UserListResponse struct {
	Users []User `json:"users" xml:"users>user"`
} // @name UserListResponse
