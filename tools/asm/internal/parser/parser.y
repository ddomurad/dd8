%{
    package parser
%}

%union{
    Source  Source
    Entries []SourceEntry
    Entry SourceEntry

    String   string
    Number   int64
}

%token <String> OP_CODE WORD IMPORT
%token <Number> DEC_NUM HEX_NUM

%type <Source> src
%type <Entries> src_entries
%type <Entry> src_entry

%%

src             :   src_entries                 {$$.Entries = $1; yylex.(*lexer).source = $$}
                ;
src_entries     :   src_entry                   {$$ = []SourceEntry{$1}}
                |   src_entries  src_entry      {$$ = append($1, $2)}
                ;
src_entry       :   OP_CODE                     {$$.Type = SourceEntryTypeOpCode; $$.OpCode = $1}
                |   OP_CODE  DEC_NUM            {$$.Type = SourceEntryTypeOpCode; $$.OpCode = $1; $$.NumParam = $2}
                |   OP_CODE  HEX_NUM            {$$.Type = SourceEntryTypeOpCode; $$.OpCode = $1; $$.NumParam = $2}
                ;
%%

