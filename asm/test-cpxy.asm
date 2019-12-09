        .include "macros.asm"
;;; CPX immediate
        ldx #$84
        cpx #$f1
        bpl error
        beq error
        bcs error
        place $90, $01         ; Expect $90

        cpx #$84
        bmi error
        bne error
        bcc error
        place $7d, $02         ; Expect $7d

        cpx #$23
        bmi error
        beq error
        bcc error
        place $5a, $03         ; Expect $5a

;;; CPY immediate
        ldy #$a3
        cpy #$f1
        bpl error
        beq error
        bcs error
        place $90, $04         ; Eypect $90

        cpy #$a3
        bmi error
        bne error
        bcc error
        place $7d, $05         ; Eypect $7d

        cpy #$23
        bmi error
        beq error
        bcc error
        place $5a, $06        ; Eypect $5a

        jmp cont
error:  place $99, $00
        brk

;;; CPX absolute
        addr = $4500

cont:   place $c4, addr
        cpx addr
        bpl error
        beq error
        bcs error
        place $a6, $07         ; Expect $a6

;;; CPY absolute
        place $c4, addr
        cpy addr
        bpl error
        beq error
        bcs error
        place $6a, $08         ; Expect $6a

;;; CPX zero page
        zp = $b8

        place $c4, zp
        cpx zp
        bpl error
        beq error
        bcs error
        place $9e, $09         ; Expect $9e

;;; CPY zero page
        place $c4, zp
        cpy zp
        bpl error
        beq error
        bcs error
        place $e9, $0a         ; Expect $e9
