package environment

import (
	"flag"
	"os"

	ut "github.com/pierback/bchain-qlearning/pkg/utils"
)

var (
	BcFlag  *string
	DplFlag *string
)

func SetEnvVars() {
	DplFlag = flag.String("dp", "", "deploy sm")
	BcFlag = flag.String("bc", "watch", "test bchain")
	flag.Parse()

	ws := ut.GetEthWsAddr()
	os.Setenv("WS", ws)

	sip := ut.GetServerIP()
	os.Setenv("SIP", sip)

	uip := ut.GetUploadIP()
	os.Setenv("UPID", uip)

	bcip := ut.GetBchainIP()
	os.Setenv("BCIP", bcip)
}
