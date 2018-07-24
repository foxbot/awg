package bot

import (
	"github.com/dop251/goja"
	"github.com/foxbot/awg/commands"
)

func init() {
	evalCommand := commands.Command{
		Aliases: []string{"eval", "debug"},
		Func:    eval,
	}
	Manager.Add(evalCommand)
}

func eval(ctx *commands.Context) commands.Result {
	script := ctx.Argument
	if script == "" {
		return commands.Text("**err:** empty script")
	}

	vm := goja.New()
	vm.Set("ctx", ctx)

	val, err := vm.RunString(script)
	if err != nil {
		return commands.Error(err)
	}

	return commands.Textf("```%s```", val.String())
}
