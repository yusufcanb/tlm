*** Settings ***
Resource        ../tlm.resource

Suite Setup     Run Command    tlm config set llm.model ${model}

Test Tags       command=suggest

Name            tlm suggest


*** Variables ***
${model}    qwen2.5-coder:1.5b


*** Test Cases ***
tlm suggest <prompt>
    ${output}=    Run Suggestion And Verify Output    tlm suggest "list all files"
    Should Contain    ${output}    ${model} is thinking...
    Should Contain    ${output}    ls

    Should Contain    ${output}    Execute
    Should Contain    ${output}    Explain
    Should Contain    ${output}    Cancel

tlm s <prompt>
    ${output}=    Run Suggestion And Verify Output    tlm s "list all files"
    Should Contain    ${output}    ${model} is thinking...
    Should Contain    ${output}    ls

    Should Contain    ${output}    Execute
    Should Contain    ${output}    Explain
    Should Contain    ${output}    Cancel
