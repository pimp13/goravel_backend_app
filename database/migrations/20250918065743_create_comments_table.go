package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250918065743CreateCommentsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250918065743CreateCommentsTable) Signature() string {
	return "20250918065743_create_comments_table"
}

// Up Run the migrations.
func (r *M20250918065743CreateCommentsTable) Up() error {
	if !facades.Schema().HasTable("comments") {
		return facades.Schema().Create("comments", func(table schema.Blueprint) {
			table.ID()
			table.String("content")
			table.Boolean("is_active").Default(false)
			table.UnsignedBigInteger("user_id")
			table.UnsignedBigInteger("post_id")
			table.Foreign("user_id").References("id").On("users").CascadeOnDelete().CascadeOnUpdate()
			table.Foreign("post_id").References("id").On("posts").CascadeOnDelete().CascadeOnUpdate()
			table.Index("is_active")
			table.Index("user_id")
			table.Index("post_id")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250918065743CreateCommentsTable) Down() error {
	return facades.Schema().DropIfExists("comments")
}
