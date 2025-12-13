.intel_syntax noprefix

.section .text
.global _start

_start:
    mov rax, 1
    mov rdi, 1
    lea rsi, [name]
    mov rdx, 19
    syscall
    
    mov rax, 60
    xor rdi, rdi
    syscall

.section .data
name:
    .ascii "M Faisal Abdillah\n"