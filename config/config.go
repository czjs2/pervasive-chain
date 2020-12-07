package config

import (
	"syscall"
)

const SysTimefrom = "2006-01-02 15:04:05"

const LocationTimeZone = "Asia/Shanghai"

//token加解密的key,谨慎修改
const SecretKey = "[]{&ds13SDF*&8a4%fsF11@#aA"

const (
	// SIGUSR1 linux SIGUSR1
	SIGUSR1 = syscall.Signal(0xa)

	// SIGUSR2 linux SIGUSR2
	SIGUSR2 = syscall.Signal(0xc)
)




// 链类型
const BeaconType = "B"
const RelayType = "R"
const SharedType = "S"

// 命令下发间隔时间
const GenCmdIntervalTime = 15

// 通用分页参数
const PageSize = 15
