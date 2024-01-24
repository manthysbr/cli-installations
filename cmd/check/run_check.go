package check

import (
    "fmt"
    "manthys/cmd/validate"
    "os/exec"
    "strings"
)

func RunCheck() {
    // Chama a função de checagem para Python
    CheckPython()
    CheckGit()
    CheckDocker()
    CheckProxy()
    CheckAzureCli()
    // Chama funções de checagem para outros softwares    // Exemplo: CheckGit(), CheckDocker(), etc.
}