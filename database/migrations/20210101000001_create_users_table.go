package migrations

import (
	"goravel_by_gin/app/enums"

	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20210101000001CreateUsersTable struct{}

// Signature The unique signature for the migration.
func (r *M20210101000001CreateUsersTable) Signature() string {
	return "20210101000001_create_users_table"
}

// Up Run the migrations.
func (r *M20210101000001CreateUsersTable) Up() error {
	return facades.Schema().Create("users", func(table schema.Blueprint) {
		table.ID("id")
		table.String("name")
		table.String("email")
		table.String("password")
		table.Boolean("is_active").Default(false)
		table.Boolean("is_supperadmin").Default(false)
		table.TinyInteger("role").Default(enums.RoleGuest)

		table.Unique("email")
		table.Index("role")
		table.Index("is_active")
		table.Index("is_supperadmin")
		table.TimestampsTz()
	})
}

// Down Reverse the migrations.
func (r *M20210101000001CreateUsersTable) Down() error {
	return facades.Schema().DropIfExists("users")
}
