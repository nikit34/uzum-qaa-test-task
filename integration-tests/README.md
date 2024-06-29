brew install golang-migrate
make postgres
make createdb
make migrateup
