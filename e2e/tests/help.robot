*** Settings ***
Library     Collections
Resource    ../tlm.resource

Name        tlm help


*** Test Cases ***
tlm help
    [Tags]    command=help
    ${rc}    ${output}=    Run Command    tlm help
    Verify Help Command Output    ${rc}    ${output}


*** Keywords ***
Verify Help Command Output
    [Arguments]    ${rc}    ${output}
    Should Be Equal As Integers    ${rc}    0
    Should Contain    ${output}    NAME:
    Should Contain    ${output}    USAGE:

    Should Contain    ${output}    VERSION:

    Should Contain    ${output}    COMMANDS:
    Should Contain    ${output}    ask, a    Asks a question
    Should Contain    ${output}    suggest, s    Suggests a command.
    Should Contain    ${output}    explain, e    Explains a command.
    Should Contain    ${output}    config, c    Configures language model, style and shell
    Should Contain    ${output}    version, v    Prints tlm version.
