*** Settings ***
Library             Collections
Library             OperatingSystem
Resource            ../tlm.resource

Suite Setup         Run Command    tlm config set llm.model ${model}
Suite Teardown      Run Command    tlm config set llm.model ${model}

Test Tags           command=ask

Name                tlm ask


*** Variables ***
${model}    qwen2.5-coder:3b


*** Test Cases ***
tlm ask <command>
    ${rc}    ${output}=    Run Command    tlm ask "Why the sky is blue? Name the concept."

    Should Be Equal As Integers    ${rc}    0
    Should Contain    ${output}    Rayleigh scatte

    ${rc}    ${output}=    Run Command    tlm a "Why the sky is blue? Name the concept."

    Should Be Equal As Integers    ${rc}    0
    Should Contain    ${output}    Rayleigh scattering

tlm ask --context=<context> --include=<patterns> <command>
    ${rc}    ${output}=    Run Command    tlm ask --context=. --include=**/*.robot "explain provided context"
    ${expected_file_list}=    Create List    tests/ask.robot    tests/suggest.robot    tests/help.robot

    Verify Ask Command Output With Context
    ...    ${rc}
    ...    ${output}
    ...    ${expected_file_list}

tlm ask --context=<context> --exclude=<patterns> <command>
    ${rc}    ${output}=    Run Command    tlm ask --context=. --exclude=**/*.robot "explain provided context"
    ${expected_file_list}=    Create List    tlm.robot    tlm.resource    tlm_lib.py    requirements.txt
    Verify Ask Command Output With Context
    ...    ${rc}
    ...    ${output}
    ...    ${expected_file_list}

tlm ask (no ollama)
    [Tags]    no-ollama

    # Test that the command fails when OLLAMA_HOST is not set
    Remove Environment Variable    OLLAMA_HOST
    ${rc}    ${output}=    Run Command    tlm ask "What is the meaning of life?"
    Should Not Be Equal As Integers    ${rc}    0
    Should Contain    ${output}    (err)
    Should Contain
    ...    ${output}
    ...    OLLAMA_HOST environment variable is not set

    # Test the command fails when OLLAMA_HOST is set but not reachable
    Set Environment Variable    OLLAMA_HOST    http://localhost:11434
    ${rc}    ${output}=    Run Command    tlm ask "What is the meaning of life?"

    Should Not Be Equal As Integers    ${rc}    0
    Should Contain    ${output}    (err)
    Should Contain
    ...    ${output}
    ...    Ollama connection failed. Please check your Ollama if it's running or configured correctly.

tlm ask (non-exist model)
    ${model}=    Set Variable    non-exist-model:1b
    Run Command    tlm config set llm.model ${model}

    ${rc}    ${output}=    Run Command    tlm ask 'What is the meaning of life?'
    Should Not Be Equal As Integers    ${rc}    0
    Should Contain    ${output}    model "${model}" not found, try pulling it first


*** Keywords ***
Verify Ask Command Output With Context
    [Arguments]    ${rc}    ${output}    ${expected_file_list}

    Should Be Equal As Numbers    ${rc}    0

    FOR    ${file}    IN    @{expected_file_list}
        Should Contain    ${output}    ${file}
    END

    Should Contain    ${output}    Context Summary:
    Should Contain    ${output}    Total Files:
    Should Contain    ${output}    Total Chars:
    Should Contain    ${output}    Total Tokens:
