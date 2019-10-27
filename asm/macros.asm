        .macro staddr addr,loc
        lda #<addr
        sta loc
        lda #>addr
        sta loc+1
        .endmacro

        .macro place val,addr
        lda #val
        sta addr
        .endmacro

        .macro xplace val,addr
        ldx #val
        stx addr
        .endmacro

        .macro yplace val,addr
        ldy #val
        sty addr
        .endmacro
