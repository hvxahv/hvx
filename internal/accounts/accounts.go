package accounts

import "gorm.io/gorm"

type Accounts struct {
	gorm.Model

	Username string `gorm:"primaryKey;type:text;preferredUsername;" validate:"required,min=4,max=16"`
	Mail     string `gorm:"index;type:text;mail;unique" validate:"required,email"`
	Password string `gorm:"type:text;password" validate:"required,min=8,max=100"`

	// When creating an account, first verify the username, email address, and password.
	// After the verification is successful, store the username and key in the actors table.
	// Then use the returned ActorID in this field, and then store the data in the accounts table.
	// At this time, the contextual interaction of creating the user is complete.
	ActorID int `gorm:"type:bigint;actor_id"`

	// Whether it is a robot or other type of account.
	AccountType string `gorm:"type:text;account_type"`

	// Whether to set as a private account.
	IsPrivate  bool   `gorm:"type:boolean;is_private"`
	PrivateKey string `gorm:"type:text;private_key"`
}

type Actors struct {
	gorm.Model

	PreferredUsername string `gorm:"type:text;preferredUsername;"`
	Domain            string `gorm:"type:text;doamin"`

	Avatar  string `gorm:"type:text;avatar"`
	Name    string `gorm:"type:text;name"`
	Summary string `gorm:"type:text;summary"`

	// The id of the matrix server.
	MatrixID string `gorm:"type:text;matrix_id;unique"`

	// inbox address.
	Inbox     string `gorm:"type:text;inbox"`
	PublicKey string `gorm:"type:text;public_key"`
}
