module github.com/mimetrix/nas/nasTest

go 1.19

replace github.com/mimetrix/nas => /home/luay/projects/mantisnet/nas
replace github.com/mimetrix/nasType => /home/luay/projects/mantisnet/nasType
replace github.com/mimetrix/nasMessage => /home/luay/projects/mantisnet/nasMessage

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/mimetrix/nas v1.0.8-0.20230206222722-e627942c4ad3
	github.com/mimetrix/nas/nasMessage v0.0.0-20230206222722-e627942c4ad3
)

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/mimetrix/nas/nasType v0.0.0-20230206222722-e627942c4ad3 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
