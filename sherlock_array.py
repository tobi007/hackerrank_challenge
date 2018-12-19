a = [1,2,3]
b = [1,2,3,3]

def solve1(a):
    for ind in range(len(a)):
        x = sum(a[:ind])
        y = sum(a[ind+1:])
        if x == y:
                return "YES"
    return "NO"

def solve(a,n):
    # Complete this function
    left_sum = 0
    right_sum = sum(a[1:])
    if right_sum == 0:
        return "YES"
    for ind in range(1,n):
        left_sum += a[ind - 1]
        right_sum -= a[ind]
        if left_sum == right_sum:
            return "YES"
    return "NO"