        .include "macros.asm"
;;; ASL A
        lda #$8f
        ldx #$ff
        asl a
        bcc error
        sta $01                 ; Expect $1e

;;; LSR A
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

;;; ASL absolute
        absolute = $5500

        clc
        xplace $9b, absolute
        asl absolute
        bcc error
        lda absolute
        sta $07                 ; Expect $3c

;;; ASL absolute, X
        offset = 5

        ldx #offset
        yplace $78, absolute+offset
        asl absolute,X
        bcs error
        lda absolute,X
        sta $08                 ; Expect $f0

;;; ASL zero page
        zp = $cc

        yplace $f9, zp
        asl zp
        bcc error
        lda zp
        sta $09                 ; Expect $f2

;;; ASL zero page, X
        yplace $19, zp+offset
        asl zp,X
        bcs error
        lda zp,X
        sta $0a                 ; Expect $32

        brk
error:  place $99, $00
