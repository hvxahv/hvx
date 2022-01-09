# Accounts

Accounts is the account service of this project, which is used to manage accounts.

About WebFinger Production Standards

[https://datatracker.ietf.org/doc/html/rfc7033](https://datatracker.ietf.org/doc/html/rfc7033)

About the HTTP Signatures Production Standard

[https://datatracker.ietf.org/doc/html/draft-cavage-http-signatures-10](https://datatracker.ietf.org/doc/html/draft-cavage-http-signatures-10)

## Account system design:

The account system is divided into two sets of account systems.
1. Accounts data, for the internal personal account system, used for internal verification such as login and registration of the current instance.
    
     2.0. Modify the structure → do not store private keys in the account system.

```Go
type Accounts struct {
	gorm.Model

	Username string `gorm:"primaryKey;type:text;preferredUsername;" validate:"required,min=4,max=16"`
	Mail     string `gorm:"index;type:text;mail;unique" validate:"required,email"`
	Password string `gorm:"type:text;password" validate:"required,min=8,max=100"`

	// When creating an account, first verify the username, email address, and password.
	// After the verification is successful, store the username and key in the actors table,
	// then use the returned ActorID in this field, and then store the data in the account table .
	// At this time, the context of creating the user is complete.
	ActorID uint `gorm:"type:bigint;actor_id"`

	// Whether to set as a private account
	IsPrivate bool `gorm:"type:boolean;is_private"`
}
```
2. Actor data, for Activitypub's actor information system, used when interacting with activitypub instances.
```Go
type Actors struct {
	gorm.Model

	PreferredUsername string `gorm:"primaryKey;type:text;preferredUsername;"`
	Domain            string `gorm:"index;type:text;domain"`
	Avatar            string `gorm:"type:text;avatar"`
	Name              string `gorm:"type:text;name"`
	Summary           string `gorm:"type:text;summary"`
	Inbox             string `gorm:"type:text;inbox"`
	Url               string `gorm:"index;test;url"`
	PublicKey         string `gorm:"type:text;public_key"`

	// Whether it is a robot or other type of account
	ActorType string `gorm:"type:text;actor_type"`

	// Set whether it is a remote actor.
	IsRemote bool `gorm:"type:boolean;is_remote"`
}
```

**register:**

When a user registers, they will query the Accounts table to confirm whether there is a user. Accounts only stores the account information of the instance, the unique user name, and the unique email address.

**storage:**

Actor not only stores the user's personal data (including public key) of this instance, but also stores the user data of other instances related to the user of this instance, and updates the user data of other instances through the Update event of activityPub.

**Inquire:**

When obtaining the WebFinger query request of other instances, it will check whether there is a user by querying the Username in Accounts, and then when obtaining the Actor request, it will query the information in the corresponding Actor table by querying the ActorID returned by Accounts.

## search

When searching for users in this instance, use the user name query to query all existing users in the actor table (fuzzy query). Sending an http request to get WebFinger information for other instances will return all users found.

example:
```Go
// If it contains a remote query (judging that the query string contains an @ symbol), it represents an exact query.
if remote {
	http.request(hostname)
}

// Querying directly by username returns a collection of actors that already exist in the database.
func FindActor(username string) {
	QUERY ACCOUNTS WHERE preferred_username FIND 
	return []actor{}
}
```
The actor's id is used as the user's unique identifier, so use ActorID to query data when querying. For example, when searching for the list of articles published by the user, or obtaining the user's following relationship, ActorID will be used as the primary key for query.

- When updating user information, if you update the username, you need to update the data in two tables, username in Accounts and PreferredUsername in Actors.
- When setting user information, it is divided into privacy settings and personal information settings.


## privacy

Regarding the storage of keys, the public key should be stored on the server, and the private key should not be stored on its own server, but there is no better way to store the private key, so it can only be stored in the database of the running instance.

## Friendship and status display

Only friends who follow each other can see the status sent by each other

Notes:

- When a user registers, he will query the Accounts table to confirm whether there is a user. Accounts only stores the account information of this instance, including the unique user name, the unique email address, and the private key password of the account.
- Actor not only stores the user's personal data (including the public key) of this instance, but also stores the user data of other instances related to the user of this instance, and updates the user data of other instances through the Update event of activitypub.
- When obtaining the WebFinger query request of other instances, it will determine whether there is a user by querying the Username in Accounts, and then when obtaining the Actor request, it will query the information in the corresponding Actor table by querying the ActorID returned by Accounts.
- When searching for users in this instance, use username query to query all existing users in the actor table (fuzzy query). If the `@` symbol is present in the query, it means that the query is for information about other instances, use remote Query and send http request to obtain WebFinger information of other instances, so that all the users found will be returned.
- Use the id of the actor as the unique identifier of the user. When it is associated with other tables, use the ActorID to query data. For example, when searching the list of articles published by the user, or obtaining the user's following relationship, the ActorID will be used as the primary key for query.
- When updating user information, if you update the user name, you need to update the data in two tables, username in Accounts and PreferredUsername in Actors.
- When setting user information, it is divided into privacy setting and personal information setting.