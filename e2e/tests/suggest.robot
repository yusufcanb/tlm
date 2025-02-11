*** Settings ***
Resource        ../tlm.resource

Suite Setup     Run Command    tlm config set llm.model ${model}

Test Tags       command=suggest

Name            tlm suggest


*** Variables ***
${model}        qwen2.5-coder:1.5b
${model2}       llama3.2:1b
${style}        balanced


*** Test Cases ***
tlm suggest <prompt>
    ${output}=    Run Hanging Command And Verify Output    tlm suggest "list all files"
    Should Contain    ${output}    ${model} is thinking...
    Should Contain    ${output}    ls

    Should Contain    ${output}    Execute
    Should Contain    ${output}    Explain
    Should Contain    ${output}    Cancel

tlm suggest --model=<model> --style=<style> <prompt>
    [Tags]    debug
    ${output}=    Run Hanging Command And Verify Output    tlm suggest --model=${model2} --style=${style} "list all files"
    Should Contain    ${output}    ${model2} is thinking...
    Should Contain    ${output}    ls

    Should Contain    ${output}    Execute
    Should Contain    ${output}    Explain
    Should Contain    ${output}    Cancel

tlm s <prompt>
    ${output}=    Run Hanging Command And Verify Output    tlm s "list all files"
    Should Contain    ${output}    ${model} is thinking...
    Should Contain    ${output}    ls

    Should Contain    ${output}    Execute
    Should Contain    ${output}    Explain
    Should Contain    ${output}    Cancel
