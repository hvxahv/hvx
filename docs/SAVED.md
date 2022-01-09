# Saved

The client needs to use RSA PublicKey to encrypt and upload files to the IPFS network.

The encryption process should be implemented by the client and encrypted by the public key of the individual user. The client must also decrypt the file and display the decrypted file to the user.

TODO - Retrofit the saved encryption method.

In the current saved client, only the CID returned by IPFS is encrypted, and the file will not be encrypted and saved. In the future, we will explore how the client can quickly encrypt files and upload them to IPFS.