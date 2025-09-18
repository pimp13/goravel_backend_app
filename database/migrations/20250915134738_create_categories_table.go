package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250915134738CreateCategoriesTable struct{}

// Signature The unique signature for the migration.
func (r *M20250915134738CreateCategoriesTable) Signature() string {
	return "20250915134738_create_categories_table"
}

// Up Run the migrations.
func (r *M20250915134738CreateCategoriesTable) Up() error {
	if !facades.Schema().HasTable("categories") {
		return facades.Schema().Create("categories", func(table schema.Blueprint) {
			table.ID()
			table.String("name", 195)
			table.Boolean("is_active").Default(true)
			table.String("description").Nullable()
			table.String("slug", 175)
			table.Unique("slug")
			table.Index("is_active")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250915134738CreateCategoriesTable) Down() error {
	return facades.Schema().DropIfExists("categories")
}
