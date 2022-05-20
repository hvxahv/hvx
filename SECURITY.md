# Security Policy Report for HVXAHV

HVX is built for privacy and we take security very seriously. If you have any security questions, please contact: [hvxahv.security@disism.com](hvxahv.security@disism.com)

## HVXAHV Security Policy Report

When the user uses the saved function, the server should not participate in the encryption process of the file, and should only save the hash of the file and the type of the file.

### in principle

All users' passwords and keys should be handled properly to prevent hacker attacks from leaking databases.

The user's password should be stored encrypted.

The server should not store the user's private key.

The issued Token should not be stored in the server, and should use the unique device ID (device hash) to verify the legitimacy of the client.

Saved feature personal files stored in IPFS should be stored encrypted.

## TOKEN

### Authorization Key generation

### private key

### Private key storage and exchange process

Regarding the transmission process of the private key when logging in, the server should not store the user's private key, and the private key should only be stored in the user's client.

On first registration, an asymmetric key pair is generated. The logged-in device and device address should be stored in the server, so that the second client sends the address requesting the private key when logging in.

The client saves the private key in the local storage. When the user successfully registers or logs out, it is necessary to send a warning to the user to remind the user that he must back up his private key.

In the case of only one client logging in, if the private key is not saved after exiting, the next login will not be possible.

When there is already one client, when the second or third client logs in, a private key call request will be sent through the existing client, and then an authorization login request needs to be made on the specified client , when you click agree, send your own private key to the client requesting the call.

When the clients call each other's private keys, the server only plays the role of forwarding, and the private key transmission between the two clients is encrypted end-to-end.

When the last client logs out and logs in again, after entering the user name and password, the user needs to manually add his private key to the client before continuing to use it. That is to say, the server will never save the user's private key, and the account can continue to be used only after the private key is added to the client, otherwise the account will be permanently lost.

Although there are some cumbersome processes in the above logic, it is unavoidable.

When logging in, it is necessary to retrieve the number of logged-in devices from the server, and return a list of logged-in devices. When the user selects any client, send the request to the selected client, and then confirm the login in the client that received the request. , start the private key exchange, and send the existing private key to the requesting client.

### Encryption algorithm for exchanging private keys

The request and transmission of the key is achieved by communicating with the client end-to-end encrypted. You can exchange your own private key without trusting the server.

In Diffie-Hellman key exchange (ECDH), a new RSA key is regenerated for each exchange. The data during the exchange is set in the cache for two minutes, and is discarded immediately after the exchange. The private keys of both parties will not be stored, and only the public keys can be intercepted, so a certain degree of security is guaranteed.

## file encryption

###Saved

Regarding the encryption of Saved files, all saved files are stored in IPFS. As we all know, IPFS does not encrypt files. When other people know the hash, they can access this file. Therefore, when uploading the file, it needs to be encrypted and uploaded, and then it will return The cid of the file is presented to the client, which decrypts the file upon download. The current setting of hvxahv-web is to use openpgp for asymmetric encryption of the file, and when downloading, it will use the client's local private key for decryption.