# Activity

## INBOX:

Event-driven Activitypub inbox (inbox), when the activity event is obtained, it will respond accordingly according to the type of event activity Activity Type, and send the data in the inbox to the client.

- Use the library in gateway for type processing and then send the data to the inbox server.

A message to notify the user

To request attention or chat:

Send to the user's inbox and push the message to the client. The client should implement methods to notify the user such as displaying unread messages.

Store in inboxes.

- Events without client notification (background notification events), such as Create :

Send an event to the client, the user will not get the message in the inbox, but the client will handle the event accordingly when it takes over the event.

It needs to be stored in the corresponding table by type, such as articles or status stored in the articles table.

Push the message to the client, then the client updates the data, but the user is not notified, not stored in the inbox table.