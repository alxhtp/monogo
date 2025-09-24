migrate-create:
	@sql-migrate new -env=$(env) $(file_name)

migrate-up: # make migrate-up -env=staging
	@sql-migrate up -config=dbconfig.yml -env=$(env)

migrate-down:
	@sql-migrate down -env=$(env)

migrate-status:
	@sql-migrate status -env=$(env)
