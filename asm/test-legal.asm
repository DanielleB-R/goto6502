        SEI
        NOP
        CLI
        NOP
        SED
        NOP
        CLD
        NOP
        LSR $2000
        NOP
        LSR $4444,X
        NOP
        LSR $52
        NOP
        LSR $52,X
        NOP

        ROL $2000
        NOP
        ROL $4444,X
        NOP
        ROL $52
        NOP
        ROL $52,X
        NOP
        ROR $2000
        NOP
        ROR $4444,X
        NOP
        ROR $52
        NOP
        ROR $52,X
        NOP

        BIT $2000
        NOP
        BIT $84
        NOP

        SBC #$02
        NOP
        SBC $2000
        NOP
        SBC $4444,X
        NOP
        SBC $8888,Y
        NOP
        SBC $52
        NOP
        SBC $52,X
        NOP
        SBC ($38,X)
        NOP
        SBC ($89),Y
        NOP
