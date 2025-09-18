package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250918083335CreateUserFollowersTable struct{}

// Signature The unique signature for the migration.
func (r *M20250918083335CreateUserFollowersTable) Signature() string {
	return "20250918083335_create_user_followers_table"
}

// Up Run the migrations.
func (r *M20250918083335CreateUserFollowersTable) Up() error {
	if !facades.Schema().HasTable("user_followers") {
		return facades.Schema().Create("user_followers", func(table schema.Blueprint) {
			table.UnsignedBigInteger("follower_id")
			table.UnsignedBigInteger("followed_id")
			table.Foreign("follower_id").References("id").On("users").CascadeOnUpdate().CascadeOnDelete()
			table.Foreign("followed_id").References("id").On("users").CascadeOnUpdate().CascadeOnDelete()
			table.Primary("follower_id", "followed_id")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250918083335CreateUserFollowersTable) Down() error {
	return facades.Schema().DropIfExists("user_followers")
}
