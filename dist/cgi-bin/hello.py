#!/usr/bin/env python3

import sys
import os

print('Content-type: text/html\n\n')
print('<h1>Hello, World.</h1>')

print('<body>')

if not sys.stdin.isatty():
    print("from stdIn:")
    for line in sys.stdin:
        print(f'{line}')

print("from ENV:")
for key, value in os.environ.items():
    print(key, value, "<br>")

print('</body>')