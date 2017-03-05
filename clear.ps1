
Get-ChildItem -Recurse -Path . -Include "*.exe" | ForEach-Object {Remove-Item -Force -Path "$_" 2>&1 | Out-Null}
