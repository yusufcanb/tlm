*** Settings ***
Library     OperatingSystem
Name        deploy


*** Variables ***
${cmd}      tlm deploy


*** Test Cases ***
Should Deploy tlm Modelfiles
    [Tags]    requires=ollama
    ${cmd}=    Set Variable    ${cmd}
    ${rc}    ${output}=    Run and Return RC and Output    ${cmd}

    Should Be Equal As Integers    ${rc}    0
    Should Not Contain  ${output}    (err)

    Log    ${output}

Should Print Error When Ollama is Unreachable
    [Tags]
    ${cmd}=    Set Variable    ${cmd}
    ${rc}    ${output}=    Run and Return RC and Output    ${cmd}

    Should Be Equal As Integers    ${rc}    255
    Should Contain  ${output}    (err)
    Should Contain  ${output}    Ollama connection failed. Please check your Ollama if it's running or configured correctly.

    Log    ${output}
