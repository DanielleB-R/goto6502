;;;  Going to assume we start with C = 0
;;;  Not my favourite

;;;  BCC
        lda #$aa
        bcc cc
        jmp error
cc:     sta $01                 ; Expect $aa

;;; SEC
        lda #$ed
        sec
        bcc error
        sta $02                 ; Expect $ed

;;; BCS
        lda #$86
        bcs cs
        jmp error
cs:     sta $03                 ; Expect $86

;;; CLC
        lda #$27
        clc
        bcs error
        sta $04                 ; Expect $27

        brk

error:  lda #$99
        sta $00
