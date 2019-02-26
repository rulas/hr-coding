#!/bin/python3

import os

# Complete the freqQuery function below.
def freqQuery(queries):
    data = {}
    result = []
    for query in queries:
        action, value = query
        if action == 1:
            print("add new elements")
            data[value] = data.get(value, 0) + 1
        elif  action == 2:
            data[value] = data.get(value,1) - 1
        elif action == 3:
            print(data)
            found = any([True for v in data.values() if v == value])
            result.append("1" if found else "0")

    return result


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input().strip())

    queries = []

    for _ in range(q):
        queries.append(list(map(int, input().rstrip().split())))

    ans = freqQuery(queries)

    fptr.write('\n'.join(map(str, ans)))
    fptr.write('\n')

    fptr.close()
