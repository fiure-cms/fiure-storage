package adminstore

import (
	sniperstorage "github.com/uretgec/mydb/storage/sniper"
)

type AdminStore struct {
	Storage *sniperstorage.Store
}

func NewStore(dbPath, dbName string, bucketList, indexList []string, readOnly bool) (*AdminStore, error) {
	store, err := sniperstorage.NewStore(
		bucketList,
		indexList,
		dbPath,
		dbName,
		readOnly,
	)
	return &AdminStore{
		Storage: store,
	}, err
}

func (ms *AdminStore) CloseStore() {
	ms.Storage.CloseStore()
}

func (ms *AdminStore) SyncStore() {
	ms.Storage.SyncStore()
}
func (ms *AdminStore) HasBucket(bucketName []byte) bool {
	return ms.Storage.HasBucket(bucketName)
}

func (ms *AdminStore) Set(bucketName []byte, k []byte, data []byte) ([]byte, error) {
	return ms.Storage.Set(bucketName, k, data)
}

func (ms *AdminStore) Get(bucketName []byte, k []byte) ([]byte, error) {
	return ms.Storage.Get(bucketName, k)
}

func (ms *AdminStore) MGet(bucketName []byte, keys ...[]byte) (interface{}, error) {
	return ms.Storage.MGet(bucketName, keys...)
}

func (ms *AdminStore) List(bucketName []byte, k []byte, perpage int) ([]string, error) {
	return ms.Storage.List(bucketName, k, perpage)
}

func (ms *AdminStore) PrevList(bucketName []byte, k []byte, perpage int) ([]string, error) {
	return ms.Storage.PrevList(bucketName, k, perpage)
}

func (ms *AdminStore) KeyExist(bucketName []byte, k []byte) (bool, error) {
	return ms.Storage.KeyExist(bucketName, k)
}

func (ms *AdminStore) ValueExist(bucketName []byte, v []byte) (bool, error) {
	return ms.Storage.ValueExist(bucketName, v)
}

func (ms *AdminStore) Delete(bucketName []byte, k []byte) error {
	return ms.Storage.Delete(bucketName, k)
}

func (ms *AdminStore) StatsBucket(bucketName []byte) int {
	return ms.Storage.StatsBucket(bucketName)
}

func (ms *AdminStore) ListBucket() ([]string, error) {
	return ms.Storage.ListBucket()
}

func (ms *AdminStore) DeleteBucket(bucketName []byte) error {
	return ms.Storage.DeleteBucket(bucketName)
}

func (ms *AdminStore) Backup(path, filename string) error {
	return ms.Storage.Backup(path, filename)
}

func (ms *AdminStore) Restore(path, filename string) error {
	return ms.Storage.Restore(path, filename)
}
