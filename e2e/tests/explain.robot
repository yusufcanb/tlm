*** Settings ***
Library     OperatingSystem
Name        explain


*** Variables ***
${cmd}      tlm explain "ls -all"

*** Test Cases ***
Should Explain Given Command
    [Tags]   requires=ollama
    ${cmd}=    Set Variable    ${cmd}
    ${rc}    ${output}=    Run and Return RC and Output    ${cmd}

    Should Be Equal As Integers    ${rc}    0
    Should Not Contain  ${output}    (err)
    Should Contain  ${output}    llm.host = http://ollama:8080

    Log    ${output}

Should Print Error When Ollama is Unreachable
    ${cmd}=    Set Variable    ${cmd}
    ${rc}    ${output}=    Run and Return RC and Output    ${cmd}

    Should Be Equal As Integers    ${rc}    255
    Should Contain  ${output}    (err)
    Should Contain  ${output}    Ollama connection failed. Please check your Ollama if it's running or configured correctly.

    Log    ${output}