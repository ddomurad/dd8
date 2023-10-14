grammar DD8ASM;

options { caseInsensitive = true; }

prog
  : statement* EOF
  ;

statement
  : (instruction | prep_instruction)? EOL
  ;

prep_instruction
  : prep_name prep_arglist?
  | prep_name '(' prep_arglist? ')'
  ;

instruction
  : name arglist?
  ;

prep_arglist
  : argument (',' arglist)?
  | argument ('=' arglist)?
  ;

arglist
  : argument (',' arglist)?
  ;

argument
  : (num | name)
  ;

num
  : (DEC_NUM | HEX_NUM | BIN_NUM)
  ;

name
  : NAME
  ;

prep_name
  : PREP_NAME
  ;

NAME
  : [A-Z_-]+
  ;

PREP_NAME
  : '.'[A-Z_-]+
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

COMMENT
   : ';' ~ [\r\n]* -> skip
   ;

EOL
   : [\r\n] +
   ;

WS
   : [ \t] -> skip
   ;
