LDFLAGS=-t none

.PHONY: clean all

%.bin: %.o
	ld65 $(LDFLAGS) $< -o $@

%.o: %.asm
	ca65 $<

all: test1.bin test2.bin test4.bin test-lda.bin test-sta.bin test-inc.bin test-dec.bin test-jmp.bin test-branch.bin strcpy.bin test-and.bin test-ora.bin test-carry.bin test-shift.bin test-stack.bin

clean:
	rm *.o *.bin
