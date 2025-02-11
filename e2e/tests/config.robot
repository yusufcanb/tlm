*** Settings ***
Library         Collections
Library         OperatingSystem
Resource        ../tlm.resource

Test Setup      Remove Config File

Test Tags       command=explain

Name            tlm explain


*** Variables ***
${model}    qwen2.5-coder:1.5b
${style}    balanced


*** Test Cases ***
tlm config
    ${rc}    ${output}=    Run Hanging Command And Verify Output    tlm config "ls -all"
    Should Contain    ${output}
    ...    Sets a default model from the list of all available models.
    ...    Use `ollama pull <model_name>` to download new models.

    Should Contain    ${output}
    ...    Sets a default model from the list of all available models.
    ...    Use `ollama pull <model_name>` to download new models.

tlm config ls
    ${rc}    ${output}=    Run Command    tlm config ls

tlm config set <key> <value>
    ${rc}    ${output}=    Run Command    tlm config set llm.model ${model}

tlm config get <key>
    ${rc}    ${output}=    Run Command    tlm config get llm.model


*** Keywords ***
Remove Config File
    ${HOME_DIR}=    Get Environment Variable    HOME
    Remove File    path=${HOME_DIR}/.tlm.yml
