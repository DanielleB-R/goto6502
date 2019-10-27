        .include "macros.asm"
;;; BVS
        lda #$7f
        ldx #$00
        ldy #$00
        adc #$1f
        bvs overflow
        jmp error
overflow: place $24, $01          ; Expect $24 at $01

;;; BVC
        lda #$01
        adc #$10
        bvc no
        jmp error
no:     place $8e, $02          ; Expect $8e at $02

;;; CLV
        lda #$6a
        adc #$6b
        clv
        bvc cleared
        jmp error
cleared:        place $d4, $03  ; Expect $d4 at $03

        brk
error:  place $99, $00
