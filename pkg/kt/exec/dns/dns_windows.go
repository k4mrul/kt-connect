package dns

import (
	"fmt"
	"github.com/alibaba/kt-connect/pkg/kt/util"
	"os/exec"
)

// SetDnsServer set dns server records
func (s *Cli) SetDnsServer(dnsServers []string, isDebug bool) (err error) {
	// Windows dns config is set on device, so explicit removal is unnecessary
	for i, dns := range dnsServers {
		if i == 0 {
			// run command: netsh interface ip set dnsservers name=KtConnectTunnel source=static address=8.8.8.8
			err = util.RunAndWait(exec.Command("netsh",
				"interface",
				"ip",
				"set",
				"dnsservers",
				fmt.Sprintf("name=%s", s.GetName()),
				"source=static",
				fmt.Sprintf("address=%s", dns),
			), isDebug)
		} else {
			// run command: netsh interface ip add dnsservers name=KtConnectTunnel address=4.4.4.4
			err = util.RunAndWait(exec.Command("netsh",
				"interface",
				"ip",
				"add",
				"dnsservers",
				fmt.Sprintf("name=%s", s.GetName()),
				fmt.Sprintf("address=%s", dns),
			), isDebug)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
