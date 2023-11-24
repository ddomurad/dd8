import math

Fb = 50_000
# Fb = 125_000
Tb = 1/Fb

S_5k = Fb/5_000
S_1k = Fb/1_000
S_500 = Fb/500
S_100 = Fb/100
S_50 = Fb/50

print("5k ", S_5k)
print("1k ", S_1k)
print("500 ", S_500)
print("100 ", S_100, math.ceil(S_100/255))
print("50 ", S_50, math.ceil(S_50/255))
