# Security Policy Report for HVXAHV

When the user uses the saved function, the server should not participate in the encryption process of the file, and should only save the hash of the file and the type of the file.

in principle:

All users' passwords and keys should be handled properly to prevent hacker attacks from leaking databases.

The user's password should be stored encrypted.

The server should not store the user's private key.

The issued Token should not be stored on the server, and should use the unique device ID to verify the legitimacy of the client instead of the token.

Save the private key:

About the transfer process of the private key when logging in

The server should not store the user's private key, the private key should only be stored in the user's client.

On first registration, an asymmetric key pair is generated. The logged-in device and device address should be stored in the server, so that the second client sends the address requesting the private key when logging in.

The client saves the private key in the local storage. When the user successfully registers or logs out, it is necessary to send a warning to the user to remind the user that he must back up his private key.

In the case of only one client logging in, if the private key is not saved after exiting, the next login will not be possible.

When a client already exists, when the second or third client logs in, a private key call request will be sent through the existing client URL. To authorize the login request, the user needs to click agree, and then send his private key to the client calling the request.

When the clients call each other's private keys, the server only plays the role of forwarding, and the private key transmission between the two clients is encrypted end-to-end.

When the last client logs out and logs in again, after entering the user name and password, the user needs to manually add his private key to the client before continuing to use it. That is to say, the server will never save the user's private key, and the account can continue to be used only after the private key is added to the client, otherwise the account will be permanently lost.

Although there are some cumbersome processes in the above logic, it is unavoidable.

When logging in, it is necessary to retrieve the number of logged-in devices from the server, and return a list of logged-in devices. When the user selects any client, send the request to the selected client, and then confirm the login in the client that received the request. , start the private key exchange, and send the existing private key to the requesting client

Encryption algorithm for exchanging private keys:

The request and transmission of the key is achieved by communicating with the client end-to-end encrypted. You can exchange your own private key without trusting the server.

Diffie-Hellman key exchange
https://datatracker.ietf.org/doc/html/rfc2631