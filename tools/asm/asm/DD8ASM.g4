grammar DD8ASM;

options { caseInsensitive = true; }

prog
  : statement* EOF
  ;

statement
  : label 
  | (instruction | prep_instruction)? EOL
  ;

prep_instruction
  : P_ORG argument
  | P_INC str
  | P_DEF '(' EOL prep_def_arg_lines EOL ')'
  | P_DEF prep_def_args
  | P_TMPL name '(' namelist? ')'
  | P_DB arglist
  | P_DB '(' EOL arglist_lines EOL ')'
  | P_DW arglist
  | P_DW '(' EOL arglist_lines EOL ')'
  | P_BYTE num?
  | P_WORD num?
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

arglist_lines
  : arglist (EOL arglist_lines)?
  ;

argument
  : expr
  ;

expr
  : '~' expr 
  | expr ('.l' | '.h' ) 
  | expr ('*' | '/' | '%' ) expr
  | expr ('+' | '-') expr 
  | '-' expr 
  | expr ('<<' | '>>' ) expr 
  | expr '&' expr 
  | expr '^' expr 
  | expr '|' expr 
  | '(' expr ')'
  | (num | name | reg | str)
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

namelist
  : name (',' namelist)?
  ;

name
  : NAME
  ;

label
  : NAME ':'
  ;

P_DEF: '.def';
P_TMPL: '.tmpl';
P_ORG: '.org';
P_INC: '.inc';
P_DB: '.db';
P_DW: '.dw';
P_BYTE: '.byte';
P_WORD: '.word';

REG
  : [AXYZ]
  ;

STR 
  : '"' ( '\\"' | ~["\r\n] )* '"'
  ;

HEX_NUM
  : '0x'[0-9abcdef_]+
  ;

BIN_NUM
  : '0b'[01_]+
  ;

DEC_NUM
  : [0-9_]+
  ;

NAME
  : ([A-Z_][A-Z0-9_-]+|[B-W])
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

