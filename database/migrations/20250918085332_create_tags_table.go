package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250918085332CreateTagsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250918085332CreateTagsTable) Signature() string {
	return "20250918085332_create_tags_table"
}

// Up Run the migrations.
func (r *M20250918085332CreateTagsTable) Up() error {
	if !facades.Schema().HasTable("tags") {
		return facades.Schema().Create("tags", func(table schema.Blueprint) {
			table.ID()
			table.String("name", 195)
			table.String("slug", 175)
			table.Unique("slug")
			table.Index("name")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250918085332CreateTagsTable) Down() error {
	return facades.Schema().DropIfExists("tags")
}
