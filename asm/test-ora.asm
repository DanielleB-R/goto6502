;;; ORA immediate
        lda #$f0
        ldx #$00
        ldy #$00
        ora #$44
        sta $00                 ; Expect $f4

;;; ORA absolute
        lda #$f0
        ldx #$99
        stx $5000
        ora $5000
        sta $01                 ; Expect $f9

;;; ORA absolute,X
        lda #$0f
        ldx #$02
        ora $4ffe,X
        sta $02                 ; Expect $9f

;;; ORA absolute,Y
        lda #$85
        ldy #$03
        ora $4ffd,Y
        sta $03                 ; Expect $9d

;;; ORA zero page
        ldx #$11
        stx $cc
        lda #$c8
        ora $cc
        sta $04                 ; Expect $d9

;;; ORA zero page,X
        lda #$00
        ldx #$0c
        ora $c0,X
        sta $05                 ; Expect $11

;;; ORA indexed indirect,X
        lda #$00
        sta $c1
        lda #$40
        sta $c2
        lda #$f1
        sta $4000
        lda #$c5
        ldx #$11
        ora ($b0,X)
        sta $06                 ; Expect $f5

;;; ORA indirect indexed,Y
        lda #$88
        sta $4015
        ldy #$15
        lda #$55
        ora ($c1),Y
        sta $07                 ; Expect $dd
