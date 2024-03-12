*** Settings ***
Library     OperatingSystem
Name        suggest


*** Variables ***
${cmd}      tlm suggest "list all network interfaces"


*** Test Cases ***
Should Suggest a Command
    [Tags]  requires=ollama
    ${result}=    Run Process    ./tlm    shell=True
    Process Should Be Running    ${result}
    ${output}=    Get Process Result    ${result}
    Log    ${output.stdout}

Should Print Error When Ollama is Unreachable
    ${cmd}=    Set Variable    ${cmd}
    ${rc}    ${output}=    Run and Return RC and Output    ${cmd}

    Should Be Equal As Integers    ${rc}    255
    Should Contain  ${output}    (err)
    Should Contain  ${output}    Ollama connection failed. Please check your Ollama if it's running or configured correctly.

    Log    ${output}