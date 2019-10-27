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
