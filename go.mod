module github.com/newmetric/bbloom

go 1.20

require (
	github.com/golang/protobuf v1.5.3
	github.com/newmetric/tetrapod-proto v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/newmetric/tetrapod-proto => ../tetrapod-proto
