module github.com/wostzone/directory

go 1.14

require (
	github.com/sirupsen/logrus v1.8.0
	github.com/wostzone/hub v0.0.0-20210227062304-ae0ec41fc4b7
)

// Until Hub is stable
replace github.com/wostzone/hub => ../hub
