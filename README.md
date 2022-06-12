# Testing go-mysqldump

To keep it simple did not documented bash/mysql output for most of the commands.

```sh
## running mysql in a docker container
docker run --name mysql -e MYSQL_ROOT_PASSWORD=password -p 3306:3306/tcp -d mysql
mysql -h127.0.0.1 -uroot -ppassword --execute="show databases;"

## creating test db and table
CREATE DATABASE testdb;
CREATE TABLE testdb.testtb (testcol INT);
INSERT INTO testdb.testtb VALUES (1);
INSERT INTO testdb.testtb VALUES (2);
INSERT INTO testdb.testtb VALUES (3);
SELECT * FROM testdb.testtb;

## created dumps dir
mkdir dumps

## setting up golang project
mkdir godump
cd godump/
touch main.go
go mod init github.com/mol-george/godump.git
go get github.com/jamf/go-mysqldump
go mod tidy
code .

## Testing
➜ ls -l dumps/
total 0

➜ go build -o godump

➜ ./godump
[1 2 3]

➜ ls -l dumps/
-rw-r--r--  1 gm  staff  1786 12 Jun 14:33 testdb-220612T143348.sql
```