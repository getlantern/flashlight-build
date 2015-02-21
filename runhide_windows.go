package runhide

import (
	"os"
	"os/exec"
	"path"

	"github.com/getlantern/filepersist"
	"github.com/getlantern/golog"
)

var (
	log = golog.LoggerFor("runhide")

	scriptPath = path.Join(os.Getenv("APPDATA"), "runhide.vbs")
)

func init() {
	log.Tracef("Placing script in %s", scriptPath)
	err := filepersist.Save(scriptPath, script, 0644)
	if err != nil {
		panic(err)
	}
}

func command(name string, arg ...string) *exec.Cmd {
	fullArgs := make([]string, 0, len(arg)+2)
	fullArgs = append(fullArgs, scriptPath, name)
	fullArgs = append(fullArgs, arg...)
	return exec.Command("wscript.exe", fullArgs...)
}

var script = []byte(`Option Explicit

Dim i, strArguments, wshShell

If WScript.Arguments.Count = 0 Then Syntax
If WScript.Arguments(0) = "/?" Then Syntax

strArguments = ""

For i = 0 To WScript.Arguments.Count - 1
    strArguments = strArguments & " " & WScript.Arguments(i)
Next

Set wshShell = CreateObject( "WScript.Shell" )
wshShell.Run Trim( strArguments ), 0, False
Set wshShell = Nothing


Sub Syntax
    Dim strMsg
    strMsg = "RunNHide.vbs,  Version 2.00" & vbCrLf _
           & "Run a batch file or (console) command in a hidden window" & vbCrLf _
           & vbCrLf _
           & "Usage:  RUNNHIDE.VBS  some_command  [ some_arguments ]" & vbCrLf _
           & vbCrLf _
           & "Where:  ""some_command""    is the batch file or (console) command" & vbCrLf _
           & "                          you want to run hidden" & vbCrLf _
           & "        ""some_arguments""  are optional arguments for ""some_command""" & vbCrLf _
           & vbCrLf _
           & "Based on a ""one-liner"" by Alistair Johnson" & vbCrLf _
           & "www.microsoft.com/technet/scriptcenter/csc/scripts/scripts/running/cscte009.mspx" _
           & vbCrLf & vbCrLf _
           & "Written by Rob van der Woude" & vbCrLf _
           & "http://www.robvanderwoude.com"
    WScript.Echo strMsg
    WScript.Quit 1
End Sub`)
