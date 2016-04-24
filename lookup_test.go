package lookup

import (
	"os/exec"
	"regexp"
	"testing"
)

var testIPs = []string{"8.8.8.8", "127.0.0.1"}

func TestLookupAddr(t *testing.T) {
	for _, ip := range testIPs {
		host, err := LookupAddr(ip)
		if err != nil {
			t.Error(ip, err)
			continue
		}
		if host == "" {
			t.Errorf("empty host")
			continue
		}
		expected := ping(ip)
		if expected == "" {
			t.Errorf("empty expected host")
		}
		if expected != host {
			t.Errorf("different results %s:\texp:%v\tgot:%v", ip, expected, host)
		}
	}
}

func ping(ip string) string {
	output, err := exec.Command("ping", "-n", "1", "-a", ip).Output()
	if err != nil {
		return ""
	}

	re := regexp.MustCompile(`Pinging (.*?) \[\d+\.\d+\.\d+\.\d+\]`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}
