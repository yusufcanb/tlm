import subprocess
import time
import threading
import queue

from robot.api.deco import keyword


@keyword("Version Should Be Correct")
def check_version(version: str):
    import re

    pattern = r"^tlm\s\d+\.\d+(?:-pre)?\s\([0-9a-f]+\)\son\s\w+/\w+$"
    if re.match(pattern, version):
        return True
    else:
        raise AssertionError(f"Version does not match regex pattern '{pattern}'")


@keyword("Run Hanging Command And Verify Output")
def test_hanging_command_output(command) -> str:
    """
    Executes a command that produces output and then waits for user input
    (i.e. it hangs until input is provided). The function waits a few seconds
    for the command to produce its initial output, checks that the expected
    output is present, and asserts accordingly.

    Note:
      This function does NOT send any input to finish the process.
      After the assertion, the process will continue to wait for input.

    Parameters:
      command (str): The shell command to run.
      expected_output (str): A substring expected to be present in the output.

    Returns:
      proc (subprocess.Popen): The process object, still running and waiting for input.
    """
    # Start the process.
    proc = subprocess.Popen(
        command,
        shell=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        stdin=subprocess.PIPE,
        text=True,  # Work with strings rather than bytes.
        bufsize=1,  # Use line-buffering.
    )

    # A thread-safe queue to hold the output lines.
    output_queue = queue.Queue()

    def enqueue_output(out, queue_obj):
        """
        Reads lines from the process's stdout and places them into a queue.
        """
        for line in iter(out.readline, ""):
            queue_obj.put(line)
        out.close()

    # Start a daemon thread to continuously read from proc.stdout.
    t = threading.Thread(target=enqueue_output, args=(proc.stdout, output_queue))
    t.daemon = True  # Ensures the thread won't block program exit.
    t.start()

    # Wait 2-3 seconds for the process to generate some output.
    time.sleep(3)

    # Drain the queue to collect the output produced so far.
    output_lines = []
    while True:
        try:
            line = output_queue.get_nowait()
        except queue.Empty:
            break
        else:
            output_lines.append(line)

    partial_output = "".join(output_lines)
    return partial_output
