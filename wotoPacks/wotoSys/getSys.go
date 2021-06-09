package wotoSys

import (
	"github.com/zcalusic/sysinfo"
)

func GetSysInfo() (*sysinfo.SysInfo, error) {
	//current, err := user.Current()
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//}

	//if current.Uid != "0" {
	//	log.Println("requires superuser privilege")
	//	return nil, errors.New("requires superuser privilege")
	//}

	var si sysinfo.SysInfo

	si.GetSysInfo()

	return &si, nil
}
