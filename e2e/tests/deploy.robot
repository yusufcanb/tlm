*** Settings ***
Library     Process
Name        deploy


*** Variables ***
${tlm}      ${EMPTY}


*** Test Cases ***
Should Deploy tlm Modelfiles
    [Tags]    requires=ollama
    ${result}=    Run Process    ${tlm} deploy    shell=True
    Process Should Be Running    ${result}
    ${output}=    Get Process Result    ${result}
    Log    ${output.stdout}

Should Print Error When Ollama is Unreachable
    [Tags]
    ${result}=    Run Process    ./tlm s 'list all directories'    shell=True
    Process Should Be Running    ${result}

    ${output}=    Get Process Result    ${result}
    Log    ${output.stdout}
