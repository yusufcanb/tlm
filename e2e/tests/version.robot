*** Settings ***
Library     Collections
Resource    ../tlm.resource

Name        tlm version


*** Test Cases ***
tlm version
    ${rc}    ${output}=    Run Version
    Should Be Equal As Integers    ${rc}    0
    Version Should Be Correct    ${output}
