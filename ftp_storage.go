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
// 尚未接入真实 FTP 服务端，仅为占位实现，不可用于分布式锁。
// 真正接入时需借助服务端原子语义（如临时文件+原子重命名）实现 CAS。
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
	return []storage.StorageCapability{storage.CapabilityReliableTime, storage.CapabilityWatch}
}

func (x *FTPStorage) Watch(ctx context.Context, lockId string) (<-chan *storage.LockInformation, error) {
	return storage.NewPollingLockWatcher(x, 500*time.Millisecond).Watch(ctx, lockId)
}

func (x *FTPStorage) Init(ctx context.Context) error { return nil }

func (x *FTPStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
}

func (x *FTPStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
}

func (x *FTPStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return storage_lock.ErrStorageCapabilityMissing
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
