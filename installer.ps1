#!/usr/bin/env pwsh
# Generated with codemeta-ps1-installer.tmpl, see https://github.com/caltechlibrary/codemeta-pandoc-examples

#
# Set the package name and version to install
#
param(
  [Parameter()]
  [String]$VERSION = "1.3.4"
)
[String]$PKG_VERSION = [Environment]::GetEnvironmentVariable("PKG_VERSION")
if ($PKG_VERSION) {
	$VERSION = "${PKG_VERSION}"
	Write-Output "Using '${PKG_VERSION}' for version value '${VERSION}'"
}

$PACKAGE = "datatools"
$GIT_GROUP = "caltechlibrary"
$RELEASE = "https://github.com/${GIT_GROUP}/${PACKAGE}/releases/tag/v${VERSION}"
$SYSTEM_TYPE = Get-ComputerInfo -Property CsSystemType
if ($SYSTEM_TYPE.CsSystemType.Contains("ARM64")) {
    $MACHINE = "arm64"
} else {
    $MACHINE = "x86_64"
}


# FIGURE OUT Install directory
$BIN_DIR = "${Home}\bin"
Write-Output "${PACKAGE} v${VERSION} will be installed in ${BIN_DIR}"

#
# Figure out what the zip file is named
#
$ZIPFILE = "${PACKAGE}-v${VERSION}-Windows-${MACHINE}.zip"
Write-Output "Fetching Zipfile ${ZIPFILE}"

#
# Check to see if this zip file has been downloaded.
#
$DOWNLOAD_URL = "https://github.com/${GIT_GROUP}/${PACKAGE}/releases/download/v${VERSION}/${ZIPFILE}"
Write-Output "Download URL ${DOWNLOAD_URL}"

if (!(Test-Path $BIN_DIR)) {
  New-Item $BIN_DIR -ItemType Directory | Out-Null
}
curl.exe -Lo "${ZIPFILE}" "${DOWNLOAD_URL}"
#if ([System.IO.File]::Exists($ZIPFILE)) {
if (!(Test-Path $ZIPFILE)) {
    Write-Output "Failed to download ${ZIPFILE} from ${DOWNLOAD_URL}"
} else {
    tar.exe xf "${ZIPFILE}" -C "${Home}"
    #Remove-Item $ZIPFILE

    $User = [System.EnvironmentVariableTarget]::User
    $Path = [System.Environment]::GetEnvironmentVariable('Path', $User)
    if (!(";${Path};".ToLower() -like "*;${BIN_DIR};*".ToLower())) {
        [System.Environment]::SetEnvironmentVariable('Path', "${Path};${BIN_DIR}", $User)
        $Env:Path += ";${BIN_DIR}"
    }
    Write-Output "${PACKAGE} was installed successfully to ${BIN_DIR}"
}
