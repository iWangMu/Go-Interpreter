源代码 => 词法单元 => 抽象语法树

词法分析：用词法分析器将源代码转换为词法单元。
词法分析器/词法单元生成器（tokenizer）/扫描器（scanner）

语法分析器：将词法单元转换成抽象语法树。

// let x = 5 + 5; 
// [LET, IDENTIFIER("x"), EQUAL_SIGN, INTEGER(5), PLUS_SIGN, INTEGER(5), SEMICOLON]


词法单元（token）


语法分析器是一个软件组件，用于将输入的数据(通常是文本)构建成 一个数据结构，
通常是某种解析树、抽象语法树或其他层次结构。

词法分析器 => 语法分析器

AST 抽象语法树 Abstract Syntax Tree

语法分析器生成器（yacc、bison 或 ANTLR ），上下文无关文法（context-free grammar CFG）
Backus-Naur形式 BNF
Extended Backus-Naur形式 EBNF

语法分析策略：自上而下的分析 / 自下而上的分析。
自上而下，递归下降、Earley分析、预测分析等。

递归下降语法分析器，Vaughan Pratt 普拉特语法分析器，基于自上而下的运算符优先级分析法的语法分析器。

自上而下，从构造AST的根节点开始，然后下降。
自下而上，则以相反的方式进行构造。













