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
        lda $1000,X
        ldy $0ffe,X
        sta $06                 ; Expect $a2
        sty $07                 ; Expect $a9

;;; Load absolute index Y
        lda #$01
        ldx #$00
        ldy #$02
        lda $1002,Y
        ldx $1000,Y
        sta $08                 ; Expect $a0
        stx $09                 ; Expect $a2
