;;; ASL
        lda #$8f
        ldx #$ff
        asl
        bcc error
        sta $01                 ; Expect $1e

;;; LSR
        lda #$f2
        lsr
        bcs error
        sta $02                 ; Expect $79
        brk

error:  lda #$99
        sta $00
