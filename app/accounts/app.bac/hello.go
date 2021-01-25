package app_bac

import (
	"github.com/gin-gonic/gin"
	rand2 "math/rand"
	"strconv"
)

func LikeHandler(c *gin.Context)  {
	privKey := []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAzcU2iJkBv6GpgU02t9/6w1ZrKVhlw/iBLN5RmsC7K1jVLEPG
Cc0XWGg26ab8ZomKQfq2DpjS25L1aBjfxkWWnJ13PnLbCTkoMiVGinEFpwuDcONN
93YwZuPTqzrwZCbPc5L2owlI30AVXfl3zacmSgwXHISyJ3Lri5RUbUv1EMMYcRxw
PxPiD1AFZFTfCoB32Tv8vr4zqUJNeh+rnU8BVZfgwLHWwjWazVxTKdirQz8HoAra
ARs4pxpI0Y74FSnMANfnOuwevstV+14hFtObWuCJurVuZkBga9pm2gsFCAdyifaN
oXZMojsKSh0+d6ENuWx0247PxbQjpSNZ+atvAwIDAQABAoIBAE4qDwgejzJ7N4d9
jD6W34vaRQTEpoul13Py067kbXUO3pNc/1pzxN15YPqqyxAsSQ+9K/EpjleIYJ51
bHPFtWuyyTw4pwm+440A6eXm/w2wavKz9cb9527+NkbCMdTphT5S4kuWWE8t6a/z
9tVxHDNemlzj70Ta+6ze4J9QDjyaIMgqsrDKi57IkKPMtuXIDs+hkN0YSwVrveYb
mHkkWANZRhf9sZgAVoKFAPAPONWcEk6/5fdKXURzL9hKd12i8QN91raF2kdSpq26
o8QLmoUYcm69sbq0RTW/asHBUl9W78yVFAqcKFf3veurQXV+RUDdXoyFVZVYOU8V
APH4AoECgYEA9+ff1a5E8V0XzzxEOR5pcq+wzr5+Va+PX6PlfKWg/VcuQW7igqQG
NWeWTL46LjtJ8sjsan4a5CVkrUQVIboMjmaXUvxUQdZFurt93YHb4RYOIJ7As4ub
ngSjzljB607dFbm7FrwPaQpVfu+Jr9NSboHXp0qvGHzbmmAfWTE96sMCgYEA1H0m
KDOG8YFSM/5H4dubPt/uxw84XpRz/h3dfB1b9uVr5pGuUuW4iahU0rH+37M0SwOU
zu1wboPILRCtCnW0EVl631QhdDvgzttgCa41iLedDlPIOCvUUezrTAmWomkzQY7A
DKnmvYtUaJeoH0xLy/KVvqqs9QHvWQ8RGoMBpsECgYBeElJA/qlQwde4HxFwb1bq
WzsHqQe9sioAy+vYee12VS/LSq2Pj3h/qXhCVOk1A/CSvaXg30uLCZmDmlM+giAj
pVCFABVlCk9Zha5EcPOkHT9tJ3DOcQqzUuVp9wpXJY0Fon4ZKXHblT1ONs2Tt+gO
4RyZceCxtSeMFC+xI9twCwKBgHIu4h8+OShz1smLvfLXgKAT04ryWdcxLPeD8u3s
FOwiso3PAP8Y5MZMR3CFJ7Hr6ZDZ1tAvdXhdpmbZDDOPtniQPd/epK+CMbbW2c+5
5piWGnaFfRT6MHjpuDM+/8w8fcefvwHPFugKBAzEWhqfdCefLuqrao+qP4T6/LTj
azLBAoGAOunZ7iWZsyhz/NiGd1Chq5RVZHqISd3RPEDhAQutdww22G9yor9dTWJG
FRM8yDghD7Gk6qlGuqjM3pOwaEaOgCFovKzyWhH9PT6aQsUUHCAIJjQRHAQt4o9S
QeEpnzPydEqbmGwd1VBSlBaEHgPRenDtC03sIwg1v/DeNeZ+QSk=
-----END RSA PRIVATE KEY-----`)

	da := strconv.Itoa(rand2.Int())
	//k := RsaSignWithSha256(privKey)

	//header := `keyId="https://activity.disism.com/actor",headers=(request-target) host date",signature="` + k
	LikeReq(da, privKey)
}
func HandleHello2(c *gin.Context) {
	privKey := []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAzcU2iJkBv6GpgU02t9/6w1ZrKVhlw/iBLN5RmsC7K1jVLEPG
Cc0XWGg26ab8ZomKQfq2DpjS25L1aBjfxkWWnJ13PnLbCTkoMiVGinEFpwuDcONN
93YwZuPTqzrwZCbPc5L2owlI30AVXfl3zacmSgwXHISyJ3Lri5RUbUv1EMMYcRxw
PxPiD1AFZFTfCoB32Tv8vr4zqUJNeh+rnU8BVZfgwLHWwjWazVxTKdirQz8HoAra
ARs4pxpI0Y74FSnMANfnOuwevstV+14hFtObWuCJurVuZkBga9pm2gsFCAdyifaN
oXZMojsKSh0+d6ENuWx0247PxbQjpSNZ+atvAwIDAQABAoIBAE4qDwgejzJ7N4d9
jD6W34vaRQTEpoul13Py067kbXUO3pNc/1pzxN15YPqqyxAsSQ+9K/EpjleIYJ51
bHPFtWuyyTw4pwm+440A6eXm/w2wavKz9cb9527+NkbCMdTphT5S4kuWWE8t6a/z
9tVxHDNemlzj70Ta+6ze4J9QDjyaIMgqsrDKi57IkKPMtuXIDs+hkN0YSwVrveYb
mHkkWANZRhf9sZgAVoKFAPAPONWcEk6/5fdKXURzL9hKd12i8QN91raF2kdSpq26
o8QLmoUYcm69sbq0RTW/asHBUl9W78yVFAqcKFf3veurQXV+RUDdXoyFVZVYOU8V
APH4AoECgYEA9+ff1a5E8V0XzzxEOR5pcq+wzr5+Va+PX6PlfKWg/VcuQW7igqQG
NWeWTL46LjtJ8sjsan4a5CVkrUQVIboMjmaXUvxUQdZFurt93YHb4RYOIJ7As4ub
ngSjzljB607dFbm7FrwPaQpVfu+Jr9NSboHXp0qvGHzbmmAfWTE96sMCgYEA1H0m
KDOG8YFSM/5H4dubPt/uxw84XpRz/h3dfB1b9uVr5pGuUuW4iahU0rH+37M0SwOU
zu1wboPILRCtCnW0EVl631QhdDvgzttgCa41iLedDlPIOCvUUezrTAmWomkzQY7A
DKnmvYtUaJeoH0xLy/KVvqqs9QHvWQ8RGoMBpsECgYBeElJA/qlQwde4HxFwb1bq
WzsHqQe9sioAy+vYee12VS/LSq2Pj3h/qXhCVOk1A/CSvaXg30uLCZmDmlM+giAj
pVCFABVlCk9Zha5EcPOkHT9tJ3DOcQqzUuVp9wpXJY0Fon4ZKXHblT1ONs2Tt+gO
4RyZceCxtSeMFC+xI9twCwKBgHIu4h8+OShz1smLvfLXgKAT04ryWdcxLPeD8u3s
FOwiso3PAP8Y5MZMR3CFJ7Hr6ZDZ1tAvdXhdpmbZDDOPtniQPd/epK+CMbbW2c+5
5piWGnaFfRT6MHjpuDM+/8w8fcefvwHPFugKBAzEWhqfdCefLuqrao+qP4T6/LTj
azLBAoGAOunZ7iWZsyhz/NiGd1Chq5RVZHqISd3RPEDhAQutdww22G9yor9dTWJG
FRM8yDghD7Gk6qlGuqjM3pOwaEaOgCFovKzyWhH9PT6aQsUUHCAIJjQRHAQt4o9S
QeEpnzPydEqbmGwd1VBSlBaEHgPRenDtC03sIwg1v/DeNeZ+QSk=
-----END RSA PRIVATE KEY-----`)

	da := strconv.Itoa(rand2.Int())
	//k := RsaSignWithSha256(privKey)

	//header := `keyId="https://activity.disism.com/actor",headers=(request-target) host date",signature="` + k
	Req(da, privKey)

}

func NewOutbox(c *gin.Context) {
	privKey := []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAzcU2iJkBv6GpgU02t9/6w1ZrKVhlw/iBLN5RmsC7K1jVLEPG
Cc0XWGg26ab8ZomKQfq2DpjS25L1aBjfxkWWnJ13PnLbCTkoMiVGinEFpwuDcONN
93YwZuPTqzrwZCbPc5L2owlI30AVXfl3zacmSgwXHISyJ3Lri5RUbUv1EMMYcRxw
PxPiD1AFZFTfCoB32Tv8vr4zqUJNeh+rnU8BVZfgwLHWwjWazVxTKdirQz8HoAra
ARs4pxpI0Y74FSnMANfnOuwevstV+14hFtObWuCJurVuZkBga9pm2gsFCAdyifaN
oXZMojsKSh0+d6ENuWx0247PxbQjpSNZ+atvAwIDAQABAoIBAE4qDwgejzJ7N4d9
jD6W34vaRQTEpoul13Py067kbXUO3pNc/1pzxN15YPqqyxAsSQ+9K/EpjleIYJ51
bHPFtWuyyTw4pwm+440A6eXm/w2wavKz9cb9527+NkbCMdTphT5S4kuWWE8t6a/z
9tVxHDNemlzj70Ta+6ze4J9QDjyaIMgqsrDKi57IkKPMtuXIDs+hkN0YSwVrveYb
mHkkWANZRhf9sZgAVoKFAPAPONWcEk6/5fdKXURzL9hKd12i8QN91raF2kdSpq26
o8QLmoUYcm69sbq0RTW/asHBUl9W78yVFAqcKFf3veurQXV+RUDdXoyFVZVYOU8V
APH4AoECgYEA9+ff1a5E8V0XzzxEOR5pcq+wzr5+Va+PX6PlfKWg/VcuQW7igqQG
NWeWTL46LjtJ8sjsan4a5CVkrUQVIboMjmaXUvxUQdZFurt93YHb4RYOIJ7As4ub
ngSjzljB607dFbm7FrwPaQpVfu+Jr9NSboHXp0qvGHzbmmAfWTE96sMCgYEA1H0m
KDOG8YFSM/5H4dubPt/uxw84XpRz/h3dfB1b9uVr5pGuUuW4iahU0rH+37M0SwOU
zu1wboPILRCtCnW0EVl631QhdDvgzttgCa41iLedDlPIOCvUUezrTAmWomkzQY7A
DKnmvYtUaJeoH0xLy/KVvqqs9QHvWQ8RGoMBpsECgYBeElJA/qlQwde4HxFwb1bq
WzsHqQe9sioAy+vYee12VS/LSq2Pj3h/qXhCVOk1A/CSvaXg30uLCZmDmlM+giAj
pVCFABVlCk9Zha5EcPOkHT9tJ3DOcQqzUuVp9wpXJY0Fon4ZKXHblT1ONs2Tt+gO
4RyZceCxtSeMFC+xI9twCwKBgHIu4h8+OShz1smLvfLXgKAT04ryWdcxLPeD8u3s
FOwiso3PAP8Y5MZMR3CFJ7Hr6ZDZ1tAvdXhdpmbZDDOPtniQPd/epK+CMbbW2c+5
5piWGnaFfRT6MHjpuDM+/8w8fcefvwHPFugKBAzEWhqfdCefLuqrao+qP4T6/LTj
azLBAoGAOunZ7iWZsyhz/NiGd1Chq5RVZHqISd3RPEDhAQutdww22G9yor9dTWJG
FRM8yDghD7Gk6qlGuqjM3pOwaEaOgCFovKzyWhH9PT6aQsUUHCAIJjQRHAQt4o9S
QeEpnzPydEqbmGwd1VBSlBaEHgPRenDtC03sIwg1v/DeNeZ+QSk=
-----END RSA PRIVATE KEY-----`)

	da := strconv.Itoa(rand2.Int())
	//k := RsaSignWithSha256(privKey)

	//header := `keyId="https://activity.disism.com/actor",headers=(request-target) host date",signature="` + k
	Req2(da, privKey)
}
