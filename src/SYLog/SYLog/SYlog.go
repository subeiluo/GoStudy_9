package SYLog

const (
	LogTypeToFile =iota
	LogTypeToConsole
)

type SYLog interface {

}

func NewSYlog(logType,level int,filename ,module string)  {
	var logt SYLog
	switch logType {
	case LogTypeToFile:
		logt = level,filename,module
	}
}