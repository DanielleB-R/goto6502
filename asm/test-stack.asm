        .include "macros.asm"
;;; Setting the stack pointer
        ldx #$20
        txs
        ldx #$ff
        tsx
        stx $00                 ; Expect $20
;;; Pushing
        lda #$ec
        pha                     ; Expect this stored at $0120
        lda #$8f
        pha                     ; Expect this stored at $011f
;;; Pulling
        ldx #$ff
        txs
        lda #$0e
        pha
        lda #$a6
        pha
        lda #$00
        pla
        sta $01                 ; expect $a6
        pla
        sta $02                 ; expect $0e
;;; transfer to accumulator
        ldx #$ce
        txa
        sta $03                 ; expect $ce
        ldy #$93
        tya
        sta $04                 ; expect $93

;;; pushing flags
        sec
        lda #$88
        php
        pla
        sta $05                 ; expect $81
