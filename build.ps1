
$maingo=$(Get-ChildItem -Path *.go | Select-String -InputObject {$_.Name} -Pattern \d+\.go)
$maingo -match "^(\d+)\.go$" | Out-Null
$maingo=$($matches[1])

$base=$maingo
$mainexe="${base}.exe"
$mainarc="${base}.a"
$mainsrc="${base}.go"

[string]$bindir="${env:GOROOT}/pkg/tool/${env:GOOS}_${env:GOARCH}"

$filelist=""
Get-ChildItem -Path *.go | Select-String -InputObject {$_.Name} -Pattern \d+\.go -NotMatch | ForEach-Object { $filelist="${filelist} $_"}

Write-Host "compile -pack -o m.a $filelist"
Start-Process -FilePath $bindir/compile.exe -ArgumentList "-pack -o m.a $filelist" -wait

Write-Host "compile -I . -pack -o $mainarc $mainsrc"
Start-Process -FilePath $bindir/compile.exe -ArgumentList "-I . -pack -o $mainarc $mainsrc" -Wait

Write-Host "link -L . -o $mainexe $mainarc"
Start-Process -FilePath $bindir/link.exe -ArgumentList "-L . -o $mainexe $mainarc" -Wait

Remove-Item -Force -Path *.a 2>&1 | Out-Null
