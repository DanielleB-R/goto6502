        .include "macros.asm"
;;; ORA immediate
        lda #$f0
        ldx #$00
        ldy #$00
        ora #$44
        sta $00                 ; Expect $f4

;;; ORA absolute
        addr = $5000

        place $99, addr
        lda #$f0
        ora addr
        sta $01                 ; Expect $f9

;;; ORA absolute,X
        absx = $02

        lda #$0f
        ldx #absx
        ora addr-absx,X
        sta $02                 ; Expect $9f

;;; ORA absolute,Y
        absy = $03

        lda #$85
        ldy #absy
        ora addr-absy,Y
        sta $03                 ; Expect $9d

;;; ORA zero page
        zp = $cc

        xplace $11,zp
        lda #$c8
        ora zp
        sta $04                 ; Expect $d9

;;; ORA zero page,X
        zpx = $0c

        lda #$00
        ldx #zpx
        ora zp-zpx,X
        sta $05                 ; Expect $11

;;; ORA indexed indirect,X
        iiaddr = $4000
        iiloc = $c1
        xii = $11

        staddr iiaddr,iiloc
        place $f1,iiaddr
        lda #$c5
        ldx #xii
        ora (iiloc-xii,X)
        sta $06                 ; Expect $f5

;;; ORA indirect indexed,Y
        yii = $15

        place $88,iiaddr+yii
        ldy #yii
        lda #$55
        ora (iiloc),Y
        sta $07                 ; Expect $dd
