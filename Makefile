# 运行 postgres 镜像
postgres_up:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14.0
# 删除 postgres 容器
postgres_down:
	docker stop postgres
	docker rm postgres

# 新建数据库
create_db:
	docker exec -it postgres createdb --username=root --owner=root information

# 删除数据库
drop_db:
	docker exec -it postgres dropdb information

# 初始化数据库迁移文件
migrate_init:
	migrate create -ext sql -dir pkg/db/migrations -seq $(name)

# 执行数据库迁移命令 1
migrate_up:
	migrate -path pkg/db/migrations -database "postgres://root:password@localhost:5432/information?sslmode=disable" -verbose up

# 删除数据库迁移命令 1
migrate_down:
	migrate -path pkg/db/migrations -database "postgres://root:password@localhost:5432/information?sslmode=disable" -verbose down
