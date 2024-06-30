
# Instruction
## Install dependencies
```bash
brew install golang-migrate
brew install kyleconroy/sqlc/sqlc
```
## Setup infrastructure
Start postgres container:
```bash
make postgres
```
Create simple_bank database:
```bash
make createdb
```
Run db migration:
```bash
make migrateup
```
## Connect to database
> url: `jdbc:postgresql://localhost:5432/simple_bank`
> username: `root`
> password: `secret`
## Generate handles to the database
```bash
make sqlc
```
## Run tests
```bash
make test
```
