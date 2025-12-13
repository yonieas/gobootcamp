.intel_syntax noprefix
.global printNama
.text
printNama:
  mov rax, 1
  mov rdi, 1
  lea rsi, nama
  mov rdx, 8
  syscall
  
  mov rax, 60
  xor rdi, rdi
  syscall

.data
nama:
  .ascii "Herbayu\n"