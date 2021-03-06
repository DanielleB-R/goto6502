        .include "macros.asm"
;;; Store absolute
        lda #$22
        ldx #$55
        ldy #$bb
        sta $2000               ; Expect $22
        stx $2001               ; Expect $55
        sty $2002               ; Expect $bb

;;; Store absolute index X
        ldx #$10
        sta $2000, X            ; Expect $22 at $2010

;;; Store absolute index Y
        ldy #$20
        sta $2000, Y            ; Expect $22 at $2020

;;; Store zero page X
        lda #$80
        ldy #$75
        ldx #$10
        sta $20,X               ; Expect $80 at $0030
        sty $21,X               ; Expect $75 at $0031

;;; Store zero page Y
        ldy #$22
        ldx #$ef
        stx $11,Y               ; Expect $ef at $0033

;;; Store indexed indirect
        iiaddr = $0035
        iiloc = $91
        xii = $01

        staddr iiaddr,iiloc
        lda #$ce
        ldx #xii
        sta (iiloc-xii,X)             ; Expect $ce at $0035

;;; Store indirect indexed
        lda #$dd
        ldy #$03
        sta (iiloc),Y             ; Expect $dd at $0038
