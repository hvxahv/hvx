# Channel

Channel is an extension of ActivityPub. When subscribing to a channel, ActivityPub is called follow and in hvxahv it is called subscription.

The type of Channel is a Service Details in Activitypub:

[https://www.w3.org/TR/activitystreams-vocabulary/#actor-types](https://www.w3.org/TR/activitystreams-vocabulary/#actor-types)

````go
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Service",
  "name": "Acme Web Service"
}
````

It is an IPFS based content publishing function.

## Create and subscribe/follow:

On creation, you need to use NewActor to create an ActivityPub Actor of type Service . Because Channel is based on Activitypub, it should also be saved to Actor as a service account of Activitypub for subscribers to follow or subscribe to.

Note: When a subscriber sends a subscription request, it will set the automatic reply Activitypub Accept to let the subscriber complete the subscription, which is different from the individual user's attention logic.

Create Channel

````go
if err := NewActor(); err != nil {
    return err
}

if err := NewChannel().New(); err != nil {
    return err
}
````

Subscription / Follow Channel

````go
if err := NewSubscript().New(); err != nil {
    return err
}

func (s *Sub) New() {
    stream.Request("Accept")
}

````

## Find channel:

It is the same as Activitypub's Actor, but the url is not consistent, so as to distinguish, for example: `https://domain.com/channel/<c>`. Use the link directly to access or follow the home page of the service.

The public Channel s that exist in the instance can be found by searching.