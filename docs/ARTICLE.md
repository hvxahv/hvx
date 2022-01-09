# Article / Status

The article function is used for users to create an article or status, the published article or status is only visible to your friends (people who follow each other), it uses the activitypub protocol, you can delete or modify it, and your friends will be notified after publishing.

---

The sent type is status, then statuses should be set to true. NSFW content will also use the boolean type.

The comment function of the article is set as a conversation, which is saved in the conversations table. A visibility data table interface needs to be maintained.

How should the editor interface be designed?

It is designed to switch between two editors, one is Article text priority, which is consistent with the rich text editor, including articles and illustrations.

The other is the status option, which can post a short text or image-first status.

To pay attention to article and state toggles, visibility, NSFW, you need to add a checkbox like this.

You will have a ticked option when you publish: Sync to channel (your channel and the channel you manage). For sending a public post or status, you can use this, any subscriber to your channel can see it, although you can delete it in your or followers timeline, however, because the content posted in the channel will be Using the ipfs protocol, this means that your published content will be permanently stored and unmodifiable, and it will be synchronized to other ipfs nodes.

For channel details, please check [Channel](https://www.notion.so/Channel-b403edc9892a4c2598bcd2cda7d187aa). GITHUB LINK: [Channels](https://github.com/disism/hvxahv/blob/main/app/channel /README.md) .

How will content attachments be saved?

The attachment will be stored in the current instance's server using object storage, it should not be stored in ipfs since it can be deleted and changed. Note: Pictures are also attachments to the content!

But when you sync to the channel, the attachment will be stored in IPFS.

Other ideas: replace likes with reads

There is no like function, but consider adding the function of reading volume (👁 number), but keep the comment function, and both you and your friends will see the content of the comment.

### Articles structure
https://github.com/hvxahv/hvxahv/blob/e7f2bbb3aa7e249ea49c66fc816cd395d12c6711/internal/article/article.go#L36

### Interface
https://github.com/hvxahv/hvxahv/blob/e7f2bbb3aa7e249ea49c66fc816cd395d12c6711/internal/article/article.go#L150