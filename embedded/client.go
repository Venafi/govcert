package embedded

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/Venafi/govcert"
)

var vcertPath string

func init() {
	d, err := ioutil.TempDir("", "vcert")
	if err != nil {
		log.Fatalf("Unable to find temp dir: %s", err)
	}
	assets := AssetNames()
	if len(assets) == 0 {
		log.Fatalf("vCert binary asset not found")
	}

	vcertPath = filepath.Join(d, assets[0])
	if err = RestoreAsset(d, assets[0]); err != nil {
		log.Fatalf("Unable to restore asset to tmp directory")
	}

	if err = exec.Command(vcertPath, "-accept-eula").Run(); err != nil {
		log.Fatalf("Error executing vcert: %s", err)
	}
}

func NewClient(apikey string, url string) govcert.Client {
	return govcert.NewClient(vcertPath, apikey, url)
}

func NewClientTPP(username string, password string, url string) govcert.Client {
	return govcert.NewClientTPP(vcertPath, username, password, url)
}
