import os

for root, dirs, files in os.walk('./'):
    # no recursion
    if root == './':
        for file in files:
            prefix = file.split('.')[0]
            if not os.path.exists(prefix):
                os.mkdir(prefix)
            os.system(f'move {file} ./{prefix}')