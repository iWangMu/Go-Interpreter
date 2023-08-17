源代码 => 词法单元 => 抽象语法树

词法分析：用词法分析器将源代码转换为词法单元。
词法分析器/词法单元生成器（tokenizer）/扫描器（scanner）

语法分析器：将词法单元转换成抽象语法树。

// let x = 5 + 5; 
// [LET, IDENTIFIER("x"), EQUAL_SIGN, INTEGER(5), PLUS_SIGN, INTEGER(5), SEMICOLON]


词法单元（token）