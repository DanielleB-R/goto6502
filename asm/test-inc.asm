        lda #$00
;;; Increment X index
        ldx #$a5
        inx
        inx
        stx $01                 ; Expect $a7
        inx
        stx $02                 ; Expect $a8

;;; Increment Y index
        ldy #$88
        iny
        sty $03                 ; Expect $89
        iny
        iny
        iny
        sty $04                 ; Expect $8c

;;; INC Absolute
        lda #$ab
        sta $2020
        inc $2020
        inc $2020               ; Expect $ad
        lda #$24
        sta $2021
        ldx #$01
        inc $2020,X
        inc $2020,X             ; Expect $26

;;; INC Zero Page
        lda #$1f
        sta $05
        inc $05
        inc $05                 ; Expect $21
        ldx #$04
        sta $06
        inc $02,X               ; Expect $20
