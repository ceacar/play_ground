
max = n

counting_sort_arr = [0] * n + 1


for ele in arr:
    counting_sort_arr[ele] += 1 

def counting_sort():
    for i in range(n+1):
        if counting_sort_arr[i] == 0:
            return i

