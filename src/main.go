package main

import (
	"cli"
	"fmt"
	"os"
	"qiniu/log"
	"qiniu/rpc"
	"runtime"
)

var debugMode = false

var supportedCmds = map[string]cli.CliFunc{
	"qupload": cli.QiniuUpload,
}

func main() {
	//set cpu count
	runtime.GOMAXPROCS(runtime.NumCPU())
	//set qshell user agent
	rpc.UserAgent = cli.UserAgent()

	//parse command
	args := os.Args
	argc := len(args)
	log.SetOutputLevel(log.Linfo)
	log.SetOutput(os.Stdout)
	if argc > 1 {
		cmd := ""
		params := []string{}
		option := args[1]
		if option == "-d" {
			if argc > 2 {
				cmd = args[2]
				if argc > 3 {
					params = args[3:]
				}
			}
			log.SetOutputLevel(log.Ldebug)
		} else if option == "-v" {
			cli.Version()
			return
		} else if option == "-h" {
			cli.Help("help")
			return
		} else {
			cmd = args[1]
			if argc > 2 {
				params = args[2:]
			}
		}
		if cmd == "" {
			fmt.Println("Error: no subcommand specified")
			return
		}

		if cliFunc, ok := supportedCmds[cmd]; ok {
			r := cliFunc(cmd, params...)
			if r == 1 {
				fmt.Println(fmt.Sprintf("exit code %d", r))
				os.Exit(r)
			}
		} else {
			fmt.Println(fmt.Sprintf("Error: unknown cmd `%s'", cmd))
		}
	} else {
		fmt.Println("Use help or help [cmd1 [cmd2 [cmd3 ...]]] to see supported commands.")
	}
}
