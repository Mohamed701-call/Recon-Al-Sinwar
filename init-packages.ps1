$root = Get-Location

Get-ChildItem -Path $root -Recurse -Filter *.go | ForEach-Object {

    if ($_.Length -eq 0) {

        $package = $_.Directory.Name.ToLower()

        "package $package" | Set-Content -Path $_.FullName -Encoding UTF8

        Write-Host "[OK] Initialized:" $_.FullName -ForegroundColor Green
    }

}

Write-Host ""
Write-Host "Finished!" -ForegroundColor Cyan