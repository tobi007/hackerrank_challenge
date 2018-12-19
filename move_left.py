

x = 5
y = 1

z = [1,2,3,4,5]

def move_left(x,y,z):
    new_arr = []
    for ind in range(x):
        new_ind = (ind + (y % x)) % x
        new_arr.append(z[new_ind])
    
    return new_arr

print(z, y)
print(move_left(x,y,z))