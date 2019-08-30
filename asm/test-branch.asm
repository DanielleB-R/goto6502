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

;;; BPL relative
        ldx #$de
        ldy #$8b
        lda #$ff
        bpl error
        lda #$7a
        bpl plus
        ldx #$ff
        ldy #$ff
plus:   stx $05                 ; Expect $de
        sty $06                 ; Expect $8b

;;; BMI relative
        ldx #$de
        ldy #$8b
        lda #$33
        bmi error
        lda #$f9
        bmi minus
        ldx #$ff
        ldy #$ff
minus:  stx $07                 ; Expect $de
        sty $08                 ; Expect $8b
        brk

;;;

error:  lda #$99
        sta $00
