package jsnowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	// t := time.Date(2022, 03, 20, 00, 00, 00, 00, time.Local).UnixNano()
	//开始时间戳 2022-03-20  毫秒
	startTime        int64 = 1647705600000000000
	workerIdBits     uint8 = 10
	datacenterIdBits uint8 = 12
	maxWorkerId      int64 = -1 ^ (-1 << workerIdBits)
	maxDatacenterId  int64 = -1 ^ (-1 << datacenterIdBits)
	timeShift              = workerIdBits + datacenterIdBits
	workerShift            = datacenterIdBits
)

type Machine struct {
	mutex         sync.Mutex
	lastTimestamp int64
	wId           int64
	datacenter    int64
}

// NewMachine
func NewMachine(wId int64) (*Machine, error) {
	if wId < 0 || wId > maxWorkerId {
		return nil, errors.New("wId is out of range")
	}
	return &Machine{
		lastTimestamp: 0,
		wId:           wId,
		datacenter:    0,
	}, nil
}

// GenerateId
func (w *Machine) GenerateId() int64 {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.lastTimestamp == now {
		w.datacenter++
		if w.datacenter > maxDatacenterId {
			for now <= w.lastTimestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.datacenter = 0
		w.lastTimestamp = now
	}

	ID := (now-startTime)<<timeShift | (w.wId << workerShift) | (w.datacenter)
	return ID
}
