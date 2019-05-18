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

;;; BNE relative
        ldx #$9f
        ldy #$82
        lda #$00
        bne error
        lda #$fa
        bne nonz
        ldx #$ff
        ldy #$ff
nonz:   stx $03                 ; Expect $9f
        sty $04                 ; Expect $82
        brk

;;;

error:  lda #$99
        sta $00
