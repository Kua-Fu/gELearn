%{
package yacc1

%}

%union {
    sql *Sql
    num int
    str string
    strs []string
}

%type <sql> expr 
%type <strs> columns
%token LeftBracket RightBracket

%token	<str> TokColumn TokTable
%token	<num> TokNum

%%

top:
expr
{
    yylex.(*lex).sql = $1
}

expr:
{
	$$ = &Sql{}
}
| LeftBracket columns RightBracket
{
	$$ = &Sql{columns: $2}
}

columns:
TokColumn
{
	$$ = []string{$1}
}
| columns TokColumn
{
	$$ = append($$, $2)
}

%%


	  



