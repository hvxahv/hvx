package activity

//func (h *Handler) Follow() (*pb.ActivityResponse, error) {
//	var (
//		notok []string
//		ok    []string
//		id    = fmt.Sprintf("%s/%s", h.aAddr, uuid.NewString())
//	)
//
//	//  MARSHAL DATA.
//	body := &activitypub.Follow{
//		Context: "https://www.w3.org/ns/activitystreams",
//		Id:      id,
//		Type:    "Follow",
//		Actor:   h.aAddr,
//		Object:  h.inbox,
//	}
//	marshal, err := json.Marshal(body)
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
//	if err := outbox.NewOutboxes(h.actorId, id, h.inbox, activitypub.FollowType, string(marshal)).Create(); err != nil {
//		return nil, err
//	}
//
//	return response(notok, ok)
//}
