package utils

import (
	"github.com/urfave/cli"
)

var (
	ChainIDFlag = cli.IntFlag{
		Name:  "chainID",
		Usage: "chainID to be stored in Redis",
		Value: 1,
	}
	DataportFlag = cli.IntFlag{
		Name:  "dataport, dp",
		Usage: "Data port for Redis",
		Value: 6379,
	}
	DataipFlag = cli.StringFlag{
		Name:  "dataip, di",
		Usage: "Database ip address",
		Value: "localhost",
	}
	DatabaseFlag = cli.IntFlag{
		Name:  "database, db",
		Usage: "Number of database for Redis",
		Value: 0,
	}
	ListenPortFlag = cli.IntFlag{
		Name:  "port, p",
		Usage: "Network listening port",
		Value: 1423,
	}
	DbPasswdPortFlag = cli.StringFlag{
		Name:  "passwd, pw",
		Usage: "Redis password",
		Value: "",
	}
	PassPhraseFlag = cli.StringFlag{
		Name:  "passphrase, ph",
		Usage: "Used to generate public and private key",
		Value: "",
	}
)

// MigrateFlags sets the global flag from a local flag when it's set.
// This is a temporary function used for migrating old command/flags to the
// new format.
//
// e.g. geth account new --keystore /tmp/mykeystore --lightkdf
//
// is equivalent after calling this method with:
//
// geth --keystore /tmp/mykeystore --lightkdf account new
//
// This allows the use of the existing configuration functionality.
// When all flags are migrated this function can be removed and the existing
// configuration functionality must be changed that is uses local flags
func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				ctx.GlobalSet(name, ctx.String(name))
			}
		}
		return action(ctx)
	}
}
