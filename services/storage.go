package services

import (
	"github.com/fiure-cms/fiure-storage/stores/adminstore"
	"github.com/uretgec/mydb/storage/interfaces"
)

var Store interfaces.Storage

const (
	AdminStore = "adminstore"
)

func SetupStorage(name, path string, bucketList, indexList []string, readOnly bool) error {
	var err error
	if name == AdminStore {
		Store, err = adminstore.NewStore(path, name, bucketList, indexList, readOnly)
	}

	return err
}
