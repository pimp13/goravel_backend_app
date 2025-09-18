package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250918071728CreateUserProfilesTable struct{}

// Signature The unique signature for the migration.
func (r *M20250918071728CreateUserProfilesTable) Signature() string {
	return "20250918071728_create_user_profiles_table"
}

// Up Run the migrations.
func (r *M20250918071728CreateUserProfilesTable) Up() error {
	if !facades.Schema().HasTable("user_profiles") {
		return facades.Schema().Create("user_profiles", func(table schema.Blueprint) {
			table.ID()
			table.String("location").Nullable()
			table.String("avatar").Default(facades.Config().GetString("http.url") + "/public/storage/uploads/default-profile.jpg")
			table.String("bio").Nullable()
			table.UnsignedBigInteger("user_id")
			table.Foreign("user_id").References("id").On("users").CascadeOnDelete().CascadeOnUpdate()
			table.Index("user_id")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250918071728CreateUserProfilesTable) Down() error {
	return facades.Schema().DropIfExists("user_profiles")
}
