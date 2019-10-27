;;; ASL
        lda #$8f
        ldx #$ff
        asl a
        bcc error
        sta $01                 ; Expect $1e

;;; LSR
        lda #$f2
        lsr a
        bcs error
        sta $02                 ; Expect $79

;;; ROL A
        clc
        lda #$81
        rol A
        bcc error
        sta $03                 ; Expect $02
        rol A
        bcs error
        sta $04                 ; Expect $05

;;; ROR A
        clc
        lda #$81
        ror A
        bcc error
        sta $05                 ; Expect $40
        ror A
        bcs error
        sta $06                 ; Expect $a0
        brk

error:  lda #$99
        sta $00
