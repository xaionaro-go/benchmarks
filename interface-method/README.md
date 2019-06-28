```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/interface-method
BenchmarkPtrMethodDirect-8         	2000000000	         0.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodDirect-8            	2000000000	         0.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkPtrMethodViaInterface-8   	1000000000	         1.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodViaInterface-8      	1000000000	         3.16 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/interface-method	7.052s
```

```
00000000004ed7f0 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkPtrMethodDirect>:
  4ed7f0:       64 48 8b 0c 25 f8 ff    mov    %fs:0xfffffffffffffff8,%rcx
  4ed7f7:       ff ff 
  4ed7f9:       48 3b 61 10             cmp    0x10(%rcx),%rsp
  4ed7fd:       76 3c                   jbe    4ed83b <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkPtrMethodDirect+0x4b>
  4ed7ff:       48 83 ec 10             sub    $0x10,%rsp
  4ed803:       48 89 6c 24 08          mov    %rbp,0x8(%rsp)
  4ed808:       48 8d 6c 24 08          lea    0x8(%rsp),%rbp
  4ed80d:       48 8b 44 24 18          mov    0x18(%rsp),%rax
  4ed812:       48 89 04 24             mov    %rax,(%rsp)
  4ed816:       e8 65 eb fb ff          callq  4ac380 <testing.(*B).ResetTimer>
  4ed81b:       48 8b 44 24 18          mov    0x18(%rsp),%rax
  4ed820:       31 c9                   xor    %ecx,%ecx
  4ed822:       eb 04                   jmp    4ed828 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkPtrMethodDirect+0x38>
  4ed824:       48 ff c1                inc    %rcx
  4ed827:       90                      nop
  4ed828:       48 39 88 08 01 00 00    cmp    %rcx,0x108(%rax)
  4ed82f:       7f f3                   jg     4ed824 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkPtrMethodDirect+0x34>
  4ed831:       48 8b 6c 24 08          mov    0x8(%rsp),%rbp
  4ed836:       48 83 c4 10             add    $0x10,%rsp
  4ed83a:       c3                      retq   
  4ed83b:       e8 20 7a f6 ff          callq  455260 <runtime.morestack_noctxt>
  4ed840:       eb ae                   jmp    4ed7f0 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkPtrMethodDirect>
  4ed842:       cc                      int3   
  4ed843:       cc                      int3   
  4ed844:       cc                      int3   
  4ed845:       cc                      int3   
  4ed846:       cc                      int3   
  4ed847:       cc                      int3   
  4ed848:       cc                      int3   
  4ed849:       cc                      int3   
  4ed84a:       cc                      int3   
  4ed84b:       cc                      int3   
  4ed84c:       cc                      int3   
  4ed84d:       cc                      int3   
  4ed84e:       cc                      int3   
  4ed84f:       cc                      int3   
```

```
00000000004ed850 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkMethodDirect>:
  4ed850:       64 48 8b 0c 25 f8 ff    mov    %fs:0xfffffffffffffff8,%rcx
  4ed857:       ff ff 
  4ed859:       48 3b 61 10             cmp    0x10(%rcx),%rsp
  4ed85d:       76 3c                   jbe    4ed89b <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkMethodDirect+0x4b>
  4ed85f:       48 83 ec 10             sub    $0x10,%rsp
  4ed863:       48 89 6c 24 08          mov    %rbp,0x8(%rsp)
  4ed868:       48 8d 6c 24 08          lea    0x8(%rsp),%rbp
  4ed86d:       48 8b 44 24 18          mov    0x18(%rsp),%rax
  4ed872:       48 89 04 24             mov    %rax,(%rsp)
  4ed876:       e8 05 eb fb ff          callq  4ac380 <testing.(*B).ResetTimer>
  4ed87b:       48 8b 44 24 18          mov    0x18(%rsp),%rax
  4ed880:       31 c9                   xor    %ecx,%ecx
  4ed882:       eb 04                   jmp    4ed888 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkMethodDirect+0x38>
  4ed884:       48 ff c1                inc    %rcx
  4ed887:       90                      nop
  4ed888:       48 39 88 08 01 00 00    cmp    %rcx,0x108(%rax)
  4ed88f:       7f f3                   jg     4ed884 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkMethodDirect+0x34>
  4ed891:       48 8b 6c 24 08          mov    0x8(%rsp),%rbp
  4ed896:       48 83 c4 10             add    $0x10,%rsp
  4ed89a:       c3                      retq   
  4ed89b:       e8 c0 79 f6 ff          callq  455260 <runtime.morestack_noctxt>
  4ed8a0:       eb ae                   jmp    4ed850 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkMethodDirect>
  4ed8a2:       cc                      int3   
  4ed8a3:       cc                      int3   
  4ed8a4:       cc                      int3   
  4ed8a5:       cc                      int3   
  4ed8a6:       cc                      int3   
  4ed8a7:       cc                      int3   
  4ed8a8:       cc                      int3   
  4ed8a9:       cc                      int3   
  4ed8aa:       cc                      int3   
  4ed8ab:       cc                      int3   
  4ed8ac:       cc                      int3   
  4ed8ad:       cc                      int3   
  4ed8ae:       cc                      int3   
  4ed8af:       cc                      int3   
```

```
00000000004ed8b0 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkPtrMethodViaInterface>:
  4ed8b0:       64 48 8b 0c 25 f8 ff    mov    %fs:0xfffffffffffffff8,%rcx
  4ed8b7:       ff ff 
  4ed8b9:       48 3b 61 10             cmp    0x10(%rcx),%rsp
  4ed8bd:       76 76                   jbe    4ed935 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkPtrMethodViaInterface+0x85>
  4ed8bf:       48 83 ec 28             sub    $0x28,%rsp
  4ed8c3:       48 89 6c 24 20          mov    %rbp,0x20(%rsp)
  4ed8c8:       48 8d 6c 24 20          lea    0x20(%rsp),%rbp
  4ed8cd:       48 8d 05 ac 53 02 00    lea    0x253ac(%rip),%rax        # 512c80 <type.*+0x24c80>
  4ed8d4:       48 89 04 24             mov    %rax,(%rsp)
  4ed8d8:       e8 f3 e3 f1 ff          callq  40bcd0 <runtime.newobject>
  4ed8dd:       48 8b 44 24 08          mov    0x8(%rsp),%rax
  4ed8e2:       48 89 44 24 18          mov    %rax,0x18(%rsp)
  4ed8e7:       48 8b 4c 24 30          mov    0x30(%rsp),%rcx
  4ed8ec:       48 89 0c 24             mov    %rcx,(%rsp)
  4ed8f0:       e8 8b ea fb ff          callq  4ac380 <testing.(*B).ResetTimer>
  4ed8f5:       31 c0                   xor    %eax,%eax
  4ed8f7:       eb 24                   jmp    4ed91d <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkPtrMethodViaInterface+0x6d>
  4ed8f9:       48 89 44 24 10          mov    %rax,0x10(%rsp)
  4ed8fe:       48 8d 05 bb dd 06 00    lea    0x6ddbb(%rip),%rax        # 55b6c0 <go.itab.*github.com/xaionaro-go/benchmarks/interface-method.Struct,github.com/xaionaro-go/benchmarks/interface-method.Interface>
  4ed905:       84 00                   test   %al,(%rax)
  4ed907:       48 8b 4c 24 18          mov    0x18(%rsp),%rcx
  4ed90c:       48 89 0c 24             mov    %rcx,(%rsp)
  4ed910:       e8 bb fe ff ff          callq  4ed7d0 <github.com/xaionaro-go/benchmarks/interface-method.(*Struct).PtrMethod>
  4ed915:       48 8b 44 24 10          mov    0x10(%rsp),%rax
  4ed91a:       48 ff c0                inc    %rax
  4ed91d:       48 8b 4c 24 30          mov    0x30(%rsp),%rcx
  4ed922:       48 39 81 08 01 00 00    cmp    %rax,0x108(%rcx)
  4ed929:       7f ce                   jg     4ed8f9 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkPtrMethodViaInterface+0x49>
  4ed92b:       48 8b 6c 24 20          mov    0x20(%rsp),%rbp
  4ed930:       48 83 c4 28             add    $0x28,%rsp
  4ed934:       c3                      retq   
  4ed935:       e8 26 79 f6 ff          callq  455260 <runtime.morestack_noctxt>
  4ed93a:       e9 71 ff ff ff          jmpq   4ed8b0 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkPtrMethodViaInterface>
  4ed93f:       cc                      int3   
```

```
00000000004ed940 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkMethodViaInterface>:
  4ed940:       64 48 8b 0c 25 f8 ff    mov    %fs:0xfffffffffffffff8,%rcx
  4ed947:       ff ff 
  4ed949:       48 3b 61 10             cmp    0x10(%rcx),%rsp
  4ed94d:       76 76                   jbe    4ed9c5 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkMethodViaInterface+0x85>
  4ed94f:       48 83 ec 28             sub    $0x28,%rsp
  4ed953:       48 89 6c 24 20          mov    %rbp,0x20(%rsp)
  4ed958:       48 8d 6c 24 20          lea    0x20(%rsp),%rbp
  4ed95d:       48 8d 05 1c 53 02 00    lea    0x2531c(%rip),%rax        # 512c80 <type.*+0x24c80>
  4ed964:       48 89 04 24             mov    %rax,(%rsp)
  4ed968:       e8 63 e3 f1 ff          callq  40bcd0 <runtime.newobject>
  4ed96d:       48 8b 44 24 08          mov    0x8(%rsp),%rax
  4ed972:       48 89 44 24 18          mov    %rax,0x18(%rsp)
  4ed977:       48 8b 4c 24 30          mov    0x30(%rsp),%rcx
  4ed97c:       48 89 0c 24             mov    %rcx,(%rsp)
  4ed980:       e8 fb e9 fb ff          callq  4ac380 <testing.(*B).ResetTimer>
  4ed985:       31 c0                   xor    %eax,%eax
  4ed987:       eb 24                   jmp    4ed9ad <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkMethodViaInterface+0x6d>
  4ed989:       48 89 44 24 10          mov    %rax,0x10(%rsp)
  4ed98e:       48 8d 05 2b dd 06 00    lea    0x6dd2b(%rip),%rax        # 55b6c0 <go.itab.*github.com/xaionaro-go/benchmarks/interface-method.Struct,github.com/xaionaro-go/benchmarks/interface-method.Interface>
  4ed995:       84 00                   test   %al,(%rax)
  4ed997:       48 8b 4c 24 18          mov    0x18(%rsp),%rcx
  4ed99c:       48 89 0c 24             mov    %rcx,(%rsp)
  4ed9a0:       e8 8b 00 00 00          callq  4eda30 <github.com/xaionaro-go/benchmarks/interface-method.(*Struct).Method>
  4ed9a5:       48 8b 44 24 10          mov    0x10(%rsp),%rax
  4ed9aa:       48 ff c0                inc    %rax
  4ed9ad:       48 8b 4c 24 30          mov    0x30(%rsp),%rcx
  4ed9b2:       48 39 81 08 01 00 00    cmp    %rax,0x108(%rcx)
  4ed9b9:       7f ce                   jg     4ed989 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkMethodViaInterface+0x49>
  4ed9bb:       48 8b 6c 24 20          mov    0x20(%rsp),%rbp
  4ed9c0:       48 83 c4 28             add    $0x28,%rsp
  4ed9c4:       c3                      retq   
  4ed9c5:       e8 96 78 f6 ff          callq  455260 <runtime.morestack_noctxt>
  4ed9ca:       e9 71 ff ff ff          jmpq   4ed940 <github.com/xaionaro-go/benchmarks/interface-method.BenchmarkMethodViaInterface>
  4ed9cf:       cc                      int3   
```
