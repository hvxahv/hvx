package channel

import "testing"

func TestNew(t *testing.T) {
	TestInitChannelConfig(t)
	nb := NewBroadcast("平手 友梨奈", `
みんながおかしいんじゃないのか
自分は普通だと思ってた
でも何が普通なのか?
その根拠なんかあるわけもなくて
もう誰もいないだろうと思った真夜中
こんな路地ですれ違う人がなぜいるの?
独り占めしてたはずの不眠症が
私だけのものじゃなくて落胆した
`)

	nb.New()
}
