// Derived from https://bitbucket.org/qps/tsql-grammar/
//
// -----------------------------------------------------------------------------------
// T-SQL Grammar
// -----------------------------------------------------------------------------------
// In exception to the overall bsn ModuleStore project, this grammar file is licensed
// under LGPL and MIT license - http://www.opensource.org/licenses/mit-license.php
//
/*
"Name" = 'T-SQL 2008 (strict subset)'
       
"Author" = 'Arsene von Wyss'
       
"About" = 'Parser for T-SQL.'
    | 'This subset excludes ambiguous and deprecated SQL as well as some features not used by ModuleStore.'

"Character Mapping" = 'Unicode'

"Start Symbol" = <Script>

! CLR UDTs are missing, but may be implemented later

"Version" = '2013-01-31 - modified by Karl Spaelti'
*/

grammar tsql;

COMMENT: '/*' .* '*/' -> skip; 
LINE_COMMENT: '--' ~'\n'* -> skip;
WHITESPACE: [ \t\r\n]+ -> skip;

ADD: 'ADD';
ALL: 'ALL';
ALTER: 'ALTER';
AND: 'AND';
ANY: 'ANY';
AS: 'AS';
ASC: 'ASC';
AUTHORIZATION: 'AUTHORIZATION';
BACKUP: 'BACKUP';
BEGIN: 'BEGIN';
BETWEEN: 'BETWEEN';
BREAK: 'BREAK';
BROWSE: 'BROWSE';
BULK: 'BULK';
BY: 'BY';
CASCADE: 'CASCADE';
CASE: 'CASE';
CHECK: 'CHECK';
CHECKPOINT: 'CHECKPOINT';
CLOSE: 'CLOSE';
CLUSTERED: 'CLUSTERED';
COALESCE: 'COALESCE';
COLLATE: 'COLLATE';
COLUMN: 'COLUMN';
COMMIT: 'COMMIT';
COMPUTE: 'COMPUTE';
CONSTRAINT: 'CONSTRAINT';
CONTAINS: 'CONTAINS';
CONTAINSTABLE: 'CONTAINSTABLE';
CONTINUE: 'CONTINUE';
CONVERT: 'CONVERT';
CREATE: 'CREATE';
CROSS: 'CROSS';
CURRENT: 'CURRENT';
CURRENT_DATE: 'CURRENT_DATE';
CURRENT_TIME: 'CURRENT_TIME';
CURRENT_TIMESTAMP: 'CURRENT_TIMESTAMP';
CURRENT_USER: 'CURRENT_USER';
CURSOR: 'CURSOR';
DATABASE: 'DATABASE';
DBCC: 'DBCC';
DEALLOCATE: 'DEALLOCATE';
DECLARE: 'DECLARE';
DEFAULT: 'DEFAULT';
DELETE: 'DELETE';
DENY: 'DENY';
DESC: 'DESC';
DISK: 'DISK';
DISTINCT: 'DISTINCT';
DISTRIBUTED: 'DISTRIBUTED';
DOUBLE: 'DOUBLE';
DROP: 'DROP';
DUMMY: 'DUMMY';
DUMP: 'DUMP';
ELSE: 'ELSE';
END: 'END';
ERRLVL: 'ERRLVL';
ESCAPE: 'ESCAPE';
EXCEPT: 'EXCEPT';
EXECUTE: 'EXECUTE' | 'EXEC';
EXISTS: 'EXISTS';
EXIT: 'EXIT';
FETCH: 'FETCH';
FILE: 'FILE';
FILLFACTOR: 'FILLFACTOR';
FOR: 'FOR';
FOREIGN: 'FOREIGN';
FREETEXT: 'FREETEXT';
FREETEXTTABLE: 'FREETEXTTABLE';
FROM: 'FROM';
FULL: 'FULL';
FUNCTION: 'FUNCTION';
GOTO: 'GOTO';
GRANT: 'GRANT';
GROUP: 'GROUP';
HAVING: 'HAVING';
HOLDLOCK: 'HOLDLOCK';
IDENTITY: 'IDENTITY';
IDENTITY_INSERT: 'IDENTITY_INSERT';
IDENTITYCOL: 'IDENTITYCOL';
IF: 'IF';
IN: 'IN';
INDEX: 'INDEX';
INNER: 'INNER';
INSERT: 'INSERT';
INTERSECT: 'INTERSECT';
INTO: 'INTO';
IS: 'IS';
JOIN: 'JOIN';
KEY: 'KEY';
KILL: 'KILL';
LEFT: 'LEFT';
LIKE: 'LIKE';
LINENO: 'LINENO';
LOAD: 'LOAD';
NATIONAL: 'NATIONAL';
NOCHECK: 'NOCHECK';
NONCLUSTERED: 'NONCLUSTERED';
NOT: 'NOT';
NULL: 'NULL';
NULLIF: 'NULLIF';
OF: 'OF';
OFF: 'OFF';
OFFSETS: 'OFFSETS';
ON: 'ON';
OPEN: 'OPEN';
OPENDATASOURCE: 'OPENDATASOURCE';
OPENQUERY: 'OPENQUERY';
OPENROWSET: 'OPENROWSET';
OPENXML: 'OPENXML';
OPTION: 'OPTION';
OR: 'OR';
ORDER: 'ORDER';
OUTER: 'OUTER';
OVER: 'OVER';
PERCENT: 'PERCENT';
PLAN: 'PLAN';
PRECISION: 'PRECISION';
PRIMARY: 'PRIMARY';
PRINT: 'PRINT';
PROCEDURE: 'PROC' 'EDURE'?;
PUBLIC: 'PUBLIC';
RAISERROR: 'RAISERROR';
READ: 'READ';
READTEXT: 'READTEXT';
RECONFIGURE: 'RECONFIGURE';
REFERENCES: 'REFERENCES';
REPLICATION: 'REPLICATION';
RESTORE: 'RESTORE';
RESTRICT: 'RESTRICT';
RETURN: 'RETURN';
REVOKE: 'REVOKE';
RIGHT: 'RIGHT';
ROLLBACK: 'ROLLBACK';
ROWCOUNT: 'ROWCOUNT';
ROWGUIDCOL: 'ROWGUIDCOL';
RULE: 'RULE';
SAVE: 'SAVE';
SCHEMA: 'SCHEMA';
SELECT: 'SELECT';
SESSION_USER: 'SESSION_USER';
SET: 'SET';
SETUSER: 'SETUSER';
SHUTDOWN: 'SHUTDOWN';
SOME: 'SOME';
STATISTICS: 'STATISTICS';
SYSTEM_USER: 'SYSTEM_USER';
TABLE: 'TABLE';
TEXTSIZE: 'TEXTSIZE';
THEN: 'THEN';
TO: 'TO';
TOP: 'TOP';
TRANSACTION: 'TRAN' 'SACTION'?;
TRIGGER: 'TRIGGER';
TRUNCATE: 'TRUNCATE';
TSEQUAL: 'TSEQUAL';
UNION: 'UNION';
UNIQUE: 'UNIQUE';
UPDATE: 'UPDATE';
UPDATETEXT: 'UPDATETEXT';
USE: 'USE';
USER: 'USER';
VALUES: 'VALUES';
VARYING: 'VARYING';
VIEW: 'VIEW';
WAITFOR: 'WAITFOR';
WHEN: 'WHEN';
WHERE: 'WHERE';
WHILE: 'WHILE';
WITH: 'WITH';
WRITETEXT: 'WRITETEXT';

fragment ID_CH_STANDARD: [a-zA-Z0-9_];
fragment ID_CH_DELIMITED_BRACKET: [^\[\]];
fragment ID_CH_DELIMITED_BRACKET_START: [^\[\]#];
fragment ID_CH_DELIMITED_QUOTE: [^"#];
fragment ID_CH_SINGLE_DELIMITED_QUOTE: [^'#];
fragment ID_CH_DELIMITED_QUOTE_START: ID_CH_DELIMITED_QUOTE;

STRING_LITERAL: 'N'? '\'' ( '\\\'' | ~'\'' )* '\'';
INTEGER_LITERAL: [0-9]+;
REAL_LITERAL: [0-9]*'.'[0-9]+;
HEX_LITERAL: '0x'[0-9a-zA-Z]+;
fragment LETTER: [a-zA-Z];

ID: LETTER ID_CH_STANDARD* ;

QUOTED_ID: '[' ID_CH_DELIMITED_BRACKET_START ID_CH_DELIMITED_BRACKET* ']'
| '"' ID_CH_DELIMITED_QUOTE_START ID_CH_DELIMITED_QUOTE* '"' ;

// SINGLE_QUOTED_ID: '\'' ID_CH_SINGLE_DELIMITED_QUOTE* '\'';

LOCAL_ID: '@' LETTER ID_CH_STANDARD* ;

SYSTEM_VAR_ID: '@@' LETTER ID_CH_STANDARD* ;

SYSTEM_FUNC_ID: '::' LETTER ID_CH_STANDARD* ;

TEMP_TABLE_ID: '#' LETTER ID_CH_STANDARD*
| '[#' ID_CH_DELIMITED_BRACKET+ ']'
| '"#' ID_CH_DELIMITED_QUOTE+ '"' ;

objectName: QUOTED_ID
| ID ;

objectNameQualified: (schemaName '.')? objectName ;

aliasName: objectName
| STRING_LITERAL ;

assemblyName: objectName ;

className: objectName ;

collationName: ID ;

columnName: objectName
| '$ACTION' ;

columnNameFullyQualified
: schemaName '.' tableName '.' columnName 
| tableName '.' columnName 
| columnName ;

columnNameQualified: (tableName '.')? columnName ;

columnNameList: columnName (',' columnName)* ;

columnNameQualifiedList: columnNameQualified (',' columnNameQualified)* ;

columnWild: '*' ;

columnWildQualified: ((variableName|tableName) '.')? columnWild ;

columnWildNameQualified: columnNameQualified
| columnWildQualified ;

constraintName: objectName ;

cursorName: objectName ;

functionName: objectName
| SYSTEM_FUNC_ID 
;

functionNameQualified: (schemaName '.')? functionName ;

indexName: objectName ;

labelName: objectName ;

label: labelName ':' ;

methodName: objectName ;

parameterName: LOCAL_ID ;

procedureName: objectName ;

procedureNameQualified: (schemaName '.')? procedureName ;

tableName: objectName
| TEMP_TABLE_ID ;

tableNameQualified: (schemaName '.')? tableName ;

transactionName: objectName ;

triggerName: objectName ;

triggerNameQualified: (schemaName '.')? triggerName ;

simpleTypeName: objectName ;

simpleTypeNameQualified: (schemaName '.')? simpleTypeName ;

// typeName: ID '(' ID ')'
// | QUOTED_ID '(' ID ')'
// | ID '(' integerLiteral ')'
// | QUOTED_ID '(' integerLiteral ')'
// | ID '(' integerLiteral ',' integerLiteral ')'
// | QUOTED_ID '(' integerLiteral ',' integerLiteral ')' 
// | simpleTypeName ;

typeName: objectName ('(' ( ID | (integerLiteral (',' integerLiteral)? ) ) ')')? ;

typeNameQualified: schemaName '.' typeName
| typeName ;

schemaName: objectName ;

systemVariableName: SYSTEM_VAR_ID ;

variableName: LOCAL_ID ;

viewName: objectName ;

viewNameQualified: schemaName '.' viewName
| viewName ;

xmlNamespaceName: objectName ;

xmlSchemaCollectionName: objectName ;

xmlSchemaCollectionNameQualified: schemaName '.' xmlSchemaCollectionName
| xmlSchemaCollectionName ;

externalName: 'EXTERNAL NAME' assemblyName '.' className '.' methodName ;

// optionalAs: 'AS'
// |  ;

stringLiteral: STRING_LITERAL ;

integerLiteral: INTEGER_LITERAL
| HEX_LITERAL ;

numberLiteral: REAL_LITERAL
| integerLiteral ;

priceLiteral: '$' numberLiteral;

literal: priceLiteral
| numberLiteral
| stringLiteral
| STRING_LITERAL 'COLLATE' collationName
| 'NULL' ;

toggle: 'ON'
| 'OFF' ;

script: batch ('GO' batch?)*;

// goBatchList: goBatch+;

batch: firstStatement
| statementList ;

// goBatch: 'GO' batch ;

statementList: statement+ ;

statement: statementBlock
| genericStatement ;

statementBlock: 'BEGIN' genericStatementList 'END' ;

genericStatementList: genericStatement genericStatementList? ;

TERMINATOR: ';' -> skip ;

firstStatement: createOrAlterFunctionStatement
| createOrAlterProcedureStatement
| createTriggerStatement
| createOrAlterViewStatement
| grantStatement ;

genericStatement: dDLStatement
| dMLStatement
| prgStatement
| label ;

dDLStatement: createFulltextStatement
| dropFulltextStatement
| dropFunctionStatement
| createIndexStatement
| dropIndexStatement
| dropProcedureStatement
| createTableStatement
| alterTableStatement
| dropTableStatement
| enableTriggerStatement
| disableTriggerStatement
| dropTriggerStatement
| createTypeStatement
| dropTypeStatement
| dropViewStatement
| createXmlSchemaCollectionStatement
| dropXmlSchemaCollectionStatement ;

dMLStatement: selectStatement
| insertStatement
| updateStatement
| deleteStatement
| mergeStatement ;

prgStatement
: useStatement
| beginTransactionStatement
| breakStatement
| closeStatement
| commitTransactionStatement
| continueStatement
| deallocateStatement
| declareStatement
| executeStatement
| fetchStatement
| gotoStatement
| ifStatement
| openStatement
| printStatement
| returnStatement
| rollbackTransactionStatement
| saveTransactionStatement
| setStatement
| tryCatchStatement
| waitforStatement
| whileStatement
| raiserrorStatement
| anyStatement ;

grantStatement: 'GRANT' grantPermissionList 'ON' objectNameQualified toPricipalWithOption ;

grantPermissionList: grantPermission
| grantPermission ',' grantPermissionList ;

grantPermission: EXECUTE
| 'REFERENCES'
| 'DELETE'
| 'INSERT'
| 'SELECT'
| 'UPDATE'
| 'ALL'
| 'ALL' 'PRIVILEGES' ;

databasePrincipalList: databasePrincipal (',' databasePrincipal)* ;

databasePrincipal: 'PUBLIC' ;

asDatabasePrincipal: 'AS' databasePrincipal
|  ;

withGrantOption: 'WITH' 'GRANT' 'OPTION'
|  ;

toPricipalWithOption: 'TO' databasePrincipalList withGrantOption asDatabasePrincipal ;

breakStatement: 'BREAK' ;

continueStatement: 'CONTINUE' ;

gotoStatement: 'GOTO' labelName ;

ifStatement: 'IF' predicate statement ('ELSE' statement)? ;

returnStatement: 'RETURN' expression? ;

tryCatchStatement: 'BEGIN' 'TRY' statementList 'END' 'TRY' 'BEGIN' 'CATCH' statementList 'END' 'CATCH' ;

whileStatement: 'WHILE' predicate statement ;

waitforStatement: 'WAITFOR' ID (stringLiteral|variableName);

printStatement: 'PRINT' expression ;

anyStatement: ID expressionList ;

useStatement: 'USE' ID ;

raiserrorStatement: 'RAISERROR' '(' expressionList ')' ('WITH' raiserrorOptionList)? ;

raiserrorOption: 'LOG'
| 'NOWAIT'
| 'SETERROR' ;

raiserrorOptionList: raiserrorOption (',' raiserrorOption)*;

// label: ID ':' ;

declareStatement: 'DECLARE' 
( variableName 'AS' 'TABLE' tableDefinitionGroup
| variableName 'TABLE' tableDefinitionGroup
| declareItemList
| cursorName cursorDefinition 
);

cursorDefinition: cursorType 'CURSOR' cursorOptionList 'FOR' selectStatement cursorUpdate ;

cursorType: 'INSENSITIVE'
| 'SCROLL'
|  ;

declareItem: variableName 'CURSOR'
| variableName 'AS'? typeNameQualified ('=' expression)?
;

declareItemList: declareItem (',' declareItem)* ;

cursorOptionList: ID cursorOptionList
|  ;

closeStatement: 'CLOSE' globalOrLocalCursor ;

openStatement: 'OPEN' globalOrLocalCursor ;

deallocateStatement: 'DEALLOCATE' globalOrLocalCursor ;

globalOrLocalCursor: variableName
| cursorName
| 'GLOBAL' cursorName ;

cursorUpdate: 'FOR UPDATE OF' columnNameList
| 'FOR UPDATE'
| 'FOR READ ONLY'
|  ;

fetchStatement: 'FETCH' cursorPosition globalOrLocalCursor ('INTO' variableNameList)?;

cursorPosition: 'NEXT' 'FROM'
| 'PRIOR' 'FROM'
| 'FIRST' 'FROM'
| 'LAST' 'FROM'
| 'ABSOLUTE' integerLiteral 'FROM'
| 'ABSOLUTE' variableName 'FROM'
| 'RELATIVE' integerLiteral 'FROM'
| 'RELATIVE' variableName 'FROM'
| 'FROM'
|  ;

variableNameList: variableName (',' variableName)* ;

optionalDefault: '=' literal
|  ;

destinationRowset: variableName
| tableNameQualified tableHintGroup ;

openxml: 'OPENXML' '(' variableName ',' stringLiteral ',' integerLiteral ')' optionalOpenxmlSchema
| 'OPENXML' '(' variableName ',' variableName ',' integerLiteral ')' optionalOpenxmlSchema
| 'OPENXML' '(' variableName ',' stringLiteral ')' optionalOpenxmlSchema
| 'OPENXML' '(' variableName ',' variableName ')' optionalOpenxmlSchema ;

optionalOpenxmlSchema: openxmlImplicitSchema
| openxmlExplicitSchema
|  ;

openxmlImplicitSchema: 'WITH' '(' tableNameQualified ')' ;

openxmlExplicitSchema: 'WITH' '(' openxmlColumnList ')' ;

openxmlColumnList: openxmlColumn (',' openxmlColumn)* ;

openxmlColumn: columnName typeName stringLiteral
| columnName typeName ;

setStatement: setVariableStatement
| setOptionStatement ;

setVariableStatement: 'SET' variableName 
( '=' expression
| '.' namedFunctionList
| '=' cursorDefinition 
);

setOptionStatement: 'SET' 
( ID setValueList
| TRANSACTION setValueList
| 'OFFSETS' setValueList
| 'ROWCOUNT' setValue
| 'STATISTICS' setValueList
| 'IDENTITY_INSERT' tableNameQualified toggle 
);

setValueList: setValue+ ;

setValue: ID
| 'READ' 'UNCOMMITTED'
| 'READ' 'COMMITTED'
| 'REPEATABLE' 'READ'
| toggle
| integerLiteral
| stringValue ;

stringValue: stringLiteral
| variableName ;

beginTransactionStatement: 'BEGIN' TRANSACTION (transactionIdentifier ('WITH' 'MARK' stringLiteral?)?)? ;

commitTransactionStatement: 'COMMIT'
| 'COMMIT' 'WORK'
| 'COMMIT' TRANSACTION
| 'COMMIT' TRANSACTION transactionIdentifier ;

rollbackTransactionStatement: 'ROLLBACK' ('WORK' | TRANSACTION transactionIdentifier? )? ;

saveTransactionStatement: 'SAVE' TRANSACTION transactionIdentifier ;

transactionIdentifier: transactionName
| variableName ;

createFulltextStatement: 'CREATE' 'FULLTEXT' 'INDEX' 'ON' 'TABLE' tableNameQualified fulltextColumnGroup 'KEY' 'INDEX' indexName fulltextChangeTracking ;

fulltextColumnGroup: '(' fulltextColumnList ')'
|  ;

fulltextColumnList: fulltextColumn (',' fulltextColumn)* ;

fulltextColumn: columnName fulltextColumnType optionalLanguage ;

fulltextColumnType: 'TYPE' 'COLUMN' typeNameQualified
|  ;

optionalLanguage: 'LANGUAGE' (integerLiteral|stringLiteral)
|  ;

fulltextChangeTracking: 'WITH' 'CHANGE_TRACKING' 
( auto='AUTO'
| manual='MANUAL'
| off='OFF'
| offNoPop=('OFF' ',' 'NO' 'POPULATION')
)
|  ;

dropFulltextStatement: 'DROP' 'FULLTEXT' 'INDEX' 'ON' tableNameQualified ;

createOrAlterFunctionStatement: op=('CREATE'|'ALTER') 'FUNCTION' functionNameQualified '(' optionalFunctionParameterList ')' 'RETURNS'
( typeNameQualified optionalFunctionOption 'AS'? statementBlock
| 'TABLE' optionalFunctionOption 'AS'? 'RETURN' functionInlineSelect
| variableName 'TABLE' tableDefinitionGroup optionalFunctionOption 'AS'? statementBlock
| typeNameQualified optionalFunctionOption 'AS'? externalName
| 'TABLE' tableDefinitionGroup optionalFunctionOption 'AS'? externalName
);

optionalFunctionParameterList: functionParameterList
|  ;

functionParameterList: functionParameter (',' functionParameter)* ;

functionParameter: parameterName 'AS' typeNameQualified optionalDefault optionalReadonly
| parameterName typeNameQualified optionalDefault optionalReadonly ;

functionInlineSelect: '(' functionInlineSelect ')'
| selectStatement ;

optionalFunctionOption: 'WITH' 'RETURNS' 'NULL' 'ON' 'NULL' 'INPUT'
| 'WITH' 'CALLED' 'ON' 'NULL' 'INPUT'
| 'WITH' EXECUTE 'AS' 'CALLER'
| 'WITH' 'SCHEMABINDING'
|  ;

dropFunctionStatement: 'DROP' 'FUNCTION' functionNameQualified ;

createOrAlterProcedureStatement: 'CREATE' PROCEDURE procedureNameQualified procedureParameterGroup procedureOptionGroup procedureFor asObjectBody
| 'ALTER' PROCEDURE procedureNameQualified procedureParameterGroup procedureOptionGroup procedureFor asObjectBody ;

asObjectBody: 'AS' statementList ;

procedureParameterGroup: '(' procedureParameterList ')'
| procedureParameterList
| '(' ')'
|  ;

procedureParameterList: procedureParameter (',' procedureParameter)* ;

procedureParameter: parameterName typeNameQualified optionalVarying optionalDefault optionalOutput optionalReadonly ;

optionalVarying: 'VARYING'
|  ;

optionalOutput: 'OUTPUT'
|  ;

optionalReadonly: 'READONLY'
|  ;

procedureOptionGroup: 'WITH' 'RECOMPILE'
|  ;

procedureFor: 'FOR' 'REPLICATION'
|  ;

dropProcedureStatement: 'DROP' PROCEDURE procedureNameQualified ;

executeStatement: EXECUTE 
( variableName '=' procedureNameQualified executeParameterGroup procedureOptionGroup
| procedureNameQualified executeParameterGroup procedureOptionGroup 
);

executeParameterGroup: executeParameterList
|  ;

executeParameterList: executeParameter (',' executeParameter)* ;

executeParameter: tableNameQualified
| (parameterName '=')? (variableName|systemVariableName|literal) optionalOutput
;

createTableStatement: 'CREATE' 'TABLE' tableNameQualified tableDefinitionGroup ('WITH' '(' tableOptionList ')')? ;

tableOptionList: tableOption (',' tableOption)* ;

// TODO: Support all of these.
tableOption: 
'DATA_COMPRESSION' '=' ('NONE'|'ROW'|'PAGE')
;

alterTableStatement: 'ALTER' 'TABLE' tableNameQualified 
( 'ALTER' 'COLUMN' columnName columnDefinition
| 'ALTER' 'COLUMN' columnName 'ADD' 'ROWGUIDCOL'
| 'ALTER' 'COLUMN' columnName 'DROP' 'ROWGUIDCOL'
| 'ALTER' 'COLUMN' columnName 'ADD' 'PERSISTED'
| 'ALTER' 'COLUMN' columnName 'DROP' 'PERSISTED'
| 'ALTER' 'COLUMN' columnName 'ADD' 'NOT' 'FOR' 'REPLICATION'
| 'ALTER' 'COLUMN' columnName 'DROP' 'NOT' 'FOR' 'REPLICATION'
| tableCheck 'ADD' tableDefinitionList
| tableCheck 'ADD' columnConstraint 'FOR' columnName
| tableCheck 'NOCHECK' 'CONSTRAINT' constraintName
| tableCheck 'NOCHECK' 'CONSTRAINT' 'ALL'
| tableCheck 'CHECK' 'CONSTRAINT' constraintName
| tableCheck 'CHECK' 'CONSTRAINT' 'ALL'
| 'DROP' constraintName
| 'DROP' 'CONSTRAINT' constraintName
| 'DROP' 'COLUMN' columnName 
);

tableCheck: 'WITH' 'CHECK'
| 'WITH' 'NOCHECK'
|  ;

dropTableStatement: 'DROP' 'TABLE' tableNameQualified ;

createOrAlterViewStatement: 'CREATE' 'VIEW' viewNameQualified columnNameGroup viewOptionalAttribute 'AS' selectStatement viewOptionalCheckOption
| 'ALTER' 'VIEW' viewNameQualified columnNameGroup viewOptionalAttribute 'AS' selectStatement viewOptionalCheckOption ;

columnNameGroup: '(' columnNameList ')'
|  ;

viewOptionalAttribute: 'WITH' 'VIEW_METADATA'
| 'WITH' 'SCHEMABINDING'
|  ;

viewOptionalCheckOption: 'WITH_CHECK_OPTION'
|  ;

dropViewStatement: 'DROP' 'VIEW' viewNameQualified ;

createIndexStatement: 'CREATE' 
( 'PRIMARY'? 'XML' 'INDEX' indexName 'ON' tableNameQualified '(' columnName ')' indexUsing indexOptionGroup?
| indexOptionalUnique constraintCluster 'INDEX' indexName 'ON' tableNameQualified '(' indexColumnList ')' ( 'INCLUDE' '(' columnNameList ')' )? indexOptionGroup?
);

indexColumnList: indexColumn (',' indexColumn)* ;

indexOptionalUnique: 'UNIQUE'
|  ;

indexColumn: columnName orderType ;

indexOptionGroup: 'WITH' '(' indexOptionList ')' ;

indexOptionList: indexOption (',' indexOption)* ;

indexOption: ID '=' integerLiteral
| ID '=' toggle ;

indexUsing: 'USING' 'XML' 'INDEX' indexName 'FOR' 'VALUE'
| 'USING' 'XML' 'INDEX' indexName 'FOR' 'PATH'
| 'USING' 'XML' 'INDEX' indexName 'FOR' 'PROPERTY'
|  ;

dropIndexStatement: 'DROP' 'INDEX' indexName 'ON' tableNameQualified indexOptionGroup? ;

createTriggerStatement: 'CREATE' 'TRIGGER' triggerNameQualified 'ON' tableNameQualified triggerType triggerOperationList optionalNotForReplication 'AS' statement ;

triggerType: 'FOR'
| 'INSTEAD' 'OF'
| 'AFTER' ;

triggerOperationList: triggerOperation (',' triggerOperation)* ;

triggerOperation: 'INSERT'
| 'UPDATE'
| 'DELETE' ;

optionalNotForReplication: 'NOT' 'FOR' 'REPLICATION'
|  ;

triggerTarget: tableNameQualified
| 'DATABASE'
| 'ALL' 'SERVER' ;

enableTriggerStatement: 'ENABLE' 
( 'TRIGGER' 'ALL' 'ON' triggerTarget
| 'TRIGGER' triggerNameQualifiedList 'ON' triggerTarget 
);

disableTriggerStatement: 'DISABLE' 
( 'TRIGGER' 'ALL' 'ON' triggerTarget
| 'TRIGGER' triggerNameQualifiedList 'ON' triggerTarget 
);

triggerNameQualifiedList: triggerNameQualified (',' triggerNameQualified)* ;

dropTriggerStatement: 'DROP' 'TRIGGER' triggerNameQualified ;

createXmlSchemaCollectionStatement: 'CREATE' 'XML' 'SCHEMA' 'COLLECTION' xmlSchemaCollectionNameQualified 'AS' expression ;

dropXmlSchemaCollectionStatement: 'DROP' 'XML' 'SCHEMA' 'COLLECTION' xmlSchemaCollectionNameQualified ;

createTypeStatement: 'CREATE' 'TYPE' simpleTypeNameQualified ( 'FROM' typeName typeConstraint | 'AS' 'TABLE' tableDefinitionGroup );

typeConstraint: 'NOT' 'NULL'
| 'NULL'
|  ;

dropTypeStatement: 'DROP' 'TYPE' simpleTypeNameQualified ;

tableDefinitionGroup: '(' tableDefinitionList ')' ;

tableDefinitionList: tableDefinition (',' tableDefinition)* ;

tableDefinition: columnName columnDefinition
| tableConstraint ;

tableConstraint: 'CONSTRAINT' constraintName 'PRIMARY' 'KEY' constraintCluster '(' indexColumnList ')' constraintIndex?
| 'CONSTRAINT' constraintName 'UNIQUE' constraintCluster '(' indexColumnList ')' constraintIndex?
| 'CONSTRAINT' constraintName 'FOREIGN' 'KEY' '(' columnNameList ')' 'REFERENCES' tableNameQualified columnNameGroup foreignKeyActionList
| 'CONSTRAINT' constraintName 'CHECK' optionalNotForReplication '(' predicate ')'
| 'PRIMARY' 'KEY' constraintCluster '(' indexColumnList ')' constraintIndex?
| 'UNIQUE' constraintCluster '(' indexColumnList ')' constraintIndex?
| 'FOREIGN' 'KEY' '(' columnNameList ')' 'REFERENCES' tableNameQualified columnNameGroup foreignKeyActionList
| 'CHECK' optionalNotForReplication '(' predicate ')' ;

foreignKeyActionList: foreignKeyAction* ;

foreignKeyAction: 'ON' 
( 'DELETE' 'NO' 'ACTION'
| 'DELETE' 'CASCADE'
| 'DELETE' 'SET' 'NULL'
| 'DELETE' 'SET' 'DEFAULT'
| 'UPDATE' 'NO' 'ACTION'
| 'UPDATE' 'CASCADE'
| 'UPDATE' 'SET' 'NULL'
| 'UPDATE' 'SET' 'DEFAULT' 
);

optionalForeignRefColumn: '(' columnName ')'
|  ;

columnDefinition: typeNameQualified columnConstraintList
| 'AS' expression computedColumnConstraintList ;

columnConstraintList: columnConstraint* ;

columnConstraint: namedColumnConstraint
| 'COLLATE' collationName
| 'IDENTITY'
| 'IDENTITY' '(' INTEGER_LITERAL ',' INTEGER_LITERAL ')'
| 'ROWGUIDCOL'
| 'NOT' 'NULL'
| 'NULL' ;

computedColumnConstraintList: computedColumnConstraint* ;

computedColumnConstraint: namedColumnConstraint
| 'PERSISTED'
| 'NOT' 'NULL' ;

namedColumnConstraint: 'CONSTRAINT' constraintName 'PRIMARY' 'KEY' constraintCluster constraintIndex?
| 'CONSTRAINT' constraintName 'UNIQUE' constraintCluster constraintIndex?
| 'CONSTRAINT' constraintName 'FOREIGN' 'KEY' 'REFERENCES' tableNameQualified optionalForeignRefColumn foreignKeyActionList
| 'CONSTRAINT' constraintName 'REFERENCES' tableNameQualified optionalForeignRefColumn foreignKeyActionList
| 'CONSTRAINT' constraintName 'CHECK' optionalNotForReplication '(' predicate ')'
| 'CONSTRAINT' constraintName 'DEFAULT' constraintDefaultValue
| 'PRIMARY' 'KEY' constraintCluster constraintIndex?
| 'UNIQUE' constraintCluster constraintIndex?
| 'FOREIGN' 'KEY' 'REFERENCES' tableNameQualified optionalForeignRefColumn foreignKeyActionList
| 'REFERENCES' tableNameQualified optionalForeignRefColumn foreignKeyActionList
| 'CHECK' optionalNotForReplication '(' predicate ')'
| 'DEFAULT' constraintDefaultValue ;

constraintCluster: 'CLUSTERED'
| 'NONCLUSTERED'
|  ;

constraintDefaultValue: expressionParens
| functionCall
| numberLiteral
| stringLiteral
| 'NULL' ;

constraintIndex: 'WITH' 'FILLFACTOR' '=' integerLiteral
| indexOptionGroup ;

queryOptions: 'WITH' 
( 'XMLNAMESPACES' '(' xmlNamespaceList ')' (',' cTEList)?
| cTEList
)
|  ;

xmlNamespaceList: xmlNamespace (',' xmlNamespace)* ;

xmlNamespace: stringLiteral 'AS' xmlNamespaceName ;

cTEList: cTE (',' cTE)* ;

cTE: aliasName columnNameGroup 'AS' '(' selectQuery ')' ;

top: 'TOP' '(' expression ')' optionalPercent optionalWithTies ;

optionalTop: top
|  ;

optionalPercent: 'PERCENT'
|  ;

optionalWithTies: 'WITH' 'TIES'
|  ;

outputClause: 'OUTPUT' columnItemList ('INTO' destinationRowset columnNameGroup)?
|  ;

queryHint: 'OPTION' '(' queryHintOptionList ')'
|  ;

// TODO: Add all the hint options.
queryHintOption: 'MAXRECURSION' INTEGER_LITERAL
| 'RECOMPILE' 
| ('HASH'|'ORDER') 'GROUP'
| 'FAST' INTEGER_LITERAL
| ('CONCAT'|'HASH'|'MERGE') 'UNION'
| 'OPTIMIZE' 'FOR' '(' variableName '=' expression ')'
;

queryHintOptionList: queryHintOption (',' queryHintOption)* ;

insertStatement: queryOptions 'INSERT' optionalTop optionalInto destinationRowset 
( 'DEFAULT' 'VALUES' queryHint 
| columnNameGroup outputClause 
    ( 'VALUES' '(' expressionList ')' (',' '(' expressionList ')')* queryHint
    | selectQuery queryHint
    | executeStatement queryHint )
);

optionalInto: 'INTO'
|  ;

mergeStatement: queryOptions 'MERGE' optionalTop optionalInto destinationRowset rowsetAlias? 'USING' sourceRowset 'ON' predicate mergeWhenMatchedList outputClause queryHint ;

mergeWhenMatchedList: mergeWhenMatched+ ;

mergeWhenMatched: 'WHEN' 'MATCHED' ('AND' predicate)? 'THEN' mergeMatched
| 'WHEN' 'NOT' 'MATCHED' 'BY' 'TARGET' ('AND' predicate)? 'THEN' mergeNotMatched
| 'WHEN' 'NOT' 'MATCHED' ('AND' predicate)? 'THEN' mergeNotMatched
| 'WHEN' 'NOT' 'MATCHED' 'BY' 'SOURCE' ('AND' predicate)? 'THEN' mergeMatched;

mergeMatched: 'UPDATE' 'SET' updateItemList
| 'DELETE' ;

mergeNotMatched: 'INSERT' columnNameGroup 'VALUES' '(' expressionList ')'
| 'INSERT' columnNameGroup 'DEFAULT' 'VALUES' ;

updateStatement: queryOptions 'UPDATE' optionalTop destinationRowset 'SET' updateItemList outputClause optionalFromClause whereClause queryHint ;

updateItemList: updateItem (',' updateItem)* ;

assignmentOperator: '=' | '+=' | '-=' | '*=' | '/=' | '%=' | '&=' | '^=' | '|=' ;

updateItem: columnNameFullyQualified assignmentOperator expression
| columnNameQualified '=' 'DEFAULT'
| variableName '=' (columnNameQualified '=')? expression
| ( variableName | tableName (columnName '.')? ) '.' namedFunctionList;

optionalFromClause: fromClause
|  ;

deleteStatement: queryOptions 'DELETE' optionalTop optionalFrom destinationRowset outputClause optionalFromClause whereClause queryHint ;

optionalFrom: 'FROM'
|  ;

selectStatement: queryOptions selectQuery queryHint ;

selectQuery: 'SELECT' restriction? topLegacy? columnItemList intoClause fromClause whereClause groupClause havingClause optionalOrderClause rowsetCombineClause? computeClause ;

topLegacy: 'TOP' (INTEGER_LITERAL|REAL_LITERAL) optionalPercent optionalWithTies
| top ;

columnItemList: columnItem (',' columnItem)* ;

columnItem: columnWildQualified
| expression 'AS' aliasName
| expression aliasName
| aliasName '=' expression
| variableName '=' expression
| expression 
;

restriction: 'ALL'
| 'DISTINCT' ;

intoClause: 'INTO' destinationRowset
|  ;

fromClause: 'FROM' source joinChain ;

source: '(' source joinChain ')' 
| sourceRowset ;

joinChain: join joinChain
|  ;

join: 'JOIN' source 'ON' predicate
| 'INNER' joinHint 'JOIN' source 'ON' predicate
| 'LEFT' joinHint 'JOIN' source 'ON' predicate
| 'LEFT' 'OUTER' joinHint 'JOIN' source 'ON' predicate
| 'RIGHT' joinHint 'JOIN' source 'ON' predicate
| 'RIGHT' 'OUTER' joinHint 'JOIN' source 'ON' predicate
| 'FULL' joinHint 'JOIN' source 'ON' predicate
| 'FULL' 'OUTER' joinHint 'JOIN' source 'ON' predicate
| 'CROSS' 'JOIN' source
| 'CROSS' 'APPLY' source
| 'OUTER' 'APPLY' source ;

joinHint: 'MERGE'
| 'HASH'
| 'LOOP'
|  ;

sourceRowset
: '(' 'VALUES' valuesList ')' rowsetAlias?
| '(' selectQuery ')' rowsetAlias? 
| schemaName '.' (
    tableName '.' (
        schemaName '.' tableName rowsetAlias?
        | tableName rowsetAlias?
        | namedFunction rowsetAlias?
    )
    | namedFunction rowsetAlias?
)
| variableName ('.' (columnName '.' namedFunction rowsetAlias? | namedFunction rowsetAlias?) | rowsetAlias)?
| tableNameQualified rowsetAlias? tableHintGroup
| openxml rowsetAlias?
| textTableFunction rowsetAlias?
| namedFunction rowsetAlias?
;

rowsetAlias: 'AS' aliasName '(' columnNameList ')'
| aliasName '(' columnNameList ')'
| 'AS' aliasName
| aliasName 
;

valuesList: '(' expressionList ')' (',' valuesList)* ;

textTableFunction: ('CONTAINSTABLE'|'FREETEXTTABLE') '(' 
    tableNameQualified ',' 
    (columnName|columnWild|'(' columnNameList ')') ',' 
    expression (',' 'LANGUAGE' literal)? optionalContainsTop 
')';

optionalContainsTop: ',' integerLiteral
|  ;

tableHintGroup: 'WITH' '(' tableHintList ')'
| 'WITH'? 
    ( '(NOLOCK)'
    | '(READUNCOMMITTED)'
    | '(UPDLOCK)'
    | '(REPEATABLEREAD)'
    | '(SERIALIZABLE)'
    | '(READCOMMITTED)'
    | '(TABLOCK)'
    | '(TABLOCKX)'
    | '(PAGLOCK)'
    | '(ROWLOCK)'
    | '(NOWAIT)'
    | '(READPAST)'
    | '(XLOCK)'
    | '(NOEXPAND)'
    )
|  ;

tableHintList: tableHint (',' tableHint)* ;

tableHint: 'INDEX' ('=' indexValue | '(' indexValueList ')' )
| 'HOLDLOCK'
| ID
| 'NOLOCK'
| 'NOWAIT'
| 'PAGLOCK'
| 'READCOMMITTED'
| 'READCOMMITTEDLOCK'
| 'READPAST'
| 'READUNCOMMITTED'
| 'REPEATABLEREAD'
| 'ROWLOCK'
| 'SERIALIZABLE'
| 'TABLOCK'
| 'TABLOCKX'
| 'UPDLOCK'
| 'XLOCK' ;

indexValueList: indexValue (',' indexValue)* ;

indexValue: INTEGER_LITERAL
| objectName
;

whereClause: 'WHERE' predicate
| 'WHERE' 'CURRENT' 'OF' (variableName|cursorName)
|  ;



groupClause: 'GROUP' 'BY' expressionList
|  ;

orderClause: 'ORDER' 'BY' orderList ;

optionalOrderClause: orderClause
|  ;

orderList: order (',' order)* ;

order: expression orderType ;

orderType: 'ASC'
| 'DESC'
|  ;

havingClause: 'HAVING' predicate
|  ;

computeClause: 'COMPUTE' columnItemList ('BY' columnItemList)?
|  ;

rowsetCombineClause: 'UNION' 'ALL'? ('('selectQuery')'|selectQuery)
| 'EXCEPT' selectQuery
| 'INTERSECT' selectQuery ;

forClause: 'FOR' 
( 'BROWSE'
| 'XML' 'AUTO' xmlDirectiveList
| 'XML' 'RAW' optionalElementName xmlDirectiveList
| 'XML' 'EXPLICIT' xmlDirectiveList
| 'XML' 'PATH' optionalElementName xmlDirectiveList
| 'READ' 'ONLY'
)
|  ;

optionalElementName: '(' STRING_LITERAL ')'
|  ;

xmlDirectiveList: (',' xmlDirective)* ;

xmlDirective: ID optionalElementName
| ID ID ;

predicate: predicateOr ;

predicateOr
: predicateAnd ('OR' predicateOr)?
;

predicateAnd
: predicateNot ('AND' predicateAnd)?
;

predicateNot: 'NOT'? predicateCompare;

predicateCompare: expression 
( op=('='|'>'|'<'|'>='|'<='|'<>'|'!='|'!<'|'!>') expression
| 'IS' 'NOT'? 'NULL'
| 'NOT'? 'BETWEEN' expression 'AND' expression
| 'NOT'? 'LIKE' expression ('ESCAPE' STRING_LITERAL)?
| 'NOT'? 'IN' tuple
)
| predicateFunction ;

tuple: '(' (selectQuery|expressionList) ')' ;

expressionList: expression (',' expression)* ;

predicateFunction
: 'EXISTS' '(' selectQuery ')'
| 'UPDATE' '(' columnName ')'
| ('CONTAINS'|'FREETEXT') 
    ('(' '(' columnNameQualifiedList ')' ',' expression (',' 'LANGUAGE' literal)? ')'
    |'(' columnWildNameQualified ',' expression (',' 'LANGUAGE' literal)? ')'
    )
| predicateParens ;

predicateParens: '(' predicate ')' ;

expression: expressionAdd ;

expressionAdd: expressionMult (op=('+'|'-'|'&'|'|'|'^') expressionAdd)? ;

expressionMult: expressionNegate (op=('*'|'/'|'%') expressionMult)? ;

expressionNegate: op=('-'|'+'|'~')? expressionCase ;

rankingArguments: 'PARTITION' 'BY' expressionList optionalOrderClause
| orderClause ;

expressionCase: 'CASE' ( caseWhenPredicateList | expression caseWhenExpressionList) ('ELSE' expression)? 'END'
| collatedValue ;

caseWhenExpressionList: caseWhenExpression+ ;

caseWhenExpression: 'WHEN' expression 'THEN' expression ;

caseWhenPredicateList: caseWhenPredicate+ ;

caseWhenPredicate: 'WHEN' predicate 'THEN' expression ;

collatedValue: value 'COLLATE' collationName
| value
| literal ;

functionCall: ('CAST'|'TRY_CAST') '(' expression 'AS' typeName ')'
| 'COALESCE' '(' expressionList ')'
| 'NULLIF' '(' expression ',' expression ')'
| 'LEFT' '(' expression ',' expression ')'
| 'RIGHT' '(' expression ',' expression ')'
| ('CONVERT'|'TRY_CONVERT') '(' typeName ',' expression (',' INTEGER_LITERAL)? ')'
| 'COUNT' '(' restriction? (columnWild|expression) ')'
| 'IDENTITY' '(' typeName (',' integerLiteral ',' integerLiteral)? ')'
| 'CURRENT_TIMESTAMP'
| 'CURRENT_USER'
| 'SESSION_USER'
| 'SYSTEM_USER'
| 'USER'
| namedFunction ;

namedFunction: callable '(' expressionList? ')' ;

callable: objectName ('.' objectName|SYSTEM_FUNC_ID)
| objectName
| SYSTEM_FUNC_ID
;

namedFunctionList: namedFunction ('.' namedFunction)* ;

value
: functionCall ('.' namedFunctionList)?
| schemaName '.' columnNameQualified
| columnNameQualified
| variableName
| expressionParens ('.' namedFunctionList)?
| tableName '.' namedFunctionList
| tableName '.' columnName '.' namedFunctionList
| functionCall 'OVER' '(' rankingArguments ')'
| systemVariableName 
;

expressionParens: ('ANY'|'ALL')? '(' (selectQuery|expression) ')' ;

