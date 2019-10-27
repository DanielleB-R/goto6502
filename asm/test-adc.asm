;;; ADC immediate
        clc
        lda #$22
        ldx #$00
        ldy #$00
        adc #$8d
        bcs error
        sta $01                 ; Expect $af

;;; ADC absolute
        ldx #$88
        stx $4000
        lda #$cd
        adc $4000
        bcc error
        sta $02                 ; Expect $55

;;; ADC absolute,X
        ldy #$73
        sty $4096
        ldx #$96
        lda #$d7
        adc $4000,X
        bcc error
        sta $03                 ; Expect $4b

;;; ADC absolute,Y
        ldx #$3b
        stx $4040
        ldy #$40
        lda #$44
        adc $4000,Y
        bcs error
        sta $04                 ; Expect $80

;;; ADC zero page
        ldy #$b9
        sty $88
        lda #$11
        adc $88
        bcs error
        sta $05                 ; Expect $ca

;;; ADC zero page,X
        ldx #$18
        lda #$75
        adc $70,X
        bcc error
        sta $06                 ; Expect $2e

;;; ADC indexed indirect
        lda #$00
        sta $98
        lda #$40
        sta $99
        ldx #$08
        lda #$24
        adc ($90,X)
        bcs error
        sta $07                 ; Expect $ad

;;; ADC indirect indexed
        ldy #$96
        lda #$70
        adc ($98),Y
        bcs error
        sta $08                 ; Expect $e3

        brk
error:  lda #$99
        sta $00
