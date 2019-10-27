;;; AND immediate
        lda #$f0
        ldx #$00
        ldy #$00
        and #$44
        sta $00                 ; Expect $40

;;; AND absolute
        lda #$f0
        ldx #$99
        stx $5000
        and $5000
        sta $01                 ; Expect $90

;;; AND absolute,X
        lda #$0f
        ldx #$02
        and $4ffe,X
        sta $02                 ; Expect $09

;;; AND absolute,Y
        lda #$85
        ldy #$03
        and $4ffd,Y
        sta $03                 ; Expect $81

;;; AND indexed indirect,X
        lda #$00
        sta $c1
        lda #$40
        sta $c2
        lda #$f1
        sta $4000
        lda #$c5
        ldx #$11
        and ($b0,X)
        sta $04                 ; Expect $c1

;;; AND indirect indexed,Y
        lda #$dd
        sta $4015
        ldy #$15
        lda #$bb
        and ($c1),Y
        sta $05                 ; Expect $99

;;; AND zero page
        ldx #$11
        stx $cc
        lda #$c9
        and $cc
        sta $06                 ; Expect $01

;;; AND zero page,X
        lda #$d6
        ldx #$0c
        and $c0,X
        sta $07                 ; Expect $10
