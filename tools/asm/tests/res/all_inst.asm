

; row 0 
brk
ora [0x01,x]
tsb 0x02
ora 0x03
asl 0x04
asl a
tsb 0x0501
ora 0x0602
asl 0x0603

; row 1 
bpl 100 
ora [0x07],y
ora [0x08]
trb 0x09
ora 0x0a,x 
asl 0x0b,x 
clc 
ora 0x0c01, y
inc a
trb 0x0d02
ora 0x0e03, x 
asl 0x0f04, x 

; row 2 
jsr 0x1000 
and [0x11,x]
bit 0x12 
and 0x13 
rol 0x14 
plp 
andi 0x15 
rol a
bit 0x1501 
and 0x1502 
rol 0x1503 

; row 3 
bmi 100
and [0x16], y
and [0x17]
bit 0x18, x 
and 0x19, x 
rol 0x1a, x 
sec 
and 0x1a01,y
dec a
bit 0x1a02, x
and 0x1a03, x
rol 0x1a04, x

; row 4 
rti 
eor [0x1b, x]
eor 0x1c 
lsr 0x1d 
pha 
eor 0x1e 
lsr a 
jmp 0x1e01 
eor 0x1e02
lsr 0x1e03

; row 5 
bvc 100
eor [0x1f], y
eor [0x20]
eor 0x21, x 
lsr 0x22, x 
cli 
eor 0x2201, y 
phy 
eor 0x2202, x 
lsr 0x2203, x 

; row 6 
rts 
adc [0x23, x]
stz 0x24
adc 0x25 
ror 0x26 
pla 
adci 0x27 
ror a 
jmp [0x2701]
adc 0x2702 
ror 0x2703 

; row 7 
bvs 100
adc [0x28], y 
adc [0x29]
stz 0x2a, x 
adc 0x2b, x 
ror 0x2c, x 
sei 
adc 0x2c01, y 
ply
jmp [0x2c02, x]
adc 0x2c03, x 
ror 0x2c04, x 

; row 8 
bra 100
sta [0x2e, x]
sty 0x2f
sta 0x30
stx 0x31 
dey 
biti 0x32 
txa 
sty 0x3201
sta 0x3202
stx 0x3203

; row 9 
bcc 100
sta [0x33], y 
sta [0x34]
sty 0x35, x 
sta 0x36, x 
stx 0x37, y 
tya 
sta 0x3701, y 
txs 
stz 0x3702 
sta 0x3703, x 
stz 0x3704, x 

; row a 
ldyi 0x38 
lda [0x39, x]
ldxi 0x3a
ldy 0x3b 
lda 0x3c 
ldx 0x3d 
tay 
ldai 0x3e 
tax 
ldy 0x3e01
lda 0x3e02
ldx 0x3e03

; row b 
bcs 200
lda [0x3f], y 
lda [0x40]
ldy 0x41, x
lda 0x42, x
ldx 0x43, y
clv 
lda 0x4301, y 
tsx 
ldy 0x4302, x
lda 0x4303, x
ldx 0x4304, y

; row c 
cpyi 0x44 
cmp [0x45, x]
cpy 0x46 
cmp 0x47 
dec 0x48 
iny 
cmpi 0x49 
dex 
; wai 
cpy 0x4901 
cmp 0x4902 
dec 0x4903 

; row d 
bne 200
cmp [0x4a], y 
cmp [0x4b]
cmp 0x4c,x
dec 0x4d,x
cld 
cmp 0x4d01, y 
phx 
; stp 
cmp 0x4d02, x
dec 0x4d03, x

; row e 
cpxi 0x4e 
sbc [0x4f, x]
cpx 0x50 
sbc 0x51 
inc 0x52 
inx 
sbci 0x53 
nop 
cpx 0x5301 
sbc 0x5302 
inc 0x5303 

; row f 
beq 300
sbc [0x54], y 
sbc [0x55]
sbc 0x56, x 
inc 0x57, x 
sed 
sbc 0x5701, y 
plx 
sbc 0x5702, x
inc 0x5703, x


; bbr 400,0
; bbr 400,1
; bbr 400,2
; bbr 400,3
; bbr 400,4
; bbr 400,5
; bbr 400,6
; bbr 400,7
;
; bbs 400, 0
; bbs 400, 1
; bbs 400, 2
; bbs 400, 3
; bbs 400, 4
; bbs 400, 5
; bbs 400, 6
; bbs 400, 7
;
; rmb 0x05,0
; rmb 0x05,1
; rmb 0x05,2
; rmb 0x05,3
; rmb 0x05,4
; rmb 0x05,5
; rmb 0x05,6
; rmb 0x05,7
;
;
; smb 0x15,0
; smb 0x15,1
; smb 0x15,2
; smb 0x15,3
; smb 0x15,4
; smb 0x15,5
; smb 0x15,6
; smb 0x15,7