.module crt0
.globl  _main
.area  _HEADER (ABS)

;------------------
; Z-80 starts here!
;------------------
.org 0
   jp init  

;--------------
; INTERRUPT
;--------------
.org 0x38
   DI                     ; Disable Interrupt
   call _interrupt        ; Process Interrupt
   call _schedInterrupt   ; Reschedule interrupt
   EI                     ; Enable  Interrupt
   RET 

;--------------
; INIT and MAIN
;--------------
.org 0x100
init:
   
   ld  sp,#0xd7ff         ; Setup stack
   IM 1                   ; Set Interupt mode 1
   call _schedInterrupt   ; Request first int
   call  gsinit           ; Init global variables
main:  
   call  _main            ; Call C main()
   jp    main             ; Never happens