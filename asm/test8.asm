        lda #$01
        ldx #$10
        ldy #$22
        sta $00
        stx $01
        sty $02
        lda $f2,X
        ldy $f1,X
