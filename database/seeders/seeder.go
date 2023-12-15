package seeders

func Seeder() {
	PermissionSeeder()
	RoleSeeder()
	RoleHasPermissionSeeder()
	UserSeeder()
	UserHasRoleSeeder()
	TodoSeeder()
}
