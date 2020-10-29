// 工厂模式
// author: baoqiang
// time: 2020/10/29 8:18 下午
package gopat

import (
	"io"
	"log"
	"os"
)

// 定义多种类型，不同类型返回不同实现
type StorageType int

const (
	DiskStorageType StorageType = 1 << iota
	MemoryStoryType
	//TempStorageType
)

// 定义有某个能力的类型
type Store interface {
	Open(s string) (io.ReadWriteCloser, error)
}

// 新建一个该类型的实例
func NewStore(st StorageType) Store {
	switch st {
	case DiskStorageType:
		return newDiskStorage()
		//case MemoryStory:
		//return newMemoryStorage()
	}
	return nil
}

// 该类型的具体实现
type diskStorage struct{}

func (d diskStorage) Open(s string) (io.ReadWriteCloser, error) {
	log.Printf("disk storage open")
	return os.Create(s)
}

func newDiskStorage() Store {
	return diskStorage{}
}

func RunFactory() {
	s := NewStore(DiskStorageType)
	ds, _ := s.Open("hello.txt")
	defer ds.Close()
	ds.Write([]byte("data\n"))
}
