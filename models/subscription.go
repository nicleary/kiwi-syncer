package models

import "time"

type ZimSubscription struct {
	Name          string    `json:"name" redis:"name"`
	ID            string    `json:"id" redis:"id"`
	SubscribedAt  time.Time `json:"subscribed_at" redis:"subscribed_at"`
	LastSyncedAt  time.Time `json:"last_synced_at" redis:"last_synced_at"`
	LastUpdatedAt time.Time `json:"last_updated_at" redis:"last_updated_at"`
}
