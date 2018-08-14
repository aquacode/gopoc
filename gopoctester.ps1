param(
    [string]$sleepdurationsec
)

Write-Host "This is a test script"
Write-Host `n
Write-Host $sleepdurationsec
Write-Host `n
Start-Sleep -s $sleepdurationsec
Write-Host "Ending test"
Write-Host `n 
