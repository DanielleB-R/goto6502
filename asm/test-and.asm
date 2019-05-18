;;; AND immediate
        lda #$f0
        ldx #$00
        ldy #$00
        and #$44
        sta $00                 ; Expect $40
