from robot.api.deco import keyword


@keyword("Version Should Be Correct")
def check_version(version: str):
    import re
    pattern = r'^tlm\s\d+\.\d+\s\([0-9a-f]+\)\son\s\w+/\w+$'
    if re.match(pattern, version):
        return True
    else:
        raise AssertionError(f"Version does not match regex pattern '{pattern}'")


@keyword("Run Suggestion And Wait For Form")
def run_suggestion_and_wait_for_form():
    print("Opening tlm...")


@keyword("Run Config And Wait For Form")
def run_config_and_wait_for_form():
    print("Opening tlm...")
