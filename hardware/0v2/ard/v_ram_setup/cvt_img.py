# load png image and read pixel values

from PIL import Image

im = Image.open('/home/work/tmp/img.png')
pix = im.load()
size = im.size
print(size[0], size[1])


def to_pix(p):
    r, g, b = 0, 0, 0

    if p[0] != 0:
        r = 1
    if p[1] != 0:
        g = 1
    if p[2] != 0:
        b = 1

    return r | g << 1 | b << 2


pix_raw = []

for y in range(size[1]):
    for x in range(size[0]):
        pix_raw.append(to_pix(pix[x, y]))


print("pix_raw size: ", len(pix_raw))

pix_out = []
for i in range(len(pix_raw)//2):
    pix_out.append(pix_raw[i*2] | pix_raw[i*2+1] << 4)

print("pix_out size: ", len(pix_out))
# save pix_out as c header file
f = open('./img.h', 'w')

f.write('#include <avr/pgmspace.h>\n\n')
f.write('const unsigned char img[] PROGMEM = {\n')

for i in range(len(pix_out)):
    f.write('0x%02x, ' % pix_out[i])
    if i % 16 == 15:
        f.write('\n')

f.write('};\n')
