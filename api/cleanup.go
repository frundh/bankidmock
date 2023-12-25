package api

import (
	"encoding/json"
	"time"

	"go.etcd.io/bbolt"
)

func CleanUp(db *bbolt.DB) error {
	return db.Update(func(tx *bbolt.Tx) error {
		ordersToDelete := []string{}
		bucket := tx.Bucket([]byte("orders"))
		if err := bucket.ForEach(func(k, v []byte) error {
			var order order
			if err := json.Unmarshal(v, &order); err != nil {
				return err
			}

			if time.Now().Sub(time.Unix(order.CreatedAt, 0)) >= 3*time.Minute {
				ordersToDelete = append(ordersToDelete, string(k))
			}

			return nil
		}); err != nil {
			return err
		}

		for _, key := range ordersToDelete {
			if err := bucket.Delete([]byte(key)); err != nil {
				return err
			}
		}

		return nil
	})
}
