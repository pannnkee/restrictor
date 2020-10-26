package LeakyBucket

import (
	"sync"
	"time"
)

type Bucket struct {
	capacity uint
	remaining uint
	reset time.Time
	rate time.Duration
	mutex sync.Mutex
}

func (this *Bucket) Capacity() uint {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.capacity
}

func (this *Bucket) Remaining() uint {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.remaining
}

func (this *Bucket) Reset() time.Time {
	this.remaining = this.capacity
	return this.reset
}

func (this *Bucket) Add(amount uint) (BucketState, error) {
	//判断时间有没有过期 过期重置
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if time.Now().After(this.reset) {
		this.reset = time.Now().Add(this.rate)
		this.remaining = this.capacity
	}

	//判断桶是不是满了 满了则返回state
	if amount > this.remaining {
		return BucketState{
			Capacity:  this.capacity,
			Remaining: this.remaining,
			Reset:     this.reset,
		}, ErrorFull
	}

	//将桶剩余量减一 返回状态
	this.remaining -= amount
	return BucketState{
		Capacity:  this.capacity,
		Remaining: this.remaining,
		Reset:     this.reset,
	}, nil
}

type Storage struct {
	Bucket map[string]*Bucket
}

func New() *Storage {
	return &Storage{
		Bucket: make(map[string]*Bucket),
	}
}

func (this *Storage) Create(name string, capacity uint, rate time.Duration) (BucketI, error) {
	b,ok := this.Bucket[name]
	if ok {
		return b,nil
	}
	b = &Bucket{
		capacity:  capacity,
		remaining: capacity,
		reset:     time.Now().Add(rate),
		rate:      rate,
	}
	this.Bucket[name] = b
	return b, nil
}