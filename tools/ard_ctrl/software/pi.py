total = 4
sign = -1
div = 3

for _ in range(1000):
    total += (4/div)*sign
    sign *= -1
    div += 2
    print(total)
