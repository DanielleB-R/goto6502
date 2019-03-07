;;; Store absolute
        lda #$22
        ldx #$55
        ldy #$bb
        sta $2000               ; Expect $22
        stx $2001               ; Expect $55
        sty $2002               ; Expect $bb

;;; Store absolute index X
        ldx #$10
        sta $2000, X            ; Expect $22 at $2010

;;; Store absolute index Y
        ldy #$20
        sta $2000, Y            ; Expect $22 at $2020
