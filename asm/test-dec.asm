        lda #$00
;;; Decrement X index
        ldx #$a5
        dex
        dex
        stx $01                 ; Expect $a3
        dex
        stx $02                 ; Expect $a2

;;; Decrement Y index
        ldy #$88
        dey
        sty $03                 ; Expect $87
        dey
        dey
        dey
        sty $04                 ; Expect $84

;;; DEC Absolute
        addr = $2020
        ind = $01
        lda #$ab
        sta addr
        dec addr
        dec addr               ; Expect $a9
        lda #$24
        sta addr+ind
        ldx #ind
        dec addr,X
        dec addr,X             ; Expect $22

;;; DEC Zero Page
        zp = $05
        zind = $04
        lda #$1f
        sta zp
        dec zp
        dec zp                 ; Expect $1d
        ldx #zind
        sta $06
        dec $06-zind,X         ; Expect $1e
