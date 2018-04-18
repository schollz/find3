import json
import sys
import mmap

import requests
from tqdm import tqdm


def get_num_lines(file_path):
    fp = open(file_path, "r+")
    buf = mmap.mmap(fp.fileno(), 0)
    lines = 0
    while buf.readline():
        lines += 1
    return lines


server = sys.argv[1]
fname = sys.argv[2]
print(fname)
with open(fname, 'r') as f:
    for line in tqdm(f, total=get_num_lines(fname)):
        r = requests.post(server + "/data?justsave=0", data=line.strip())
