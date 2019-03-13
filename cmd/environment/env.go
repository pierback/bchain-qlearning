package environment

import (
	"flag"
	"fmt"
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

	/* sip := ut.GetServerIP()
	fmt.Printf("ServerIP: %s", sip) */
	os.Setenv("SIP", "oc-appsrv01.informatik.uni-augsburg.de")

	uip := ut.GetUploadIP()
	fmt.Printf("  UploadIP: %s", uip)
	os.Setenv("UPID", uip)

	bcip := ut.GetBchainIP()
	fmt.Printf("  BchainIP: %s\n", bcip)
	os.Setenv("BCIP", bcip)
}
