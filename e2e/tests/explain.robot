*** Settings ***
Library             Collections
Library             OperatingSystem
Resource            ../tlm.resource

Suite Setup         Run Command    tlm config set llm.model ${model}
Suite Teardown      Run Command    tlm config set llm.model ${model}

Test Tags           command=explain

Name                tlm explain


*** Variables ***
${model}        qwen2.5-coder:1.5b
${model2}       llama3.2:1b
${style}        balanced


*** Test Cases ***
tlm explain <prompt>
    ${rc}    ${output}=    Run Command    tlm explain "ls -all"
    Should Contain    ${output}    list    ignore_case=True
    Should Contain    ${output}    file    ignore_case=True

tlm explain --model=<model> --style=<style> <prompt>
    ${rc}    ${output}=    Run Command    tlm explain --model=${model2} --style=${style} "ls -all"
    Should Contain    ${output}    list    ignore_case=True
    Should Contain    ${output}    file    ignore_case=True

tlm e <prompt>
    ${rc}    ${output}=    Run Command    tlm e "ls -all"
    Should Contain    ${output}    list    ignore_case=True
    Should Contain    ${output}    file    ignore_case=True
