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
        lda #$ab
        sta addr
        dec addr
        dec addr               ; Expect $a9
        lda #$24
        sta addr+1
        ldx #$01
        dec addr,X
        dec addr,X             ; Expect $22

;;; DEC Zero Page
        lda #$1f
        sta $05
        dec $05
        dec $05                 ; Expect $1d
        ldx #$04
        sta $06
        dec $02,X               ; Expect $1e