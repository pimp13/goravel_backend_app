package database

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"

	"goravel_by_gin/database/migrations"
	"goravel_by_gin/database/seeders"
)

type Kernel struct {
}

func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20210101000001CreateUsersTable{},
		&migrations.M20210101000002CreateJobsTable{},
		&migrations.M20250915134738CreateCategoriesTable{},
		&migrations.M20250916110648CreatePostsTable{},
		&migrations.M20250918065743CreateCommentsTable{},
		&migrations.M20250918071728CreateUserProfilesTable{},
		&migrations.M20250918083335CreateUserFollowersTable{},
		&migrations.M20250918085332CreateTagsTable{},
		&migrations.M20250918085628CreatePostTagTable{},
	}
}

func (kernel Kernel) Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.DatabaseSeeder{},
	}
}
