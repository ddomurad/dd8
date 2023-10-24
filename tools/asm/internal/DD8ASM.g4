grammar DD8ASM;

options { caseInsensitive = true; }

prog
  : statement* EOF
  ;

statement
  : (instruction | prep_instruction | label)? EOL
  ;

prep_instruction
  : P_ORG argument
  | P_INC str
  | P_DEF '(' EOL prep_def_arg_lines EOL ')'
  | P_DEF prep_def_args
  | P_DB arglist
  | P_DW arglist
  ;

prep_def_args
  : name ':=' argument
  ;

prep_def_arg_lines
  : prep_def_args (EOL prep_def_arg_lines)?
  ;

instruction
  : name 
  | name arglist_p 
  | name arglist 
  | name arglist_p ',' arglist
  | name arglist ',' arglist_p
  ;

arglist_p
  : '[' arglist ']'
  ;

arglist
  : argument (',' arglist)?
  ;

argument
  : (num | reg | name | str)
  ;

str 
  : STR
  ;

num
  : (DEC_NUM | HEX_NUM | BIN_NUM)
  ;

reg
  : REG
  ;

name
  : NAME
  ;

label
  : NAME ':'
  ;

P_DEF: '.def';
P_ORG: '.org';
P_INC: '.inc';
P_DB: '.db';
P_DW: '.dw';

REG
  : [A-Z]
  ;

STR 
  : '"' ( '""' | ~["\r\n] )* '"'
  ;

HEX_NUM
  : '0x'[0-9abcdef_]+
  ;

BIN_NUM
  : '0b'[01_]+
  ;

DEC_NUM
  : '-'?[0-9_]+
  ;

NAME
  : [A-Z_-]+[A-Z_0-9-]+
  ;

COMMENT
  : ';' ~ [\r\n]* -> skip
  ;

EOL
  : [\r\n] +
  ;

WS
  : [ \t] -> skip
  ;

