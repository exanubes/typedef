const jsdoc = require('eslint-plugin-jsdoc');
const prettier = require('eslint-plugin-prettier');
const html = require('eslint-plugin-html');
const htmlParser = require('@html-eslint/parser');
const htmlPlugin = require('@html-eslint/eslint-plugin');

module.exports = [
    {
        ignores: ['dist/', '.parcel-cache/', 'node_modules/', 'src/js/rpc/wasm_exec.js'],
    },
    // JavaScript configuration
    {
        files: ['**/*.js'],
        languageOptions: {
            ecmaVersion: 'latest',
            sourceType: 'module',
            globals: {
                window: 'readonly',
                document: 'readonly',
                console: 'readonly',
                navigator: 'readonly',
                localStorage: 'readonly',
                indexedDB: 'readonly',
                Go: 'readonly',
                WebAssembly: 'readonly',
            },
        },
        plugins: {
            jsdoc,
            prettier,
        },
        rules: {
            'prettier/prettier': 'error',
            camelcase: 'off',

            'jsdoc/require-jsdoc': [
                'warn',
                {
                    require: {
                        FunctionDeclaration: true,
                        FunctionExpression: false,
                        ArrowFunctionExpression: false,
                        ClassDeclaration: true,
                        ClassExpression: false,
                        MethodDefinition: true,
                    },
                    contexts: [
                        "ExportNamedDeclaration > VariableDeclaration[kind='const'] > VariableDeclarator > ArrowFunctionExpression",
                        "ExportNamedDeclaration > VariableDeclaration[kind='const'] > VariableDeclarator > FunctionExpression",
                    ],
                    exemptEmptyFunctions: true,
                },
            ],
            'jsdoc/require-param': 'off',
            'jsdoc/require-param-description': 'off',
            'jsdoc/require-param-type': 'off',
            'jsdoc/require-returns': 'off',
            'jsdoc/require-returns-description': 'off',
            'jsdoc/require-returns-type': 'off',
            'jsdoc/check-param-names': 'error',
            'jsdoc/check-types': 'off',
            'jsdoc/valid-types': 'off',
            'jsdoc/check-tag-names': 'error',
            'jsdoc/no-types': 'off',
            'jsdoc/tag-lines': ['warn', 'any', { startLines: 1 }],
            'jsdoc/require-hyphen-before-param-description': ['warn', 'never'],
        },
        settings: {
            jsdoc: {
                mode: 'jsdoc',
                tagNamePreference: {
                    returns: 'returns',
                    return: 'returns',
                },
                preferredTypes: {
                    String: 'string',
                    Number: 'number',
                    Boolean: 'boolean',
                    Object: 'object',
                    Array: 'array',
                    Function: 'function',
                    Null: 'null',
                    Undefined: 'undefined',
                },
                structuredTags: {
                    template: {
                        name: 'namepath-defining',
                    },
                },
            },
        },
    },
    // HTML configuration
    {
        files: ['**/*.html'],
        plugins: {
            '@html-eslint': htmlPlugin,
        },
        languageOptions: {
            parser: htmlParser,
        },
        rules: {
            '@html-eslint/indent': ['error', 4],
            '@html-eslint/require-doctype': 'error',
            '@html-eslint/require-lang': 'error',
            '@html-eslint/require-meta-charset': 'error',
            '@html-eslint/require-meta-viewport': 'error',
            '@html-eslint/require-title': 'error',
            '@html-eslint/no-duplicate-attrs': 'error',
            '@html-eslint/no-duplicate-id': 'error',
            '@html-eslint/require-button-type': 'error',
            '@html-eslint/require-img-alt': 'error',
            '@html-eslint/no-inline-styles': 'warn',
            '@html-eslint/require-meta-description': 'warn',
            '@html-eslint/attrs-newline': 'off',
            '@html-eslint/element-newline': 'off',
        },
    },
];
