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
        addr = $2020
        index = $01
        lda #$ab
        sta addr
        inc addr
        inc addr                ; Expect $ad
        lda #$24
        sta addr + index
        ldx #index
        inc addr,X
        inc addr,X             ; Expect $26

;;; INC Zero Page
        zp = $05
        ind = $04
        lda #$1f
        sta zp
        inc zp
        inc zp                  ; Expect $21
        ldx #ind
        sta $06
        inc $06-ind,X               ; Expect $20
