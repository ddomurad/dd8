; mem loader is able to load ram memory from serial port and then jump to code execution 
; this way we can load progams from serial prot 
;
; message format:
; aAbbbbb....x where a is start address low byte, A is the start address high bute, 
; b - this are the data bytes, x - is the fisnish command 
; x - can be: '$' finish memory load, new address and data can be provided. 
;             '!' jump to address, this allows  to load a memory block, and execute it 
;                 or just to execute it 
; a,A and B are expected of 8 bit hex format, eg. 00, 12, ab, a0, 0a, 1c, f4,....
; after the jump, the new routine test controll of the cpu

init_mem_loader: {
  stz LOADER_STATE
  stz BYTES_LOADED
  stz TMP_BYTE
  stz NIBLE_LOAD
  rts 
}

handle_mem_load: {
  { ; switch TMP_VAR_0
    lda TMP_VAR_0
    ; case
    cmpi '$'
    beq load_done
    ; case
    cmpi '!'
    beq execute_code
    ; default
    jmp load_byte
    ; load_byte will return

    load_done:
      jsr init_mem_loader
      rts
    exec_code:
      jmp execute_code
      ; execute_code will return
  }

  load_byte: {
    ; convert two chars like 'fa' into one byte
    jsr cvt_nibble

    sta TMP_VAR_2 
    lda TMP_BYTE 
    asl a
    asl a
    asl a
    asl a
    ora TMP_VAR_2
    sta TMP_BYTE
    ldai 0x01 
    and NIBLE_LOAD
    beq first_nibble
    stz NIBLE_LOAD ; this was the second nibble
    jsr handle_byte
    rts
    first_nibble:
    inc NIBLE_LOAD
    rts
  }

  cvt_nibble: {
    ; convert a char into a nibble 
    cmpi '0'
    bcs step_1 ; >= '0'
    rts ; < '0'

    step_1:
    cmpi ':'
    bcs step_2 ; >= ':'
    jmp cvt_0_9 ; < ':'

    step_2:
    cmpi 'a'
    bcs step_3 ; >= 'a'
    rts

    step_3:
    cmpi 'g'
    bcs ret ; >= 'g'
    jmp cvt_a_f ; < 'g'
    
    ret:
    rts
    cvt_0_9: { 
      ; convert '0' to '9' into 0 to 9
      clc
      ldai 0x100 - '0'
      adc TMP_VAR_0
      rts
    }
    cvt_a_f: {
      ; convert 'a' to 'f' into 10 to 15
      clc
      ldai 0x100 - 'W' 
      adc TMP_VAR_0
      rts
    }
  }

  handle_byte: {
    lda LOADER_STATE
    cmpi 0x00
    beq handle_addr_1
    cmpi 0x01 
    beq handle_addr_2
    cmpi 0x02 
    beq handle_data

    handle_addr_1: {
      inc LOADER_STATE
      lda TMP_BYTE 
      sta LOAD_ADDR+1
      rts
    }

    handle_addr_2: {
      inc LOADER_STATE
      lda TMP_BYTE 
      sta LOAD_ADDR
      rts
    }

    handle_data:{
      lda TMP_BYTE
      ldy BYTES_LOADED
      sta [LOAD_ADDR],y
      inc BYTES_LOADED 
      rts
    }
  }

  execute_code: {
    ldxi 0x00 
    txs 

    lda LOAD_ADDR
    bne skip 
    dec LOAD_ADDR+1
  skip: 
    dec LOAD_ADDR

    lda LOAD_ADDR+1
    pha 
    lda LOAD_ADDR 
    pha 
    rts ; this should return to the new address
  }
}

