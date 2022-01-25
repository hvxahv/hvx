package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"google.golang.org/grpc"
	"log"
)

func CreateAccountHandler(c *gin.Context) {
	conn, err := grpc.Dial("localhost:7041", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	cli := pb.NewAccountsClient(conn)

	username := c.PostForm("username")
	password := c.PostForm("password")
	mail := c.PostForm("mail")

	// https://datatracker.ietf.org/doc/html/rfc5208
	publicKey := c.PostForm("public_key")

	d := &pb.CreateData{
		Username:  username,
		Mail:      mail,
		Password:  password,
		PublicKey: publicKey,
	}
	create, err := cli.Create(c, d)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code":    create.Code,
		"message": create.Reply,
	})

}

//func SignInHandler(c *gin.Context) {
//	ua := c.GetHeader("User-Agent")
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//
//	id, mail, err := account.NewAuth(username, password).SignIn()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	deviceID := uuid.New().String()
//	token, err := policy.GenToken(mail, username, password, deviceID)
//	if err != nil {
//		log.Println(err)
//	}
//	d := device.NewDevices(id, ua, deviceID)
//	if err := d.Create(); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"code":      "200",
//		"token":     token,
//		"mail":      mail,
//		"deviceID":  deviceID,
//		"publicKey": d.PublicKey,
//	})
//}

//func GetPublicKeyHandlers(c *gin.Context) {
//	name := middleware.GetUsername(c)
//	actor, err := account.NewActorByAccountUsername(name).GetActorByAccountUsername()
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	c.JSON(200, gin.H{
//		"code":       "200",
//		"public_key": actor.PublicKey,
//	})
//}

//type ECDHDerive struct {
//	DeviceID    string `json:"device_id"`
//	ReqDeviceID string `json:"req_device_id"`
//	JWK         string `json:"jwk"`
//	IV          string `json:"iv"`
//}
//
//func GetDHPublicJWKHandlers(c *gin.Context) {
//	deviceID := c.Param("id")
//	data, err := cache.GETDHData(deviceID)
//	if err != nil {
//		log.Println(data)
//		return
//	}
//
//	log.Println(deviceID)
//
//	var e ECDHDerive
//	dec := gob.NewDecoder(bytes.NewReader(data))
//	if err := dec.Decode(&e); err != nil {
//		log.Fatal("decode error:", err)
//	}
//	c.JSON(200, gin.H{
//		"code":          "200",
//		"req_device_id": e.ReqDeviceID,
//		"device_id":     e.DeviceID,
//		"jwk":           e.JWK,
//		"iv":            e.IV,
//	})
//}
//
//func RequestPrivateKeyHandlers(c *gin.Context) {
//	a, err := account.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	// Owner device ID.
//	deviceID := c.PostForm("device_id")
//	req := c.PostForm("req_device_id")
//
//	log.Println(a)
//	jwk := c.PostForm("jwk")
//	iv := c.PostForm("iv")
//
//	var buf bytes.Buffer
//	enc := gob.NewEncoder(&buf)
//	if err := enc.Encode(ECDHDerive{
//		ReqDeviceID: req,
//		DeviceID:    deviceID,
//		JWK:         jwk,
//		IV:          iv,
//	}); err != nil {
//		log.Fatal("encode error:", err)
//	}
//	fmt.Println(buf.Bytes())
//	if err := cache.SETDHData(req, buf.Bytes()); err != nil {
//		log.Println(err)
//		return
//	}
//
//	log.Println(req)
//
//	// Get device id by req hash and sent to the requested device via a notification.
//	hash, err := device.NewDeviceByHash(a.ID, req).GetDeviceByHash()
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	d, err := json.Marshal(push.NewData(
//		"Notify",
//		fmt.Sprintf("You are preparing to login on another device: %s.", deviceID),
//		"https://avatars.githubusercontent.com/u/94792300?s=200&v=4",
//		"Authorized"),
//	)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	if err := notify.NewPush(a.ID, hash.ID, d).Push(); err != nil {
//		log.Println(err)
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"code":    "200",
//		"message": "The request is successful, please confirm login in the requested terminal.",
//	})
//
//}
//
//type ECDHRespData struct {
//	DeviceID   string `json:"device_id"`
//	PublicJWK  string `json:"public_jwk"`
//	PrivateJWK string `json:"private_jwk"`
//}
//
//func SendPrivateKeyHandlers(c *gin.Context) {
//	deviceID := c.PostForm("device_id")
//	publicKey := c.PostForm("public_jwk")
//	privateKey := c.PostForm("private_jwk")
//	var buf bytes.Buffer
//	enc := gob.NewEncoder(&buf)
//	if err := enc.Encode(ECDHRespData{
//		DeviceID:   deviceID,
//		PublicJWK:  publicKey,
//		PrivateJWK: privateKey,
//	}); err != nil {
//		log.Fatal("encode error:", err)
//	}
//	fmt.Println(buf.Bytes())
//	if err := cache.SETDHData(deviceID, buf.Bytes()); err != nil {
//		log.Println(err)
//		return
//	}
//	c.JSON(200, gin.H{
//		"code":    "200",
//		"message": "Private key sent successfully.",
//	})
//}

//func GetDHPrivateJWKHandlers(c *gin.Context) {
//	deviceID := c.Param("id")
//	data, err := cache.GETDHData(deviceID)
//	if err != nil {
//		log.Println(data)
//		return
//	}
//
//	name := middleware.GetUsername(c)
//	actor, err := account.NewActorByAccountUsername(name).GetActorByAccountUsername()
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	var e ECDHRespData
//	dec := gob.NewDecoder(bytes.NewReader(data))
//	if err := dec.Decode(&e); err != nil {
//		log.Fatal("decode error:", err)
//	}
//	c.JSON(200, gin.H{
//		"code":        "200",
//		"public_key":  actor.PublicKey,
//		"device_id":   e.DeviceID,
//		"public_jwk":  e.PublicJWK,
//		"private_jwk": e.PrivateJWK,
//	})
//}

//func LogoutHandler(c *gin.Context) {
//	acct, err := account.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	if err := device.NewDeviceByHash(acct.ID, middleware.GetDevicesID(c)).Delete(); err != nil {
//		fmt.Println(err)
//		return
//	}
//	c.JSON(200, gin.H{
//		"code":    "200",
//		"message": "logout",
//	})
//}
//
//func GetDevicesHandler(c *gin.Context) {
//	username := middleware.GetUsername(c)
//	acct, err := account.NewAccountsUsername(username).GetAccountByUsername()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	devices, err := device.NewDevicesByAccountID(acct.ID).GetDevicesByAccountID()
//	if err != nil {
//		return
//	}
//	c.JSON(200, gin.H{
//		"code":    "200",
//		"devices": devices,
//	})
//}
//
//func DeleteDevicesHandler(c *gin.Context) {
//	acct, err := account.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	if err := device.NewDeviceByHash(acct.ID, c.PostForm("device_hash")).Delete(); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"code":    "200",
//		"message": "ok!",
//	})
//}
