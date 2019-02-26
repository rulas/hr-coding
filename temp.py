#!/bin/python3

import os

inputtext = """\
8
1 5
1 6
3 2
1 10
1 10
1 6
2 5
3 2"""

# Complete the freqQuery function below.
def freqQuery(queries):
    data = {}
    result = []
    freq = {0:0}
    
    for query in queries:
        action, value = query
        action, value = int(action), int(value)
        if action == 1:
            newcount = data.get(value, 0) + 1
            data[value] = newcount
            
            freq[newcount]  = freq.get(newcount, 0) + 1
            freq[newcount-1]  -= 1
        elif  action == 2:
            if data.get(value, 0) > 0:
                # decrement for that number, and store
                newcount = data[value] - 1
                data[value] = newcount
        
                # substract 1 to freq of newcount
                # increment by 1 to freq of newcount-1
                
                freq[newcount+1]  = freq[newcount+1] - 1 
                freq[newcount]  = freq.get(newcount, 0) + 1                
        elif action == 3:
            result.append("1" if freq.get(value, 0) > 0 else "0")
        
        print(data)    
        print(freq)

    return result


if __name__ == '__main__':
    fptr = open('output.txt', 'w')

    lines = inputtext.splitlines()
    q = lines[0].strip()

    queries = []

    for line in lines[1:]:
        queries.append([int(v) for v in line.split()])

    ans = freqQuery(queries)

    fptr.write('\n'.join(map(str, ans)))
    fptr.write('\n')

    fptr.close()
