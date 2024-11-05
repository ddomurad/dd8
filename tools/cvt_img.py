# load png image and read pixel values

from PIL import Image, ImageOps

# im = Image.open('/home/work/tmp/img.png')
im = Image.open('/home/work/downloads/imp.jpeg')
im = ImageOps.posterize(im, 2)
im.save('/home/work/tmp/test_img.png')

pix = im.load()
size = im.size
print(size[0], size[1])


def reverse_bits_in_number(x):
    return int('{:08b}'.format(x)[::-1], 2)


def to_pix(p):
    r1, r2, g1, g2, b1, b2 = 0, 0, 0, 0, 0, 0

    if p[0] == 64:
        r1 = 1
    if p[0] == 128:
        r2 = 1
    if p[0] == 192:
        r1 = 1
        r2 = 1
    if p[1] == 64:
        g1 = 1
    if p[1] == 128:
        g2 = 1
    if p[1] == 192:
        g1 = 1
        g2 = 1
    if p[2] == 64:
        b1 = 1
    if p[2] == 128:
        b2 = 1
    if p[2] == 192:
        b1 = 1
        b2 = 1

    return r1 << 2 | g1 << 3 | b1 << 4 | r2 << 5 | g2 << 6 | b2 << 7


pix_raw = []

for y in range(size[1]):
    for x in range(size[0]):
        pix_raw.append(to_pix(pix[x, y]))


print("pix_raw size: ", len(pix_raw))

with open('/home/work/tmp/VMShared/img.bin', 'wb') as f:
    for i in pix_raw:
        f.write(bytes([i]))


def write_sprite_data(index, data):
    pass
