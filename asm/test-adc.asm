        .include "macros.asm"
;;; ADC immediate
        clc
        lda #$22
        ldx #$00
        ldy #$00
        adc #$8d
        bcs error
        sta $01                 ; Expect $af

;;; ADC absolute
        addr = $4000

        place $88,addr
        lda #$cd
        adc addr
        bcc error
        sta $02                 ; Expect $55

;;; ADC absolute,X
        absx = $96

        place $73,addr+absx
        ldx #absx
        lda #$d7
        adc addr,X
        bcc error
        sta $03                 ; Expect $4b

;;; ADC absolute,Y
        absy = $40

        place $3b,addr+absy
        ldy #absy
        lda #$44
        adc addr,Y
        bcs error
        sta $04                 ; Expect $80

;;; ADC zero page
        zp = $88

        place $b9,zp
        lda #$11
        adc zp
        bcs error
        sta $05                 ; Expect $ca

;;; ADC zero page,X
        zpx = $28

        ldx #zpx
        lda #$75
        adc zp-zpx,X
        bcc error
        sta $06                 ; Expect $2e

;;; ADC indexed indirect
        iiaddr = $4000
        iiloc = $98
        xii = $08

        staddr iiaddr, iiloc
        ldx #xii
        lda #$24
        adc (iiloc-xii,X)
        bcs error
        sta $07                 ; Expect $ad

;;; ADC indirect indexed
        yii = $96

        ldy #yii
        lda #$70
        adc (iiloc),Y
        bcs error
        sta $08                 ; Expect $e3

        brk
error:  lda #$99
        sta $00
