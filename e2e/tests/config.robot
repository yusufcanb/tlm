*** Settings ***
Library     OperatingSystem
Name        config

*** Test Cases ***
set llm.host <value>
    [Tags]    config

    Set Ollama Host Should Pass w/ Valid Value
    Set Ollama Host Should Fail w/ Invalid Value

set llm.suggest <value>
    [Tags]    config

    ${cmd}=    Set Variable    tlm config set shell auto
    ${rc}    ${output}=    Run and Return RC and Output    ${cmd}

    Should Be Equal As Integers    ${rc}    0
    Should Not Contain  ${output}    (err)
    Should Contain  ${output}    shell = auto

    Log    ${output}

set llm.explain <value>
    [Tags]    config

    ${cmd}=    Set Variable    tlm config set shell auto
    ${rc}    ${output}=    Run and Return RC and Output    ${cmd}

    Should Be Equal As Integers    ${rc}    0
    Should Not Contain  ${output}    (err)
    Should Contain  ${output}    shell = auto

    Log    ${output}

set shell <value>
    [Tags]    config

    ${cmd}=    Set Variable    tlm config set shell auto
    ${rc}    ${output}=    Run and Return RC and Output    ${cmd}

    Should Be Equal As Integers    ${rc}    0
    Should Not Contain  ${output}    (err)
    Should Contain  ${output}    shell = auto

    Log    ${output}

*** Keywords ***

Set Ollama Host Should Fail w/ Invalid Value
    ${cmd}=    Set Variable    tlm config set llm.host invalid-url
    ${rc}    ${output}=    Run and Return RC and Output    ${cmd}

    Should Be Equal As Integers    ${rc}    1
    Should Contain  ${output}    invalid url

    Log    ${output}

Set Ollama Host Should Pass w/ Valid Value
    ${cmd}=    Set Variable    tlm config set llm.host http://ollama:8080
    ${rc}    ${output}=    Run and Return RC and Output    ${cmd}

    Should Be Equal As Integers    ${rc}    0
    Should Not Contain  ${output}    (err)
    Should Contain  ${output}    llm.host = http://ollama:8080

    Log    ${output}