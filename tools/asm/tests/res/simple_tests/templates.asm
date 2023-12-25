.def (
    STR_BEG := 0x10
    RANGE_10 := [0..10] 
  )

.def test_block := (v1) {
  .repeate N := v1 {
    ldai N 
      nop 
      sta STR_BEG + N
  }
}


.tmpl test_block, RANGE_10
.tmpl test_block, [0x10, 0x01, 0x01<<2]
