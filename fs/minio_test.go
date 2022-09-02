package fs

import (
	"fmt"
	"os"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
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
}

func TestGetMinio(t *testing.T) {
	_, err := NewDefaultMinio().Dial()
	if err != nil {
		t.Error(err)
		return
	}

}

func TestMinio_CreateBucket(t *testing.T) {
	minio, err := NewDefaultMinio().Dial()
	if err != nil {
		t.Error(err)
		return
	}
	if err := minio.MakeBucket("avatar", ""); err != nil {
		t.Error(err)
		return
	}
}

func TestMinio_ListBuckets(t *testing.T) {
	minio, err := NewDefaultMinio().Dial()
	if err != nil {
		t.Error(err)
		return
	}
	if err := minio.ListBuckets(); err != nil {
		t.Error(err)
		return
	}
}

func TestMinio_GetBucketPolicy(t *testing.T) {
	minio, err := NewDefaultMinio().Dial()
	if err != nil {
		t.Error(err)
		return
	}
	policy, err := minio.GetBucketPolicy("avatar")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(policy)
}

func TestMinioFiles_Remove(t *testing.T) {
	minio, err := NewDefaultMinio().Dial()
	if err != nil {
		t.Error(err)
		return
	}
	if err := NewMinioRemoveFile(minio.Client, minio.Ctx, "attach", "fn").Remove(); err != nil {
		t.Error(err)
		return
	}
}
