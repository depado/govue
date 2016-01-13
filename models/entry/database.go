package entry

import (
	"fmt"
	"strconv"

	"github.com/Depado/govue/database"
	"github.com/boltdb/bolt"
)

// Save saves an Entry in the database.
func (e *Entry) Save() error {
	if !database.Main.Opened {
		return fmt.Errorf("db must be opened before saving")
	}
	return database.Main.DB.Update(func(tx *bolt.Tx) error {
		var err error
		var b *bolt.Bucket
		var enc []byte
		var id uint64

		if b, err = tx.CreateBucketIfNotExists([]byte(Bucket)); err != nil {
			return fmt.Errorf("Error creating bucket : %s", err)
		}
		if enc, err = e.Encode(); err != nil {
			return fmt.Errorf("Could not encode : %s", err)
		}
		if e.ID == 0 {
			if id, err = b.NextSequence(); err != nil {
				return fmt.Errorf("Could not generate ID : %s", err)
			}
			e.ID = int(id)
		}
		return b.Put([]byte(strconv.Itoa(e.ID)), enc)
	})
}

// Delete deletes an Entry from the database.
func (e Entry) Delete() error {
	if !database.Main.Opened {
		return fmt.Errorf("db must be opened before deleting")
	}
	return database.Main.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(Bucket))
		if b == nil {
			return fmt.Errorf("Could not find bucket : %s", Bucket)
		}
		return b.Delete([]byte(strconv.Itoa(e.ID)))
	})
}
