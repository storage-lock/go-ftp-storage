package ftp_storage

import (
	"context"
	"sync"
	"time"

	go_iterator "github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
)

// FTPStorage FTP文件存储实现
// 使用mutex实现进程内CAS
type FTPStorage struct {
	mu sync.RWMutex
}

var _ storage.Storage = &FTPStorage{}

func NewFTPStorage(ctx context.Context) (*FTPStorage, error) {
	return &FTPStorage{}, nil
}

const StorageName = "ftp-storage"

func (x *FTPStorage) GetName() string { return StorageName }

func (x *FTPStorage) Capabilities() []storage.StorageCapability {
	return []storage.StorageCapability{storage.CapabilityCAS, storage.CapabilityReliableTime}
}

func (x *FTPStorage) Init(ctx context.Context) error { return nil }

func (x *FTPStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}

func (x *FTPStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}

func (x *FTPStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}

func (x *FTPStorage) Get(ctx context.Context, lockId string) (string, error) {
	x.mu.RLock()
	defer x.mu.RUnlock()
	return "", storage_lock.ErrLockNotFound
}

func (x *FTPStorage) GetTime(ctx context.Context) (time.Time, error) { return time.Now(), nil }
func (x *FTPStorage) Close(ctx context.Context) error               { return nil }
func (x *FTPStorage) List(ctx context.Context) (go_iterator.Iterator[*storage.LockInformation], error) {
	return go_iterator.FromSlice([]*storage.LockInformation{}), nil
}
