

        brk
        ora [0xff,x]
        .db 0x02, 0x03

        tsb 0xff
        ora 0xff
        asl 0xff
        rmb 0, 0xff

        php
        orai 0xff
        asl a
        .db 0x0b
        tsb 0xffff
        ora 0xffff
        asl 0xffff
        bbr 0, 0xff, [-3] ; make this better pls!!!!
lbl1:
        bpl lbl1
        ora [0xff],y
        ora [0xff]
        .db 0x13
        trb 0xff
        ora 0xff,x
        asl 0xff,x
        rmb 1, 0xff

        clc
        ora 0xffff,y
        inc a
        .db 0x1b
        trb 0xffff
        ora 0xffff,x
        asl 0xffff,x
        bbr 1, 0xff, [-3]

        jsr 0xffff
        and [0xff,x]
        .db 0x22, 0x23

        bit 0xff
        and 0xff
        rol 0xff
        rmb 2, 0xff

        plp
        andi 0xff
        rol a
        .db 0x2b
        bit 0xffff
        and 0xffff
        rol 0xffff
        bbr 2, 0xff, [-3]
lbl2:
        bmi lbl2
        and [0xff],y
        and [0xff]
        .db 0x33
        bit 0xff,x
        and 0xff,x
        rol 0xff,x
        rmb 3, 0xff

        sec
        and 0xffff,y
        dec a
        .db 0x3b
        bit 0xffff,x
        and 0xffff,x
        rol 0xffff,x
        bbr 3, 0xff, [-3]

        rti
        eor [0xff,x]
        .db 0x42, 0x43, 0x44


        eor 0xff
        lsr 0xff
        rmb 4, 0xff

        pha
        eori 0xff
        lsr a
        .db 0x4b
        jmp 0xffff
        eor 0xffff
        lsr 0xffff
        bbr 4, 0xff, [-3]
lbl3:
        bvc lbl3
        eor [0xff],y
        eor [0xff]
        .db 0x53, 0x54

        eor 0xff,x
        lsr 0xff,x
        rmb 5, 0xff

        cli
        eor 0xffff,y
        phy
        .db 0x5b, 0x5c

        eor 0xffff,x
        lsr 0xffff,x
        bbr 5, 0xff, [-3]

        rts
        adc [0xff,x]
        .db 0x62, 0x63

        stz 0xff
        adc 0xff
        ror 0xff
        rmb 6, 0xff

        pla
        adci 0xff
        ror a
        .db 0x6b
        jmp [0xffff]
        adc 0xffff
        ror 0xffff
        bbr 6, 0xff, [-3]
lbl4:
        bvs lbl4
        adc [0xff],y
        adc [0xff]
        .db 0x73
        stz 0xff,x
        adc 0xff,x
        ror 0xff,x
        rmb 7, 0xff

        sei
        adc 0xffff,y
        ply
        .db 0x7b
        jmp [0xffff,x]
        adc 0xffff,x
        ror 0xffff,x
        bbr 7, 0xff, [-3]
lbl5:
        bra lbl5
        sta [0xff,x]
        .db 0x82, 0x83

        sty 0xff
        sta 0xff
        stx 0xff
        smb 0, 0xff

        dey
        biti 0xff
        txa
        .db 0x8b
        sty 0xffff
        sta 0xffff
        stx 0xffff
        bbs 0, 0xff, [-3]
lbl6:
        bcc lbl6
        sta [0xff],y
        sta [0xff]
        .db 0x93
        sty 0xff,x
        sta 0xff,x
        stx 0xff,y
        smb 1, 0xff

        tya
        sta 0xffff,y
        txs
        .db 0x9b
        stz 0xffff
        sta 0xffff,x
        stz 0xffff,x
        bbs 1, 0xff, [-3]

        ldyi 0xff
        lda [0xff,x]
        ldxi 0xff
        .db 0xa3
        ldy 0xff
        lda 0xff
        ldx 0xff
        smb 2, 0xff

        tay
        ldai 0xff
        tax
        .db 0xab
        ldy 0xffff
        lda 0xffff
        ldx 0xffff
        bbs 2, 0xff, [-3]
lbl7:
        bcs lbl7
        lda [0xff],y
        lda [0xff]
        .db 0xb3
        ldy 0xff,x
        lda 0xff,x
        ldx 0xff,y
        smb 3, 0xff

        clv
        lda 0xffff,y
        tsx
        .db 0xbb
        ldy 0xffff,x
        lda 0xffff,x
        ldx 0xffff,y
        bbs 3, 0xff, [-3]

        cpyi 0xff
        cmp [0xff,x]
        .db 0xc2, 0xc3

        cpy 0xff
        cmp 0xff
        dec 0xff
        smb 4, 0xff

        iny
        cmpi 0xff
        dex
        .db 0xcb
        cpy 0xffff
        cmp 0xffff
        dec 0xffff
        bbs 4, 0xff, [-3]
lbl8:
        bne lbl8
        cmp [0xff],y
        cmp [0xff]
        .db 0xd3, 0xd4

        cmp 0xff,x
        dec 0xff,x
        smb 5, 0xff

        cld
        cmp 0xffff,y
        phx
        .db 0xdb, 0xdc

        cmp 0xffff,x
        dec 0xffff,x
        bbs 5, 0xff, [-3]

        cpxi 0xff
        sbc [0xff,x]
        .db 0xe2, 0xe3

        cpx 0xff
        sbc 0xff
        inc 0xff
        smb 6, 0xff

        inx
        sbci 0xff
        nop
        .db 0xeb
        cpx 0xffff
        sbc 0xffff
        inc 0xffff
        bbs 6, 0xff, [-3]
lbl9:
        beq lbl9
        sbc [0xff],y
        sbc [0xff]
        .db 0xf3, 0xf4

        sbc 0xff,x
        inc 0xff,x
        smb 7, 0xff

        sed
        sbc 0xffff,y
        plx
        .db 0xfb, 0xfc

        sbc 0xffff,x
        inc 0xffff,x
        bbs 7, 0xff, [-3]

