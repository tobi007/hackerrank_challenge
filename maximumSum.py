class MaximunDS:
    def __init__(self, arr):
        self.level = 0
        self.searchMap = []
        self.generateSearchMap(arr)
        # self.searchMap.append(arr[:])
        # self.searchMap.append(arr[:])

    def generateSearchMap(self, arr):
        arrLen = len(arr)
        for ind in range(arrLen):
            x = arrLen - ind
            self.searchMap.append(arrLen[:])
    
    def print_searchMap(self):
        print(self.searchMap)

a = [3,3,9,9,5]
#a = [1,2,3]
x = len(a)
sumCol = [0] * x
modArr = [0] * x

maxi = 0
for i in range(x):
    for j in range(x - i):
        sumCol[j] += a[j+i]
        modArr[j] = (sumCol[j] % 7)
    print(modArr)
    maxiMod = max(modArr)
    modArr.pop()
    sumCol.pop()
    if maxiMod > maxi:
        maxi = maxiMod
    if maxi == (7 -1):
        print(maxi)
        exit()
    
#arrrr = MaximunDS(a)

print(maxi)