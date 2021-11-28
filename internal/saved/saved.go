package saved

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/ipfs"
	"github.com/spf13/viper"
	"strings"
)

// Upload When uploading files, you need to use asymmetric encryption for encryption.
func Upload() {

	broad := fmt.Sprintf(`
<!doctype html>
<html>
<head>
<meta charset='UTF-8'><meta name='viewport' content='width=device-width initial-scale=1'>
</style><title></title>
</head>
<body>
<div>
<h1>HELLO</h1>
<p>THIS HVXAHV SAVED FILE2.</p>
</div>
</body>
</html>
`)

	cid, err := ipfs.GetIPFS().Add(strings.NewReader(broad))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cid)
	fmt.Println(fmt.Sprintf("%s%s", viper.GetString("ipfs_gateway"), cid))
}

