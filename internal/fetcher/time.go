package fetcher

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func Fetch(server string) (time.Time, error) {
	return ntp.Time(server)
}

func FormatRFC1123(t time.Time) string {
	return t.Format(time.RFC1123)
}

func ValidateServer(server string) error {
	if server == "" {
		return fmt.Errorf("ntp server name is empty")
	}
	return nil
}
