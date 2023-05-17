package dnscache

import "os"

var logDns bool

func init() {
	logDns = os.Getenv("AM_LOG_DNS") == "true"
}
