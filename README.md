# hvxahv

[![Go Report Card](https://goreportcard.com/badge/github.com/hvxahv/hvxahv)](https://goreportcard.com/report/github.com/hvxahv/hvxahv) [![Go](https://github.com/hvxahv/hvxahv/actions/workflows/lint.yml/badge.svg)](https://github.com/hvxahv/hvxahv/actions/workflows/lint.yml)

Instance: [halfmemories.com](https://halfmemories.com)

Preview version: [dev.halfmemories.com](https://dev.halfmemories.com)

## Next plan

- In the process of implementation, the developer believes that the user's private key and some personal sensitive data should not be stored in the server, even if the instance custodian is not allowed to see it, so hvxahv is revamping the design of user privacy.
- The current hvxahv does not represent the final design. The developer is pursuing the best solution for user privacy and security.
- The current hvxahv is under active development and cannot be used as any production environment for the time being. The source code will be frequently modified or rewritten, please be aware.
- hvxahv pursuit of privacy is the ultimate, please believe that the future hvxahv will become safe and reliable.

## Introduction

A completely decentralized social network implementation.

Based on the decentralized [activitypub](https://www.w3.org/TR/activitypub/) decentralized social network protocol, [ipfs](https://ipfs.io/) distributed persistent storage protocol and [matrix.org](https://matrix.org/) real time communication protocol social platform, an open source decentralized social network where the data is really in your hands.

- Find and add your social network friends as easily as an email format (`name@domain`) address.
- Cross-server content publishing channel (channel), the content can not be deleted or modified.
- End-to-end encrypted instant messaging directly with your friends.
- They are all decentralized.

## Build an open source, decentralized open network.

It's a platform for content distribution, social networking, and instant messaging all in one.

Anyone can build their own instance and design their own interface.

You can communicate socially with each other and all platforms that implement the Activitypub protocol.

Store broadcast channel data in the ipfs file system, which is persistent and cannot be deleted or modified.

Communicate with all platforms that implement Matrix protocol with end-to-end encryption.

## Easy to use

Find and add your friends just like email.

### Find friends

You only need to know the ID of the other person, i.e. search for friends like an email format (example: `name@domain`). It is decentralized and you can easily find friends under any instance.

### Socializing

Share your life, interests, and some things between you and your friends.

### Channels

Find public channels across servers.

It provides a decentralized channel feature, you can create or find the broadcast channels you are interested in, the search is as easy as just searching the public index or directly for the keywords of the channels you are interested in, it is IPFS based, once published, the content will not be changed and deleted.

### Instant Messaging

End-to-end encrypted communication directly with your friends. Protect your instant messaging privacy and keep your chat messages away from third parties.

We have implemented a simple end-to-end encrypted chat application based on the matrix protocol, which simplifies the matrix search function, since you can just click on a friend in your address book to start sending messages since the backend already does the rest of the necessary things for you, and you can manually add other chatters from the matrix client, as long as you pass the matrix format `@name:domain` to add friends and talk to them.

## For developers

DEPLOY:

- DEVLOPMENT: [DEVLOPMENT_ENV](build/PRODUCTION_ENV.md)
- PRODUCTION: [PRODUCTION_ENV](build/PRODUCTION_ENV.md)

FOR CLIENT DEVELOPERS: 

- WEB CLIENT REPO: [HVXAHV-WEB](https://github.com/hvxahv/hvxahv-web)

- RESTful API DOCS: [RESTful.md](./app/gateway/RESTful.md)

## Finally

Everyone has the right to control their own data and privacy and to choose who they share it with.

MIT License: [LISENSE](https://github.com/hvxahv/hvxahv/blob/main/LICENSE).

[disism.com](https://disism.com/) OPEN SOURCE. 



🍬 A love letter to the future~

