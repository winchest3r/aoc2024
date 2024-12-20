import os
import sys

def create_folder(name):
    try:
        os.mkdir(name)
    except FileExistsError:
        print(f"Folder {name} is already exists.")

def create_file(name, content):
    if os.path.exists(name):
        print(f"File {name} is already exists.")
        return
    with open(name, "w", encoding="utf-8") as file:
        file.write(content)

def create_main_content(day_num):
    return f"""package day{day_num}

import "os"
    
func ReadInput(fname string) {{
    file, err := os.Open(fname)
    if err != nil {{
        panic(err)
    }}
    defer file.Close()
}}

func SolvePartOne(fname string) {{

}}

func SolvePartTwo(fname string) {{

}}
"""

def create_test_content(day_num):
    return f"""package day{day_num}

import "testing"

func TestXxx(t *testing.T) {{

}}
"""

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(f"usage: python {sys.argv[0]} day_num")
        sys.exit(1)
    day_num = sys.argv[1]
    day_dir = "day" + day_num
    day_data_dir = day_dir + "/data"
    create_folder(day_dir)
    create_file(day_dir + f"/day{day_num}.go", create_main_content(day_num))
    create_file(day_dir + f"/day{day_num}_test.go", create_test_content(day_num))
    create_folder(day_data_dir)
    create_file(day_data_dir + "/example", "")
    create_file(day_data_dir + "/input", "")