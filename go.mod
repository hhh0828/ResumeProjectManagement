module Project/ResumeManagement

go 1.22.4

require gorm.io/driver/postgres v1.5.9

require welcome/hyunho/myapp v0.0.0-00010101000000-000000000000 // indirect

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	gorm.io/gorm v1.25.11

)

replace p.com/test => ./1

replace welcome/hyunho/myapp => ./3
