LDFLAGS=-t none

.PHONY: clean all

%.bin: %.o
	ld65 $(LDFLAGS) $< -o $@

%.o: %.asm
	ca65 $<

all: load-immediate.bin test2.bin test4.bin test-lda.bin test-sta.bin test-inc.bin test-dec.bin test-jmp.bin test-branch.bin strcpy.bin test-and.bin test-ora.bin test-carry.bin test-shift.bin test-stack.bin test-eor.bin test-adc.bin test-overflow.bin test-cmp.bin test-cpxy.bin test-subroutine.bin test-legal.bin

clean:
	rm *.o *.bin
