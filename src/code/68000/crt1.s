.align 4
VSync:
    jsr onVSync
	rte

.align 4
_boot:
	* Enable interrupts
	move.w	#0x2000, sr

	* Jump to C main()
	jbsr	main
