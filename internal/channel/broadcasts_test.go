package channel

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	TestInitChannelConfig(t)

	title := "角が曲がる"
	article := `
みんながおかしいんじゃないのか
自分は普通だと思ってた
でも何が普通なのか?
その根拠なんかあるわけもなくて
もう誰もいないだろうと思った真夜中
こんな路地ですれ違う人がなぜいるの?
独り占めしてたはずの不眠症が
私だけのものじゃなくて落胆した
らしさって 一体何?
あなたらしく生きればいいなんて
人生がわかったかのように
上から何を教えてくれるの?
周りの人間に決めつけられた
思い通りのイメージになりたくない
そんなこと考えてたら眠れなくなった
だからまたそこの角を曲がる
星空さえも中途半端だ
街の灯りが明るすぎて
明日が晴れようと雨だろうと
変わらない今日がやって来るだけ
本当の自分はそうじゃない
こうなんだと 否定したところで
みんな他人のことに興味ないし
えっ なんで泣いてんだろ?
だって近くにいたって 誰もちゃんと見てはくれず
まるで何かの景色みたいに映っているんだろうな
フォーカスのあってない被写体が
泣いていようと 睨みつけようと どうだっていいんだ
わかってもらおうとすれば ギクシャクするよ
与えられた場所で 求められる私でいれば 嫌われないんだよね?
問題起こさなければ しあわせをくれるんでしょう?
らしさって 一体何?
あなたらしく微笑んで なんて
微笑みたくない そんな一瞬も
自分をどうやれば殺せるだろう?
みんなが期待するような人に
絶対になれなくてごめんなさい
ここいいるのに気づいてもらえないから
一人きりで角を曲がる
Ah, ah, ah, ah
`
	nb, err := NewBroadcast(title, article, 692910076694757377, 692635608323948545)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := nb.New(); err != nil {
		fmt.Println(err)
		return 
	}
}

func TestBroadcasts_QueryLisByCID(t *testing.T) {
	TestInitChannelConfig(t)

	nbc := NewBroadcastCID(692668434193383425)
	r, err := nbc.QueryLisByCID()
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(r)
}