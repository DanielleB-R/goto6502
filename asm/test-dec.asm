        .include "macros.asm"
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

        place $ab, addr
        dec addr
        dec addr               ; Expect $a9
        place $24, addr+ind
        ldx #ind
        dec addr,X
        dec addr,X             ; Expect $22

;;; DEC Zero Page
        zp = $05
        zind = $04

        place $1f, zp
        dec zp
        dec zp                 ; Expect $1d
        ldx #zind
        sta $06
        dec $06-zind,X         ; Expect $1e
