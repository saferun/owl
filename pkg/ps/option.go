package ps

type Option func(*ExecuteInfo)

type ExecuteInfo struct {
	File    string
	Command string
	Admin   bool
}

func NewExecuteInfo() *ExecuteInfo {
	return &ExecuteInfo{
		File: "C:\\Windows\\System32\\cmd.exe",
	}
}

func WithFilePath(file string) Option {
	return func(ei *ExecuteInfo) {
		ei.File = file
	}
}

func WithCommand(cmd string) Option {
	return func(ei *ExecuteInfo) {
		ei.Command = cmd
	}
}

func WithAdmin() Option {
	return func(ei *ExecuteInfo) {
		ei.Admin = true
	}
}
