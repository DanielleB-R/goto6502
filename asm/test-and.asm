;;; AND immediate
        lda #$f0
        ldx #$00
        ldy #$00
        and #$44
        sta $00                 ; Expect $40

;;; AND absolute
        lda #$f0
        ldx #$99
        stx $5000
        and $5000
        sta $01                 ; Expect $90

;;; AND absolute,X
        lda #$0f
        ldx #$02
        and $4ffe,X
        sta $02                 ; Expect $09

;;; AND absolute,Y
        lda #$85
        ldy #$03
        and $4ffd,Y
        sta $03                 ; Expect $81
