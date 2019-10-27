;;; EOR immediate
        lda #$f0
        ldx #$00
        ldy #$00
        eor #$44
        sta $00                 ; Expect $b4

;;; EOR absolute
        lda #$f0
        ldx #$99
        stx $5000
        eor $5000
        sta $01                 ; Expect $69

;;; EOR absolute,X
        lda #$0f
        ldx #$02
        eor $4ffe,X
        sta $02                 ; Expect $96

;;; EOR absolute,Y
        lda #$85
        ldy #$03
        eor $4ffd,Y
        sta $03                 ; Expect $1c

;;; EOR zero page
        ldx #$11
        stx $cc
        lda #$db
        eor $cc
        sta $04                 ; Expect $ca

;;; EOR zero page,X
        lda #$79
        ldx #$0c
        eor $c0,X
        sta $05                 ; Expect $68

;;; EOR indexed indirect,X
        lda #$00
        sta $c1
        lda #$40
        sta $c2
        lda #$f1
        sta $4000
        lda #$c5
        ldx #$11
        eor ($b0,X)
        sta $06                 ; Expect $f5

;;; EOR indirect indexed,Y
        lda #$77
        sta $4015
        ldy #$15
        lda #$55
        eor ($c1),Y
        sta $07                 ; Expect $22
