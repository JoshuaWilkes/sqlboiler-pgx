module github.com/volatiletech/sqlboiler/v4

go 1.16

require (
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/davecgh/go-spew v1.1.1
	github.com/ericlagergren/decimal v0.0.0-20211103172832-aca2edc11f73
	github.com/friendsofgo/errors v0.9.2
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/google/go-cmp v0.5.8
	github.com/google/uuid v1.3.0 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jackc/pgx/v5 v5.3.1
	github.com/lib/pq v1.10.6
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.12.0
	github.com/stretchr/testify v1.8.2
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/randomize v0.0.1
	github.com/volatiletech/strmangle v0.0.4
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)

retract (
	v4.10.0 // Generated models are invalid due to a wrong assignment
	v4.9.0 // Generated code shows v4.8.6, messed up commit tagging and untidy go.mod
)
