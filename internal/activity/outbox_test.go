package activity

import (
	"encoding/json"
	"log"
	"testing"
)

func TestMessages_Outbox(t *testing.T) {
	IniTestConfig(t)

	//Prepare the data first.
	nf := NewFollowRequest("hvturingga", "https://mas.to/users/hvturingga")
	data, err := json.Marshal(nf)
	if err != nil {
		log.Println(err)
		return
	}

	nar := NewActivityRequest(nf.Actor, nf.Object, data, []byte(getPrivk()))

	nar.Send()

}

func TestNewAccept(t *testing.T) {
	IniTestConfig(t)

	//name := "hvturingga"
	//actor := "https://mas.to/users/hvturingga"
	//oid := "https://mas.to/47a0d162-db65-4518-b875-b743b70734c6"
	//object := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), name)

	//na := NewFollowAccept(name, actor, oid, object)

	//data, err := json.Marshal(na)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//nar := NewActivityRequest(object, actor, data, []byte(getPrivk()))
	//nar.Accept()
}

func TestNewArticle(t *testing.T) {
	IniTestConfig(t)

	//na := NewArticle(698898124205195265, "\n<p>\n二人セゾン 二人セゾン\n春夏で恋をして\n二人セゾン 二人セゾン\n秋冬で去って行く\n一緒に過ごした季節よ\n後悔はしてないか?\n二人セゾン\n道端咲いてる雑草にも\n名前があるなんて忘れてた\n気づかれず踏まれても\n悲鳴を上げない存在\n誰かと話すのが面倒で\n目を伏せて聴こえない振りしてた\n君は突然\n僕のイアホン外した\nWhat did you say now?\n太陽が戻って来るまでに\n大切な人ときっと出会える\n見過ごしちゃ もったいない\n愛を拒否しないで\n君はセゾン 君はセゾン\n僕の前に現れて\n君はセゾン 君はセゾン\n日常を輝かせる\n昨日と違った景色よ\n生きるとは変わること\n君はセゾン\nHa\n街を吹き抜ける風の中\n何かの香りがしてたのに\n振り返る余裕とか\n興味もなかった\n自分の半径1メートル\n見えないバリア張った別世界\nそんな僕を\n連れ出してくれたんだ\nWhat made you do that?\n一瞬の光が重なって\n折々の色が四季を作る\nそのどれが欠けたって\n永遠は生まれない\n二人セゾン 二人セゾン\n春夏で恋をして\n二人セゾン 二人セゾン\n秋冬で去って行く\n儚く切ない月日よ\n忘れないで\n花のない桜を見上げて\n満開の日を想ったことはあったか?\n想像しなきゃ\n夢は見られない\n心の窓\n春夏秋冬 生まれ変われると\n別れ際 君に教えられた\n君はセゾン 君はセゾン\n僕の前に現れて\n君はセゾン 君はセゾン\n日常を輝かせる\n二人セゾン 二人セゾン\n春夏で恋をして\n二人セゾン 二人セゾン\n秋冬で去って行く\n初めて感じたときめき\n思い出はカレンダー\n二人セゾン\nHa\n僕もセゾン\n</p>\n")
	//data, err := json.Marshal(na)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//name := "hvturingga"
	//actor := "https://mas.to/users/hvturingga"
	//object := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), name)
	//
	//nar := NewActivityRequest(object, actor, data, []byte(getPrivk()))
	//nar.Article()
}

func TestNewDelete(t *testing.T) {
	IniTestConfig(t)

	//698895044614160385
	//na := NewDelete(698895044614160385)
	//data, err := json.Marshal(na)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//name := "hvturingga"
	//actor := "https://mas.to/users/hvturingga"
	//object := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), name)
	//
	//nar := NewActivityRequest(object, actor, data, []byte(getPrivk()))
	//nar.Article()
}

// reply
/*
{
	"@context": "https://www.w3.org/ns/activitystreams",

	"id": "https://my-example.com/create-hello-world",
	"type": "Create",
	"actor": "https://my-example.com/actor",

	"object": {
		"id": "https://my-example.com/hello-world",
		"type": "Note",
		"published": "2018-06-23T17:17:11Z",
		"attributedTo": "https://my-example.com/actor",
		"inReplyTo": "https://mastodon.social/@Gargron/100254678717223630",
		"content": "<p>Hello world</p>",
		"to": "https://www.w3.org/ns/activitystreams#Public"
	}
}
 */