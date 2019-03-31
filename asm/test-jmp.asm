        .org $1000
;;; Absolute jump
        lda #$05
        ldx #$20
        ldy #$99
        jmp over
        lda #$50
        ldx #$ff
over:   sta $01                 ; Expect $05
        stx $02                 ; Expect $20

;;; Indirect jump
        dest = $bb00
        lda #<target
        sta dest
        lda #>target
        sta dest+1
        jmp (dest)
        lda #$f3
        jmp save
target: lda #$22
save:   sta $03                 ; Expect $22
