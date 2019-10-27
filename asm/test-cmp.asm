        .include "macros.asm"
;;; CMP immediate
        lda #$99
        cmp #$f1
        bpl error
        beq error
        bcs error
        xplace $90, $01         ; Expect $90

        cmp #$99
        bmi error
        bne error
        bcc error
        xplace $7d, $02         ; Expect $7d

        cmp #$23
        bmi error
        beq error
        bcc error
        xplace $5a, $03         ; Expect $5a

;;; CMP absolute
        addr = $4500

        xplace $c4, addr
        cmp addr
        bpl error
        beq error
        bcs error
        xplace $a6, $04         ; Expect $a6

;;; CMP absolute,X
        xabs = $43

        yplace $99, addr+xabs
        ldx #xabs
        cmp addr,X
        bmi error
        bne error
        bcc error
        xplace $bb, $05         ; Expect $bb

;;; CMP absolute,Y
        yabs = $24

        xplace $1d, addr+yabs
        ldy #yabs
        cmp addr,Y
        bmi error
        beq error
        bcc error
        xplace $da, $06         ; Expect $da

;;; CMP zero page
        zp = $aa

        xplace $c4, zp
        cmp zp
        bpl error
        beq error
        bcs error
        xplace $9e, $07         ; Expect $9e

;;; CMP zero page,X
        zpx = $07

        yplace $99, zp+zpx
        ldx #zpx
        cmp zp,X
        bmi error
        bne error
        bcc error
        xplace $02, $08         ; Expect $02

        jmp past
error:  place $99, $00

;;; CMP indexed indirect
        iiaddr = $5000
        iiloc = $b5
        xii = $06

past:   staddr iiaddr, iiloc
        lda #$99
        yplace $44, iiaddr
        ldx #xii
        cmp (iiloc-xii,X)
        bmi error2
        beq error2
        bcc error2
        yplace $e8, $09         ; Expect $e8

;;; CMP indirect indexed
        yii = $11

        xplace $a3, iiaddr+yii
        ldy #yii
        cmp (iiloc),Y
        bpl error2
        beq error2
        bcs error2
        xplace $23, $0a         ; Expect $23


        brk
error2:  place $99, $00
