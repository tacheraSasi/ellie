; -----------------------------------------------------------------
; Ellie CLI NSIS installer
; -----------------------------------------------------------------
!include "MUI2.nsh"
!include "Path.nsh"        ; provides ${PathContains}
!include "StrFunc.nsh"     ; provides ${UnStrReplace}

!define APPNAME "Ellie CLI"
!define COMPANY "Ellie"
!define VERSION "1.0.0"
!define INSTALLDIR "$PROGRAMFILES\\Ellie"
!define EXEFILE "ellie.exe"

Name "${APPNAME} ${VERSION}"
OutFile "Ellie-Setup-${VERSION}.exe"
InstallDir "${INSTALLDIR}"
RequestExecutionLevel admin    ; need admin rights to modify HKLM PATH

; -----------------------------------------------------------------
; Pages
; -----------------------------------------------------------------
!insertmacro MUI_PAGE_DIRECTORY
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_UNPAGE_CONFIRM
!insertmacro MUI_UNPAGE_INSTFILES
!insertmacro MUI_LANGUAGE "English"

; -----------------------------------------------------------------
; Install Section
; -----------------------------------------------------------------
Section "Install Ellie" SEC01
    ; Ensure target dir exists and copy binary
    SetOutPath "$INSTDIR"
    File ".\${EXEFILE}"

    ; ---------------- Add install dir to system PATH ----------------
    ; Read existing PATH into $0
    ReadRegStr $0 HKLM "SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment" "Path"

    ; If $INSTDIR is NOT already in PATH, append it
    ${IfNot} ${PathContains} $0 $INSTDIR
        StrCpy $0 "$0;$INSTDIR"
        WriteRegExpandStr HKLM "SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment" "Path" "$0"
        ; Notify running shells of env-var change
        System::Call 'kernel32::SendMessageTimeoutA(i 0xffff, i ${WM_SETTINGCHANGE}, i 0, t "Environment", i 0, i 5000, *i .r0)'
    ${EndIf}

    ; Add uninstaller entry
    WriteRegStr HKLM "Software\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\${APPNAME}" "DisplayName" "${APPNAME}"
    WriteRegStr HKLM "Software\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\${APPNAME}" "UninstallString" "$INSTDIR\\Uninstall.exe"

    ; Generate uninstaller
    WriteUninstaller "$INSTDIR\\Uninstall.exe"
SectionEnd

; -----------------------------------------------------------------
; Uninstall Section
; -----------------------------------------------------------------
Section "Uninstall"
    Delete "$INSTDIR\\${EXEFILE}"
    Delete "$INSTDIR\\Uninstall.exe"
    RMDir  "$INSTDIR"

    ; Remove install dir from PATH
    ReadRegStr $1 HKLM "SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment" "Path"
    ${UnStrReplace} $1 "$INSTDIR;" ""
    WriteRegExpandStr HKLM "SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment" "Path" "$1"

    ; Remove uninstaller registry key
    DeleteRegKey HKLM "Software\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\${APPNAME}"

    ; Broadcast env-var change again
    System::Call 'kernel32::SendMessageTimeoutA(i 0xffff, i ${WM_SETTINGCHANGE}, i 0, t "Environment", i 0, i 5000, *i .r0)'
SectionEnd
