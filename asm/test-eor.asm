        .include "macros.asm"
;;; EOR immediate
        lda #$f0
        ldx #$00
        ldy #$00
        eor #$44
        sta $00                 ; Expect $b4

;;; EOR absolute
        abs = $5000

        place $99, abs
        lda #$f0
        eor abs
        sta $01                 ; Expect $69

;;; EOR absolute,X
        xabs = $02

        lda #$0f
        ldx #xabs
        eor abs-xabs,X
        sta $02                 ; Expect $96

;;; EOR absolute,Y
        yabs = $03

        lda #$85
        ldy #yabs
        eor abs-yabs,Y
        sta $03                 ; Expect $1c

;;; EOR zero page
        zp = $cc

        place $11,zp
        lda #$db
        eor zp
        sta $04                 ; Expect $ca

;;; EOR zero page,X
        zpx = $0c

        lda #$79
        ldx #zpx
        eor zp-zpx,X
        sta $05                 ; Expect $68

;;; EOR indexed indirect,X
        iiaddr = $4000
        iiloc = $c1
        xii = $11

        staddr iiaddr, iiloc
        place $f1,iiaddr
        lda #$c5
        ldx #xii
        eor (iiloc-xii,X)
        sta $06                 ; Expect $f5

;;; EOR indirect indexed,Y
        yii = $15

        place $77, iiaddr+yii
        ldy #yii
        lda #$55
        eor (iiloc),Y
        sta $07                 ; Expect $22
