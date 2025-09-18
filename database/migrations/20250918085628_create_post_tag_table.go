package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250918085628CreatePostTagTable struct{}

// Signature The unique signature for the migration.
func (r *M20250918085628CreatePostTagTable) Signature() string {
	return "20250918085628_create_post_tag_table"
}

// Up Run the migrations.
func (r *M20250918085628CreatePostTagTable) Up() error {
	if !facades.Schema().HasTable("post_tag") {
		return facades.Schema().Create("post_tag", func(table schema.Blueprint) {
			table.UnsignedBigInteger("post_id")
			table.UnsignedBigInteger("tag_id")
			table.Foreign("post_id").References("id").On("posts").CascadeOnDelete().CascadeOnUpdate()
			table.Foreign("tag_id").References("id").On("tags").CascadeOnDelete().CascadeOnUpdate()
			table.Primary("post_id", "tag_id")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250918085628CreatePostTagTable) Down() error {
	return facades.Schema().DropIfExists("post_tag")
}
