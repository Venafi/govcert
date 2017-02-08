package vcert

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"github.com/opencredo/govcert"
)

var vcertPath string

func init() {
	d, err := ioutil.TempDir("", "vcert")
	if err != nil {
		log.Fatalf("Unable to find temp dir: %s", err)
	}
	log.Println(d)
	assets := AssetNames()
	if len(assets) == 0 {
		log.Fatalf("vCert binary asset not found")
	}
	vcertPath = strings.Join([]string{d, assets[0]}, string(os.PathSeparator))
	if err = RestoreAsset(d, assets[0]); err != nil {
		log.Fatalf("Unable to restore asset to tmp directory")
	}

	if err = exec.Command(vcertPath, "-accept-eula").Run(); err != nil {
		log.Fatalf("Error executing vcert: %s", err)
	}
	// RestoreAssets(d, "vcert")

	// f, _ := ioutil.TempFile("", "")
	// f.Chmod(0755)
	// stat, _ := f.Stat()
	// spew.Dump(stat)
}

func NewClient() *govcert.Client {
	return govcert.NewClient(vcertPath)
}
