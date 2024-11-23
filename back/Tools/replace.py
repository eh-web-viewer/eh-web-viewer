import os
import re

# Get the path of the current directory
TOOLS_DIR = os.path.dirname(os.path.abspath(__file__))


def get_module_name():
  """
  Reads the module name from the go.mod file located in the parent directory.

  Returns:
    str: The module name.
  """
  module_file = os.path.join(os.path.dirname(TOOLS_DIR), "go.mod")

  with open(module_file, "r", encoding="utf8") as f:
    first_line = f.readline()
    module_name = first_line.split(" ")[-1].strip()

  return module_name


def replace_go_file(file_path, module_name):
  """
  Replaces the import paths in a Go file to match the current module name.

  Args:
    file_path (str): The path of the Go file to be modified.
    module_name (str): The current module name.
  """
  with open(file_path, "r", encoding="utf8") as f:
    file_content = f.read()

  # Replace import paths with the correct module name
  updated_content = re.sub(
    r'"(.*?)/Tools(.*?)"', rf'"{module_name}/Tools\2"', file_content
  )

  with open(file_path, "w", encoding="utf8") as f:
    f.write(updated_content)


def replace_go_files(module_name):
  """
  Walks through the TOOLS_DIR and modifies all Go files to update their import paths.

  Args:
    module_name (str): The current module name.
  """
  for root, _, files in os.walk(TOOLS_DIR):
    for file in files:
      if file.endswith(".go"):
        file_path = os.path.join(root, file)
        replace_go_file(file_path, module_name)
        print(f"Updated: {file_path}")


if __name__ == "__main__":
  # Get the current module name from go.mod
  module_name = get_module_name()
  print(f"Module name: {module_name}")

  # Replace import paths in all Go files under TOOLS_DIR
  replace_go_files(module_name)
