;;; STRCPY implementation, with hard-coded src and dest
        dest=$3000
        ldx #$00
loop:   lda src,X
        sta dest,X
        beq done
        inx
        jmp loop
done:   brk

src:    .asciiz "Running a program!"
