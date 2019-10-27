;; Load zero page
        lda #$01
        ldx #$10
        ldy #$22
        sta $a0
        stx $a1
        sty $a2
        lda $a2
        ldx $a0
        ldy $a1
        sta $00                 ; Expect $22
        stx $01                 ; Expect $01
        sty $02                 ; Expect $10

;;; Load absolute
        lda $1000
        ldx $1002
        ldy $1004
        sta $03                ; Expect $a9
        stx $04                ; Expect $a2
        sty $05                ; Expect $a0

;;; Load absolute index X
        lda #$01
        ldx #$02
        ldy #$00
        lda $1002-$02,X
        ldy $1000-$02,X
        sta $06                 ; Expect $a2
        sty $07                 ; Expect $a9

;;; Load absolute index Y
        lda #$01
        ldx #$00
        ldy #$02
        lda $1004-$02,Y
        ldx $1002-$02,Y
        sta $08                 ; Expect $a0
        stx $09                 ; Expect $a2

;;; Load zero page index X
        ldx #$10
        lda $f2,X
        ldy $f1,X
        sta $0a                 ; Expect $10
        sty $0b                 ; Expect $01

;;; Load zero page index Y
        lda #$44
        ldy #$11
        ldx $f0,Y
        stx $0c                 ; Expect $01

;;; Load indexed indirect
        lda #$00
        sta $81
        lda #$10
        sta $82
        lda ($80,X)
        sta $0d                 ; Expect $a9

;;; Load indirect indexed
        ldy #$02
        lda ($81),Y
        sta $0e                 ; Expect $a2
