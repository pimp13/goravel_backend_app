package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250916110648CreatePostsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250916110648CreatePostsTable) Signature() string {
	return "20250916110648_create_posts_table"
}

// Up Run the migrations.
func (r *M20250916110648CreatePostsTable) Up() error {
	if !facades.Schema().HasTable("posts") {
		return facades.Schema().Create("posts", func(table schema.Blueprint) {
			table.ID()
			table.String("title", 195)
			table.String("summary").Nullable()
			table.Text("content")
			table.Boolean("is_active").Default(false)
			table.String("slug", 175)
			table.String("image_url").Nullable()
			table.UnsignedBigInteger("user_id")
			table.UnsignedBigInteger("category_id")
			table.TimestampsTz()

			table.Foreign("category_id").References("id").On("categories").CascadeOnDelete().CascadeOnUpdate()
			table.Foreign("user_id").References("id").On("users").CascadeOnDelete().CascadeOnUpdate()
			table.Unique("slug")
			table.Index("is_active")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250916110648CreatePostsTable) Down() error {
	return facades.Schema().DropIfExists("posts")
}
