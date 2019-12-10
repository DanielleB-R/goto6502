        .include "macros.asm"
        ldx #$ff
        txs

;;; Main program
        place $ff, $00
        place $ff, $01
        place $ff, $02
        jsr sub
        place $ab, $03
        brk

;;; Subroutine
sub:    place $12, $00
        place $34, $01
        place $45, $02
        rts
