*** Settings ***
Library     OperatingSystem
Library     tlm_lib.py


*** Variables ***
${help}         tlm help
${version}      tlm version
${config}       tlm config
${suggest}      tlm suggest 'list all files in cwd'
${explain}      tlm explain 'ls -all'


*** Keywords ***
Run Command
    [Arguments]    ${command}
    ${rc}    ${output}=    Run and Return RC and Output    ${command}
    Log    Return Code: ${rc}${\n}Output:${\n}${output}

    RETURN    ${rc}    ${output}

Run Help
    ${rc}    ${output}=    Run Command    ${help}
    RETURN    ${rc}    ${output}

Run Version
    ${rc}    ${output}=    Run Command    ${version}
    RETURN    ${rc}    ${output}

Run Suggestion
    ${rc}    ${output}=    Run Command    ${suggest}
    RETURN    ${rc}    ${output}

Run Explain
    ${rc}    ${output}=    Run Command    ${explain}
    RETURN    ${rc}    ${output}

Set Config
    [Arguments]    ${key}    ${value}
    Log    ${config} set ${key} ${value}
    ${rc}    ${output}=    Run Command    ${config} set ${key} ${value}
    RETURN    ${rc}    ${output}

Get Config
    [Arguments]    ${key}
    Log    ${config} get ${key}
    ${rc}    ${output}=    Run Command    ${config} get ${key}
    RETURN    ${rc}    ${output}

Remove Config File
    Remove File    ~/.tlm.yaml
