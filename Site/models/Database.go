package models

var Dsn string = "host=localhost user=<user> password=<password> dbname=<db_name> port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func DbMigration() {
	Article{}.Migrate()
	Category{}.Migrate()
	Message{}.Migrate()
	Services{}.Migrate()
	Feature{}.Migrate()
	User{}.Migrate()
	Menu{}.Migrate()
	About{}.Migrate()
}
