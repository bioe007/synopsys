let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/go/src/github.com/bioe007/synopsys
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +28 synopsys.go
badd +62 ~/go/src/github.com/bioe007/synopsys/disk/disk.go
badd +7 README.md
badd +0 disk_test.go
badd +24 .gitignore
badd +4 stat.go
badd +142 memory/memory.go
badd +45 ~/go/src/github.com/bioe007/synopsys/load/load.go
badd +0 cpu/cpu.go
argglobal
%argdel
$argadd synopsys.go
edit ~/go/src/github.com/bioe007/synopsys/load/load.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
wincmd _ | wincmd |
split
1wincmd k
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 82 + 83) / 166)
exe '2resize ' . ((&lines * 33 + 35) / 70)
exe 'vert 2resize ' . ((&columns * 83 + 83) / 166)
exe '3resize ' . ((&lines * 33 + 35) / 70)
exe 'vert 3resize ' . ((&columns * 83 + 83) / 166)
argglobal
balt cpu/cpu.go
setlocal fdm=expr
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=50
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 45 - ((44 * winheight(0) + 33) / 67)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 45
normal! 0
wincmd w
argglobal
if bufexists(fnamemodify("cpu/cpu.go", ":p")) | buffer cpu/cpu.go | else | edit cpu/cpu.go | endif
if &buftype ==# 'terminal'
  silent file cpu/cpu.go
endif
balt ~/go/src/github.com/bioe007/synopsys/load/load.go
setlocal fdm=expr
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=50
setlocal fml=1
setlocal fdn=20
setlocal fen
let s:l = 107 - ((20 * winheight(0) + 16) / 33)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 107
normal! 052|
wincmd w
argglobal
enew | setl bt=help
help 'wrap'@en
balt cpu/cpu.go
setlocal fdm=manual
setlocal fde=nvim_treesitter#foldexpr()
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=50
setlocal fml=1
setlocal fdn=20
setlocal nofen
silent! normal! zE
let &fdl = &fdl
let s:l = 7201 - ((18 * winheight(0) + 16) / 33)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 7201
normal! 049|
wincmd w
exe 'vert 1resize ' . ((&columns * 82 + 83) / 166)
exe '2resize ' . ((&lines * 33 + 35) / 70)
exe 'vert 2resize ' . ((&columns * 83 + 83) / 166)
exe '3resize ' . ((&lines * 33 + 35) / 70)
exe 'vert 3resize ' . ((&columns * 83 + 83) / 166)
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
nohlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
