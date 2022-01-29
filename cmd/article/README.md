# Articles

The article function is used for users to create an article or status. The published article or status is only visible to your friends (people who follow each other). It uses the activitypub protocol. You can delete or modify it, and your friends will be notified after publishing.

You can also check the channel (your channel and the channel you manage) when you publish

If you want to send a public article or status, you can publish it to the channel, any subscribers of your channel can see it, and you can delete it. However, the content published in the channel will use the ipfs protocol, which means you publish The content of will be stored permanently and cannot be modified, and it will be synchronized to other ipfs nodes.

For details, please refer to Channel's [README](../channel/README.md).

---

How will the images of the content be saved?

The picture will be saved in the server of the current instance. Because it can be deleted and changed, it should not be stored in IPFS.

How to save shared files?

The shared file needs to be uploaded to the IPFS node by yourself. At this time, the file can be encrypted and saved, just like a network disk. After saving, it will be shared with other users.