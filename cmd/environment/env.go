package environment

import (
	"flag"
	"fmt"
	"os"

	ut "github.com/pierback/bchain-qlearning/pkg/utils"
)

var (
	BcFlag     *string
	DplFlag    *string
	parentFlag *string
	SimFlag    *int
)

func SetEnvVars() {
	DplFlag = flag.String("dp", "", "deploy sm")
	BcFlag = flag.String("bc", "watch", "test bchain")
	parentFlag = flag.String("pt", "upgrade", "deploy or upgrade parent contract")
	SimFlag = flag.Int("sim", 0, "simulate user")
	flag.Parse()

	ws := ut.GetEthWsAddr()
	os.Setenv("WS", ws)

	// sip := ut.GetServerIP()
	sip := "oc-appsrv01.informatik.uni-augsburg.de"
	// sip := "localhost"
	fmt.Printf("ServerIP: %s", sip)
	os.Setenv("SIP", sip)

	uip := ut.GetUploadIP()
	fmt.Printf("  UploadIP: %s", uip)
	os.Setenv("UPID", uip)

	bcip := ut.GetBchainIP()
	fmt.Printf("  BchainIP: %s\n", bcip)
	os.Setenv("BCIP", bcip)
}
