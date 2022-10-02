package activity

//
//func (h *Handler) Accept(data []byte) (*pb.ActivityResponse, error) {
//	var (
//		notok []string
//		ok    []string
//		id    = fmt.Sprintf("%s/%s", h.aAddr, uuid.NewString())
//	)
//
//	var b Object
//	if err := json.Unmarshal(data, &b); err != nil {
//		return nil, err
//	}
//
//	//  MARSHAL DATA.
//	marshal, err := json.Marshal(&activitypub.Accept{
//		Context: "https://www.w3.org/ns/activitystreams",
//		Id:      id,
//		Type:    activitypub.AcceptType,
//		Actor:   h.aAddr,
//		Object: struct {
//			Id     string `json:"id"`
//			Type   string `json:"type"`
//			Actor  string `json:"actor"`
//			Object string `json:"object"`
//		}{
//			Id:     b.Id,
//			Type:   b.Type,
//			Actor:  b.Actor,
//			Object: b.Object,
//		},
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	// DELIVERY ...
//	do, err := NewDelivery(marshal, h.aAddr, h.privateKey).Do(fmt.Sprintf("%s/inbox", h.inbox))
//	if err != nil {
//		return nil, err
//	}
//	if do.StatusCode != 202 {
//		notok = append(notok, h.inbox)
//		return nil, nil
//	}
//	ok = append(ok, h.inbox)
//
//	// CREATE FOLLOW OUTBOX ...
//	if err := outbox.NewOutboxes(h.actorId, id, h.inbox, activitypub.AcceptType, string(marshal)).Create(); err != nil {
//		return nil, err
//	}
//
//	switch b.Type {
//	case activitypub.FollowType:
//		object, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActorByAddress(h.inbox)
//		if err != nil {
//			return nil, err
//		}
//
//		// IF ACCEPT FOLLOW REQUEST, ADD FOLLOWER.
//		if err := friendship.NewFollower(h.actorId, uint(object.Id)).Follow(); err != nil {
//			return nil, err
//		}
//	default:
//
//	}
//	return response(notok, ok)
//}
