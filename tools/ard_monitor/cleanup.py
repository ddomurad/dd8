import sys

if len(sys.argv) != 3:
    print("Usage: cleanup.py <input_file> <output_file>")
    exit(1)

FNAME = sys.argv[1]
ONAME = sys.argv[2]

all_lines = open(FNAME, "r").readlines()


def parse_line(line):
    line = line.strip()
    sg = line.split(" ")

    if len(sg) != 4:
        print("Error: ", line)
        exit(1)

    return tuple(sg)


def merge_lines(parsed_lines):
    merged_lines = []
    last_line = ("", "", "", "")
    for line in parsed_lines:
        if line == last_line:
            continue

        merged_lines.append(line)
        last_line = line

    return merged_lines


def validate_lines(lines):
    prev_line = ("", "", "", "")
    for (i, line) in enumerate(lines):
        if prev_line[3] == line[3]:
            print("Error: ", i, prev_line, line)
            exit(1)
        prev_line = line


parsed_lines = [parse_line(line) for line in all_lines]
merged_lines = merge_lines(parsed_lines)
validate_lines(merged_lines)

with open(ONAME, "w") as f:
    for line in merged_lines:
        f.write(" ".join(line) + "\n")
