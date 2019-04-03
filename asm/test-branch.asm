;;; BEQ relative
        lda #$55
        ldx #$28
        ldy #$cd
        beq error
        lda #$00
        beq zero
        ldx #$ff
        ldy #$ff
zero:   stx $01                 ; Expect $28
        sty $02                 ; Expect $cd
        brk

error:  lda #$99
        sta $00
