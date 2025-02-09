*** Settings ***
Library         Collections
Resource        tlm.resource
Suite Setup     Log Variables

*** Variables ***
${OLLAMA_ENABLED}       ${TRUE}


*** Test Cases ***
help
    tlm help

version
    tlm version

config
    [Setup]    Remove Config File

    # Shell
    tlm config set shell invalid (negative)
    @{shells}=    Create List    bash    powershell    auto    zsh
    FOR    ${value}    IN    @{shells}
        tlm config set shell ${value} (positive)
        tlm config get shell (positive)    ${value}    # Check if the value is set
    END

    @{prefs}=    Create List    stable    balanced    creative
    # Explain Preference
    tlm config set llm.explain invalid-pref (negative)
    FOR    ${value}    IN    @{prefs}
        tlm config set llm.explain ${value} (positive)
        tlm config get llm.explain (positive)    ${value}    # Check if the value is set
    END

    # Suggest Preference
    tlm config set llm.suggest invalid-pref (negative)
    FOR    ${value}    IN    @{prefs}
        tlm config set llm.suggest ${value} (positive)
        tlm config get llm.suggest (positive)    ${value}    # Check if the value is set
    END
    [Teardown]    Remove Config File

suggest (p)
    [Tags]    suggest    requires=ollama
    tlm suggest 'list all hidden files in cwd' (positive)

suggest (n)
    [Tags]    suggest
    tlm suggest 'list all hidden files in cwd' (negative)

explain (p)
    [Tags]    explain    requires=ollama
    tlm explain 'ls -all' (positive)

explain (n)
    [Tags]    explain
    tlm explain 'ls -all' (negative)


*** Keywords ***
# ------ Config ------

tlm config set ${key} ${value} (positive)
    ${rc}    ${output}=    Set Config    ${key}    ${value}

    Should Be Equal As Integers    ${rc}    0
    Should Contain    ${output}    ${key} = ${value}
    Should Contain    ${output}    (ok)
    Should Not Contain    ${output}    (err)

tlm config set ${key} ${value} (negative)
    ${rc}    ${output}=    Set Config    ${key}    ${value}
    Should Not Be Equal As Integers    ${rc}    0

tlm config get ${key} (positive)
    [Arguments]    ${value}
    ${rc}    ${output}=    Get Config    ${key}

    Should Be Equal As Integers    ${rc}    0
    Should Contain    ${output}    ${key} = ${value}

# ------ Version & Help ------

tlm help
    ${rc}    ${output}=    Run Help
    Should Be Equal As Integers    ${rc}    0
    Should Contain    ${output}    NAME:
    Should Contain    ${output}    USAGE:

    Should Contain    ${output}    VERSION:

    Should Contain    ${output}    COMMANDS:
    Should Contain    ${output}    ask, a      Asks a question
    Should Contain    ${output}    suggest, s  Suggests a command.
    Should Contain    ${output}    explain, e  Explains a command.
    Should Contain    ${output}    config, c   Configures language model, style and shell
    Should Contain    ${output}    version, v  Prints tlm version.

tlm version
    ${rc}    ${output}=    Run Version
    Should Be Equal As Integers    ${rc}    0
    Version Should Be Correct    ${output}


# ------ Suggest --------

tlm suggest '${prompt}' (positive)
    ${rc}    ${output}=    Run Command    tlm suggest '${prompt}'

    Should Be Equal As Integers    ${rc}    0

tlm suggest '${prompt}' (negative)
    ${rc}    ${output}=    Run Command    tlm suggest '${prompt}'

    Should Be Equal As Integers    ${rc}    255
    Should Contain    ${output}    (err)
    Should Contain
    ...    ${output}
    ...    Ollama connection failed. Please check your Ollama if it's running or configured correctly.

# ------ Explain --------

tlm explain '${prompt}' (positive)
    ${rc}    ${output}=    Run Command    tlm explain '${prompt}'

    Should Be Equal As Integers    ${rc}    0

tlm explain '${prompt}' (negative)
    ${rc}    ${output}=    Run Command    tlm explain '${prompt}'

    Should Be Equal As Integers    ${rc}    255
    Should Contain    ${output}    (err)
    Should Contain
    ...    ${output}
    ...    Ollama connection failed. Please check your Ollama if it's running or configured correctly.
