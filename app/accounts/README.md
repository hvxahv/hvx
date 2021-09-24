# Accounts

About WebFinger production standards

[https://datatracker.ietf.org/doc/html/rfc7033](https://datatracker.ietf.org/doc/html/rfc7033)

About http signatures production standard

[https://datatracker.ietf.org/doc/html/draft-cavage-http-signatures-10](https://datatracker.ietf.org/doc/html/draft-cavage-http-signatures-10)

---

Accounts is divided into two tables, the Accounts data table of the internal personal account system, and the actor data table for the activitypub protocol.

Notes:

- When a user registers, he will query the Accounts table to confirm whether there is a user. Accounts only stores the account information of this instance, including the unique user name, the unique email address, and the private key password of the account.
- Actor not only stores the user's personal data (including the public key) of this instance, but also stores the user data of other instances related to the user of this instance, and updates the user data of other instances through the Update event of activitypub.
- When obtaining the WebFinger query request of other instances, it will determine whether there is a user by querying the Username in Accounts, and then when obtaining the Actor request, it will query the information in the corresponding Actor table by querying the ActorID returned by Accounts.
- When searching for users in this instance, use username query to query all existing users in the actor table (fuzzy query). If the `@` symbol is present in the query, it means that the query is for information about other instances, use remote Query and send http request to obtain WebFinger information of other instances, so that all the users found will be returned.
- Use the id of the actor as the unique identifier of the user. When it is associated with other tables, use the ActorID to query data. For example, when searching the list of articles published by the user, or obtaining the user's following relationship, the ActorID will be used as the primary key for query.
- When updating user information, if you update the user name, you need to update the data in two tables, username in Accounts and PreferredUsername in Actors.
- When setting user information, it is divided into privacy setting and personal information setting.