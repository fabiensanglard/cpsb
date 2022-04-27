.align 4
_boot:
    * Enable auto-interrupts
    move.w  #0x2000, sr

    * Init .BSS
    jbsr    clearBSS

    * Init .DATA
    jbsr    copyDATA

    * Jump to C main()
    jbsr    main