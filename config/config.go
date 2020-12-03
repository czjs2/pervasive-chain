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

// 通用分页参数
const PageSize = 15

// 出块时间 秒
const BlockTime = 15

const NodeOffLineTime = 45 // 秒

const BChain = "b"

const RChain = "r"

const SChain = "s"

// 心跳时间
const HeartBeatTime = 15







