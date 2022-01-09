# Message

Matrix.org

An open network for secure, decentralized communication.

Matrix Server.

matrix.halfmemories.com

hvxahv will not activate this function by default. When the client implements it, the user needs to click to activate, and then enter the password to indicate that the user agrees to activate this function.

Server:

Build matrix server

[https://github.com/matrix-org/dendrite](https://github.com/matrix-org/dendrite)

In the application, the server side of matrix is ​​still used for account management, as the middleware integrated with hvxahv.

Investigate how to log in

When establishing a dialogue, there will be two ways

1. It is to enter the ID for the dialogue, and the ID should follow the [Matrix.io](http://Matrix.io) standard.
2. Send a message directly in the contact list, there will be two sets for your followers and people you follow.

A dialog box appears when opened, and a button for contacts

In the contact button, people who follow each other and whose MatrixID ≠ false will appear in the contact list.

Chat needs to implement the client itself through the matrix standard on the front end

At present, HVXAHV will only complete the basic interface, such as registration and login