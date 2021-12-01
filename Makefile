bind_migrations:
	go-bindata -pkg migrations -o db/migrations.go -prefix='db/migrations/' db/migrations/
