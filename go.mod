module github.com/hth0919/migrationclient

go 1.13

require (
	github.com/hth0919/migcore v0.0.2
	k8s.io/api v0.18.2 // indirect
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89 // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.18.2