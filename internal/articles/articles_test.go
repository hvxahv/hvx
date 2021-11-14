package articles

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cache"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
)


func TestInitDB(t *testing.T) {

	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	// Initialize the database.
	n :=  cockroach.NewDBAddr()
	if err2 := n.InitDB(); err2 != nil {
		return
	}

	// If a configs file is found, read it in.
	if err3 := viper.ReadInConfig(); err3 == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}


	cache.InitRedis(1)

}

func TestNewArticles(t *testing.T) {
	TestInitDB(t)

	na := NewArticles(696901249244790785, "HVTURINGGA", "二人セゾン", "僕もセゾン", `
<p>
二人セゾン 二人セゾン
春夏で恋をして
二人セゾン 二人セゾン
秋冬で去って行く
一緒に過ごした季節よ
後悔はしてないか?
二人セゾン
道端咲いてる雑草にも
名前があるなんて忘れてた
気づかれず踏まれても
悲鳴を上げない存在
誰かと話すのが面倒で
目を伏せて聴こえない振りしてた
君は突然
僕のイアホン外した
What did you say now?
太陽が戻って来るまでに
大切な人ときっと出会える
見過ごしちゃ もったいない
愛を拒否しないで
君はセゾン 君はセゾン
僕の前に現れて
君はセゾン 君はセゾン
日常を輝かせる
昨日と違った景色よ
生きるとは変わること
君はセゾン
Ha
街を吹き抜ける風の中
何かの香りがしてたのに
振り返る余裕とか
興味もなかった
自分の半径1メートル
見えないバリア張った別世界
そんな僕を
連れ出してくれたんだ
What made you do that?
一瞬の光が重なって
折々の色が四季を作る
そのどれが欠けたって
永遠は生まれない
二人セゾン 二人セゾン
春夏で恋をして
二人セゾン 二人セゾン
秋冬で去って行く
儚く切ない月日よ
忘れないで
花のない桜を見上げて
満開の日を想ったことはあったか?
想像しなきゃ
夢は見られない
心の窓
春夏秋冬 生まれ変われると
別れ際 君に教えられた
君はセゾン 君はセゾン
僕の前に現れて
君はセゾン 君はセゾン
日常を輝かせる
二人セゾン 二人セゾン
春夏で恋をして
二人セゾン 二人セゾン
秋冬で去って行く
初めて感じたときめき
思い出はカレンダー
二人セゾン
Ha
僕もセゾン
</p>
`, false)
	err := na.New()
	if err != nil {
		return 
	}

	
}

func TestNewArticles2(t *testing.T) {
	TestInitDB(t)

	na := NewStatus(696901249244790785,"HVTURINGGA", `
<p>
二人セゾン 二人セゾン
春夏で恋をして
二人セゾン 二人セゾン
秋冬で去って行く
一緒に過ごした季節よ
後悔はしてないか?
</p>
`, false)
	err := na.New()
	if err != nil {
		return
	}

}

func TestArticles_FindArticlesByAccountID(t *testing.T) {
	TestInitDB(t)

	n := NewArticlesByAccountID(696901249244790785)
	articles, err := n.FindByAccountID()
	if err != nil {
		return
	}
	fmt.Println(articles)

}

func TestArticles_FindArticleByID(t *testing.T) {
	TestInitDB(t)

	n := NewArticleID(701716343427432449)
	article, err := n.FindByID()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(article)
}

func TestArticles_DeleteByURL(t *testing.T) {
	TestInitDB(t)
	
	d := NewArticleURL("https://mas.to/users/hvturingga/statuses/107043507276353003")
	err := d.DeleteByURL()
	if err != nil {
		log.Println(err)
		return 
	}
}