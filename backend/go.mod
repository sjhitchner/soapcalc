module github.com/sjhitchner/soapcalc/backend

go 1.14

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/friendsofgo/errors v0.9.2
	github.com/kat-co/vala v0.0.0-20170210184112-42e1d8b61f12
	github.com/lib/pq v1.2.1-0.20191011153232-f91d3411e481
	github.com/spf13/viper v1.6.3
	github.com/vektah/gqlparser/v2 v2.1.0
	github.com/volatiletech/null v8.0.0+incompatible
	github.com/volatiletech/randomize v0.0.1
	github.com/volatiletech/sqlboiler v3.7.1+incompatible
	github.com/volatiletech/sqlboiler/v4 v4.2.0
	github.com/volatiletech/strmangle v0.0.1
	github.com/web-ridge/sqlboiler-graphql-schema v1.0.4 // indirect
	github.com/web-ridge/utils-go/boilergql v0.0.0-20200622151500-e8881aa1e17a // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
)

replace github.com/web-ridge/gqlgen-sqlboiler/v2 => /Users/steve/src/github.com/sjhitchner/gqlgen-sqlboiler

replace github.com/web-ridge/sqlboiler-graphql-schema => /Users/steve/src/github.com/sjhitchner/sqlboiler-graphql-schema
