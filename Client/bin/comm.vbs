Dim shell,command
command = "powershell.exe -nologo -command ./app.exe"
Set shell = CreateObject("WScript.Shell")
shell.Run command,0