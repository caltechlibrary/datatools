<#
generated with CMTools 1.3.5 8109536

.SYNOPSIS
PowerShell script for running pandoc on all Markdown docs ending in .md
#>
$project = "CMTools"
Write-Output "Building website for ${project}"
$pandoc = Get-Command pandoc | Select-Object -ExpandProperty Source

# Get all markdown files except 'nav.md'
$mdPages = Get-ChildItem -Filter *.md | Where-Object { $_.Name -ne "nav.md" }

# Generate HTML page names from markdown files
$htmlPages = $mdPages | ForEach-Object { [System.IO.Path]::ChangeExtension($_.Name, ".html") }

function Build-HtmlPage {
    param($htmlPages, $mdPages)

    foreach ($htmlPage in $htmlPages) {
        $mdPage = [System.IO.Path]::ChangeExtension($htmlPage, ".md")
        if (Test-Path $pandoc) {
            & $pandoc "--metadata" "title=$($htmlPage.Replace('.html', ''))" "-s" "--to" "html5" $mdPage "-o" $htmlPage `
                "--lua-filter=links-to-html.lua" `
                "--lua-filter=add-col-scope.lua" `
                "--template=page.tmpl"
        }

        if ($htmlPage -eq "README.html") {
            Move-Item -Path "README.html" -Destination "index.html" -Force
        }
    }
}

function Invoke-PageFind {
    # Run PageFind
    pagefind --verbose --glob="{*.html,docs/*.html}" --force-language en-US --exclude-selectors="nav,header,footer" --output-path ./pagefind --site .
    git add pagefind
}

# Build HTML page
Build-HtmlPage -htmlPages $htmlPages -mdPages $mdPages

# Invoke PageFind
# Invoke-PageFind

