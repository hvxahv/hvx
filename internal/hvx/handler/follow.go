package handler

//
//func FollowReqHandler(c *gin.Context) {
//	name := middleware.GetUsername(c)
//	id, err := strconv.Atoi(c.PostForm("activity_id"))
//	if err != nil {
//		return
//	}
//
//	privateKey := c.PostForm("private_key")
//
//	f, addr := activity.NewFoAPData(name, uint(id))
//	data, err := json.Marshal(f)
//	if err != nil {
//		return
//	}
//
//	if err := activity.NewAPData(name, addr, privateKey, data).Send(); err != nil {
//		return
//	}
//	c.JSON(200, gin.H{
//		"code":    200,
//		"message": "ok",
//	})
//}
//
//func FollowAcceptHandler(c *gin.Context) {
//	name := middleware.GetUsername(c)
//
//	id := c.PostForm("id")
//	actor := c.PostForm("actor")
//	object := c.PostForm("object")
//
//	aID, err := strconv.Atoi(actor)
//	if err != nil {
//		return
//	}
//
//	oID, err := strconv.Atoi(object)
//	if err != nil {
//		return
//	}
//
//	privateKey := c.PostForm("private_key")
//
//	fa, addr := activity.NewFoAPAccept(name, id, uint(aID))
//	data, err := json.Marshal(fa)
//	if err != nil {
//		return
//	}
//
//	fmt.Println(addr)
//
//	if err := activity.NewAPData(name, addr, privateKey, data).Send(); err != nil {
//		return
//	}
//
//	// objectID (remote) -> actorID (user)
//	if err := activity.NewFollows(uint(oID), uint(aID)).Create(); err != nil {
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"code":    200,
//		"message": "ok",
//	})
//}
//
//func FollowerHandler(c *gin.Context) {
//	name := middleware.GetUsername(c)
//	acct, err := account.NewAccountsUsername(name).GetAccountByUsername()
//	if err != nil {
//		return
//	}
//	fmt.Println(acct.ActorID)
//	f, err := activity.NewByActorID(acct.ActorID).GetFollowers()
//	if err != nil {
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"code":    "200",
//		"message": f,
//	})
//}
//
//func FollowingHandler(c *gin.Context) {
//	name := middleware.GetUsername(c)
//	acct, err := account.NewAccountsUsername(name).GetAccountByUsername()
//	if err != nil {
//		return
//	}
//	fmt.Println(acct.ActorID)
//	f, err := activity.NewByActorID(acct.ActorID).GetFollowing()
//	if err != nil {
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"code":    "200",
//		"message": f,
//	})
//}
