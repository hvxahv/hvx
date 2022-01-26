package handler

type Inboxes struct {
	ID           string `json:"id"`
	From         string `json:"from"`
	FromID       string `json:"from_id"`
	ActivityType string `json:"activity_type"`
	ActivityID   string `json:"activity_id"`
}

//func GetInboxesHandler(c *gin.Context) {
//	a, err := account.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	inboxes, err := activity.NewInboxesAccountID(a.ID).GetInboxes()
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	var ibx []Inboxes
//	for _, i := range *inboxes {
//		from, err := account.NewActorID(i.FromID).GetByActorID()
//		if err != nil {
//			log.Println(err)
//		}
//		inboxes := Inboxes{
//			ID:           strconv.FormatUint(uint64(i.ID), 10),
//			From:         from.Name,
//			FromID:       strconv.FormatUint(uint64(i.FromID), 10),
//			ActivityType: i.ActivityType,
//			ActivityID:   i.ActivityID,
//		}
//		ibx = append(ibx, inboxes)
//	}
//	c.JSON(200, gin.H{
//		"code":    "200",
//		"inboxes": ibx,
//	})
//}
