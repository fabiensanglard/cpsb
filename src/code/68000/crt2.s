
.align 4
_boot:
	* Enable auto-interrupts
	move.w	#0x2000, sr

	* Jump to C main()
	jbsr	main