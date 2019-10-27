        .include "macros.asm"
;;; AND immediate
        lda #$f0
        ldx #$00
        ldy #$00
        and #$44
        sta $00                 ; Expect $40

;;; AND absolute
        abs = $5000
        lda #$f0
        ldx #$99
        stx abs
        and abs
        sta $01                 ; Expect $90

;;; AND absolute,X
        xa = $02
        lda #$0f
        ldx #xa
        and abs-xa,X
        sta $02                 ; Expect $09

;;; AND absolute,Y
        ya = $03
        lda #$85
        ldy #ya
        and abs-ya,Y
        sta $03                 ; Expect $81

;;; AND indexed indirect,X
        iiadr = $4000
        iiloc = $c1
        xii = $11

        staddr iiadr, iiloc
        lda #$f1
        sta iiadr
        lda #$c5
        ldx #xii
        and (iiloc-xii,X)
        sta $04                 ; Expect $c1

;;; AND indirect indexed,Y
        yii = $15

        lda #$dd
        sta iiadr+yii
        ldy #yii
        lda #$bb
        and (iiloc),Y
        sta $05                 ; Expect $99

;;; AND zero page
        zp = $cc

        ldx #$11
        stx zp
        lda #$c9
        and zp
        sta $06                 ; Expect $01

;;; AND zero page,X
        zpx = $c
        lda #$d6
        ldx #zpx
        and zp-zpx,X
        sta $07                 ; Expect $10
