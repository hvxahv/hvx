# Accounts

---

Accounts is the account system of hvxahv. The account system is divided into two data tables.

## Accounts objects

1. Accounts table, for the personal account system of the current instance, used for login, registration and verification.
    
    https://github.com/hvxahv/hvxahv/blob/89ded9bde8b4e00a0499654c43f430de5aca36cb/internal/account/account.go#L12
    
    STRUCT:
    
    1.0. Split the Actor into a second data table.
    
    2.0. Private keys are not stored in the account system.
    
2. Actor https://www.w3.org/TR/activitypub/#actor-objects data sheet for ActivityPub.
    
    https://github.com/hvxahv/hvxahv/blob/89ded9bde8b4e00a0499654c43f430de5aca36cb/internal/account/actor.go#L12
    

## WebFinger

Use Actor information to combine WebFinger [rfc7033 (ietf.org)](https://datatracker.ietf.org/doc/html/rfc7033) to address and open to the outside world.

Example:

````yaml
{
  "subject": "acct:hvturingga@helfmemories.com",
  "links": [
	{
		"rel": "self",
		"type": "application/activity+json",
		"href": "https://helfmemories.com/u/hvturingga"
    }
  ]
}
````

## Actors
Actor not only stores the user's personal data (including public key) of this instance, but also stores the user data of other instances related to the user of this instance, and updates the user data of other instances through the Update event of activitypub.

[https://www.w3.org/TR/activitypub/#actor-objects](https://www.w3.org/TR/activitypub/#actor-objects)

## Registration/Authorization

**register:**

When a user registers, they will query the Accounts table to confirm whether there is a user. Accounts only stores the account information of the instance, the unique user name, and the unique email address.

**Login:**

After the user logs in, a Token will be issued, and the Token will have a unique device ID, which is used to manage the device.

**Actor object storage:**

Actor not only stores the user's personal data (including public key) of this instance, but also stores the user data of other instances related to the user of this instance, and updates the user data of other instances through the Update event of activityPub.

**Inquire:**

When obtaining WebFinger query requests from other instances, it will determine whether there is a user by querying the Username in Accounts.

When an Actor request is obtained, the information in the corresponding Actor table will be queried by querying the ActorID returned by Accounts.

The id of the Actor field is used as the primary key, so the ActorID will be used to query the data when you need to query the user details. For example, when viewing the user's user profile or obtaining a list of follower relationships, ActorID will be used for query.

When updating the user information of the local instance, if you update the username, you need to update the data in two tables, username in Accounts and PreferredUsername in Actors.

**search:**

When searching for users in this example, use the user name query to query all existing users in the actor table (a list of multiple users will be returned). If the `@` symbol exists in the query, it means that the specified user needs to be searched. , if the user data does not exist in the local instance, send an HTTP request to obtain the WebFinger information of other instances, and return all the found users.

Example:
```Go
// If it contains a remote query (judging that the query string contains an @ symbol), it represents an exact query.
func IsLocal(address: string) {
	QUERY Actor WHERE address FIRST
	if Actor == nil {
		HTTP.GET(address)
	}
}
// Querying directly by username returns a collection of actors that already exist in the database.
func FindActor(username string) *[]Actor {
	QUERY Actor WHERE preferred_username FIND 
	return &[]Actor{}
}
```
## User relationship design

pay attention to

followed

mutual concern

Following relationships are not publicly visible

## privacy

Regarding the storage of keys, the public key should be stored on the server, and the private key should not be stored on your own server, please check https://github.com/hvxahv/hvxahv/blob/main/SECURITY.md Privacy Policy Report Get hvxahv's implementation of personal key storage.

## Friendship and status display

Different from Twitter and mastodon, you can only see the status sent by the other party after you follow. The status and articles sent by users will not be publicly displayed to the society. Only after sending a follow request, the other party can view the other party's postings with consent.

---

## set up

- [ ] name
- [ ] username
- [ ] Bio
- [ ] Avatar
- [ ] Private setting Private / Public
- [ ] Change E-Mail
- [ ] Password setting
- [ ] Delete account
- [ ] Push type
    - [ ] Channel update notification
    - [ ] New post notification
    - [ ] Follow updates
    - [ ] follow request
    - [ ] export follow

## Cache:

Store account data in the cache when registering an account.

When registering, check whether the user name of e-mail and username exists, and whether the user's email is used.

Then when registering, updating the data to the redis data will never expire.

You should log in with your email when logging in.

## Block User:

## Delete your account:

Deleting your account will delete it permanently, and there will be no data retention and cannot be recovered.

When someone finds it, it will show that there is no such user, and there will be no prompt that the account has been deleted.

---

## APIs

`SingUp` | Create Account

`Signin` | Login account

`UpdateProfile` | account update

`UpdateUsername` | update username

`ResetPassword` | reset password

`logout` | logout

`DeleteAccount` | delete account

## requires attention

Please take care to prevent sql injection and deduplicate when updating personal data

## Client management

The device unique ID should be saved.

When logging in, write the authenticated device ID into the database as a whitelist, and when logging out (signup) the device, remove this piece of data.

At the time of request, it is judged whether the Token carried by the user exists under the authentication device.

In the account settings, there should be a client management function that displays a list of authorized clients and provides the client offline function.
