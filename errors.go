package argparse

// BreakAfterHelp will be thrown after help showed
type BreakAfterHelp struct {
}

func (b BreakAfterHelp) Error() string {
	return "help showed and break"
}

// BreakAfterShellScript will be thrown after shell script showed
type BreakAfterShellScript struct {
}

func (b BreakAfterShellScript) Error() string {
	return "shell script showed and break"
}

// BreakAfterHelpError indicates that is's a break after help call
var BreakAfterHelpError = BreakAfterHelp{}

// BreakAfterShellScriptError indicates that it's a break after shell script call
var BreakAfterShellScriptError = BreakAfterShellScript{}
