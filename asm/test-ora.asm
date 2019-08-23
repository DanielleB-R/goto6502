;;; ORA immediate
        lda #$f0
        ldx #$00
        ldy #$00
        ora #$44
        sta $00                 ; Expect $f4

;;; ORA absolute
        lda #$f0
        ldx #$99
        stx $5000
        ora $5000
        sta $01                 ; Expect $f9

;;; ORA absolute,X
        lda #$0f
        ldx #$02
        ora $4ffe,X
        sta $02                 ; Expect $9f

;;; ORA absolute,Y
        lda #$85
        ldy #$03
        ora $4ffd,Y
        sta $03                 ; Expect $9d
