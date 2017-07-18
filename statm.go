package statm

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Statm struct {
	PID    int
	VmSize int
	VmRss  int
	Shared int
	Trs    int
	Lrs    int
	Drs    int
	Dt     int
}

func readStatm(pid int) (string, error) {
	bin, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/statm", pid))
	if err != nil {
		return "", err
	}
	return string(bin), nil
}

func New(pid ...int) *Statm {
	var p int
	var ret = &Statm{}
	if len(pid) != 0 && pid[0] != 0 {
		p = pid[0]
	} else {
		p = os.Getpid()
	}

	statmString, err := readStatm(p)
	if err != nil {
		return nil
	}

	for i, v := range strings.Split(statmString, " ") {
		switch i {
		case 0:
			ret.VmSize, _ = strconv.Atoi(v)
		case 1:
			ret.VmRss, _ = strconv.Atoi(v)
		case 2:
			ret.Shared, _ = strconv.Atoi(v)
		case 3:
			ret.Trs, _ = strconv.Atoi(v)
		case 4:
			ret.Lrs, _ = strconv.Atoi(v)
		case 5:
			ret.Drs, _ = strconv.Atoi(v)
		case 6:
			ret.Dt, _ = strconv.Atoi(v)
		}
	}
	ret.PID = p

	return ret
}
